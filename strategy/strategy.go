package strategy

import (
	"github.com/quantPlatform/quant/core"
	"github.com/quantPlatform/quant/strategy/btc_arbitrage"
	"github.com/quantPlatform/quant/strategy/macd"
)

func Init() {
	register()
	btc_arbitrage.Init()
}

func registerMACDStrategy() {
	macdStrategy := &macd.MACDStrategy{}
	macdStrategy.Init()
	macdStrategy.Register(core.Strategies())
}

func register() {
	registerMACDStrategy()
}
