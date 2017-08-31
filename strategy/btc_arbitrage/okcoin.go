package btc_arbitrage

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

const (
	okCoinChinaURL          = "https://www.okcoin.cn"
	okCoinChinaWebSocketURL = "wss://real.okcoin.cn:10440/websocket/okcoinapi"
)

type Param struct {
	key   string
	value string
}

type Params []Param

func (params Params) Sort() {
	sort.Slice(params, func(i, j int) bool { return params[i].key < params[j].key })
}

func (params Params) ToString() string {
	params.Sort()

	ret := ""
	for i := 0; i < len(params); i++ {
		ret += fmt.Sprintf("%s=%s", params[i].key, params[i].value)
	}

	return ret
}

////////////////////////////////////////////////////////////
type OKCoin struct {
	appKey     string
	secrectKey string
}

func (okCoin *OKCoin) SetKeys(appKey, secrectKey string) {
	okCoin.appKey = appKey
	okCoin.secrectKey = secrectKey
}

func (okCoin *OKCoin) md5(params Params) string {
	paramStr := params.ToString()
	paramStr += fmt.Sprintf("&secrect_key=%s", okCoin.secrectKey)

	ctx := md5.New()
	ctx.Write([]byte(paramStr))
	return hex.EncodeToString(ctx.Sum(nil))
}

func (okCoin *OKCoin) OnTicker() {

}

func (okCoin *OKCoin) GetBars(period string) {
	url := okCoinChinaURL + fmt.Sprintf("/api/v1/kline.do?symbol=btc_cny&type=%s", period)
	log.Debug(url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err.Error())
		return
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	retStr := string(body)
	cnt := strings.Count(retStr, ",")
	log.Debugf("okcoin return : %d", cnt)
}

func (okCoin *OKCoin) StartTickEngine() {
	conn, _, err := websocket.DefaultDialer.Dial(okCoinChinaWebSocketURL, nil)
	if err != nil {
		log.Errorf("Websocket to OKCoin fail : %s", err.Error())
		return
	}

	cmd := WebSocketOKCoinCmd{Event: EventOKCoinAddChannel, Channel: EventOKCoinSubscribeCNYBTCTick}
	if err := conn.WriteJSON(cmd); err != nil {
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

			tickRsp := WebSocketOKCoinTickRsp{}
			size := len(message)
			if err := json.Unmarshal(message[1:size-1], &tickRsp); err != nil {
				log.Errorf("%s", err.Error())
				continue
			}

			if tickRsp.Data.Sell == 0 || tickRsp.Data.Buy == 0 {
				log.Debug("%s", string(message))
				continue
			}

			log.Debugf("[okcoin] ask: %f\tbid: %f", tickRsp.Data.Sell, tickRsp.Data.Buy)
			bboMgr.UpdateOKCoin(int(tickRsp.Data.Sell*toInt), int(tickRsp.Data.Buy*toInt))
		}
	}()
}
