package btc_arbitrage

import (
	"encoding/json"
)

const (
	EventOKCoinAddChannel          = "addChannel"
	EventOKCoinSubscribeCNYBTCTick = "ok_sub_spotcny_btc_ticker"
)

// "{'event':'addChannel','channel':'ok_sub_spotcny_btc_ticker'}"
type WebSocketOKCoinCmd struct {
	Event   string `json:"event"`
	Channel string `json:"channel"`
}

// [{
//     "channel":"ok_sub_spotcny_btc_ticker",
//     "data":{
//         "buy":2478.3,
//         "high":2555,
//         "last":2478.51,
//         "low":2466,
//         "sell":2478.5,
//         "timestamp":1411718074965,
//         "vol":49020.30
//     }
// }]

type WebSocketOKCoinTick struct {
	Buy       float64 `json:"buy"`
	Sell      float64 `json:"sell"`
	Timestamp int     `json:"timestamp“`
	Volume    float64 `json:"vol"`
}

type WebSocketOKCoinTickRsp struct {
	Channel string              `jason:"channel"`
	Data    WebSocketOKCoinTick `json:"data"`
}

///////////////////////////////////////////////////////////////////

const (
	EventHUOBIPing = "ping"
)

const (
	HUOBISubDepthStep0 = "market.btccny.depth.step0"
)

const (
	HUOBISubRspChannelKey = "ch"
	HUOBISubRspTickKey    = "tick"
)

const (
	HUOBISubRspDepth = 0
)

type HUOBIPing struct {
	Ping int `json:"ping"`
}

type HUOBIPong struct {
	Pong int `json:"pong"`
}

type HUOBISubReq struct {
	Sub string `json:"sub"`
	ID  string `json:"id"`
}

type HUOBISubRsp struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	Subbed    string `json:"subbed"`
	Timestamp int    `json:"ts"`
}

type HUOBITick struct {
	ID        int             `json:"id"`
	Timestamp int             `json:"ts"`
	Bids      json.RawMessage `json:"bids"`
	Asks      json.RawMessage `json:"asks"`
}

type HUOBITickRsp struct {
	Channel   string    `json:"ch"`
	Timestamp int       `json:"ts"`
	Tick      HUOBITick `json:"tick"`
}
