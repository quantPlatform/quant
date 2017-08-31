package btc_arbitrage

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/quantPlatform/quant/util"
	log "github.com/sirupsen/logrus"
)

const (
	huobiWebsocketURL = "wss://api.huobi.com/ws"
)

type HUOBI struct {
}

func (huobi *HUOBI) PongIfNecessary(conn *websocket.Conn, ret string) bool {
	if !strings.Contains(ret, EventHUOBIPing) {
		return false
	}

	ping := HUOBIPing{}
	if err := json.Unmarshal([]byte(ret), &ping); err != nil {
		log.Errorf("Unmarshal fail ： %s", err.Error())
		return false
	}

	huobi.Pong(conn, ping.Ping)
	return true

}

func (huobi *HUOBI) Pong(conn *websocket.Conn, value int) {
	pong := HUOBIPong{Pong: value}
	if err := websocket.WriteJSON(conn, pong); err != nil {
		log.Error("Pong error : %s", err.Error())
	}
}

func (huobi *HUOBI) StartTickEngine() {
	conn, _, err := websocket.DefaultDialer.Dial(huobiWebsocketURL, nil)
	if err != nil {
		log.Errorf("Websocket to HUOBI fail : %s", err.Error())
		return
	}

	// cmd := "{'req': 'market.btccny.depth.step0','id': 'subTick'}"
	subDepthStep0 := HUOBISubReq{Sub: HUOBISubDepthStep0, ID: "subTick"}
	if err := conn.WriteJSON(subDepthStep0); err != nil {
		log.Errorf("write message fail : %s", err.Error())
		return
	}

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Error("websocket read message fail : ", err)
				return
			}

			r, _ := gzip.NewReader(bytes.NewReader(message))
			defer r.Close()
			data, _ := ioutil.ReadAll(r)
			retStr := string(data)

			if huobi.PongIfNecessary(conn, retStr) {
				log.Debugf("%s", data)
				continue
			}

			// log.Debugf("%s", retStr)
			if HUOBISubRspDepth == huobi.subRspType(retStr) {
				huobi.OnDepthStep0(retStr)
			}
		}
	}()
}

func (huobi *HUOBI) subRspType(rsp string) int {
	if strings.Contains(rsp, "ch") && strings.Contains(rsp, "tick") {
		return HUOBISubRspDepth
	}

	return -1
}

func (huobi *HUOBI) OnDepthStep0(str string) {
	rsp := HUOBITickRsp{}
	if err := json.Unmarshal([]byte(str), &rsp); err != nil {
		log.Errorf("%s", err.Error())
		return
	}

	askStr := string(rsp.Tick.Asks)
	asks := strings.Split(askStr[2:20], ",")
	ask := util.StrToFloat64(asks[0], 8)

	bidStr := string(rsp.Tick.Bids)
	bids := strings.Split(bidStr[2:20], ",")
	bid := util.StrToFloat64(bids[0], 8)
	log.Debugf("[huobi] ask : %f\tbid : %f", ask, bid)

	bboMgr.UpdateHUOBI(int(ask*toInt), int(bid*toInt))
}
