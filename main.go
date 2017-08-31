package main

import (
	"github.com/quantPlatform/quant/core"
	"github.com/quantPlatform/quant/strategy/btc_arbitrage"

	"github.com/quantPlatform/quant/lp"
	"github.com/quantPlatform/quant/strategy"
	log "github.com/sirupsen/logrus"
)

func init() {
	core.Init()
	lp.Init()
	strategy.Init()
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debug("Enter main...")

	lp := core.GetLp("ib")
	if lp == nil {
		return
	}

	// dataLevels := []string{"tick", "minute", "day"}
	// dataHandler
	// strategy
	// contracts := []core.Contract{
	// 	{
	// 		Symbol:       "USD",
	// 		SecurityType: "CASH",
	// 		Exchange:     "IDEALPRO",
	// 		Currency:     "JPY",
	// 	},
	// }

	// // contracts := []core.Contract{
	// // 	{
	// // 		Symbol:       "700",
	// // 		SecurityType: "STK",
	// // 		Exchange:     "SEHK",
	// // 		Currency:     "HKD",
	// // 	},
	// // }

	// lp.Start("macd", contracts)

	//okCoin := btc_arbitrage.OKCoin{}
	// okCoin.GetBars("1day")
	//okCoin.StartTickEngine()

	runOKCoin()
	runHUOBI()

	exit := make(chan bool)
	c := <-exit
	log.Print(c)
}

func runOKCoin() {
	okCoin := btc_arbitrage.OKCoin{}
	okCoin.StartTickEngine()
}

func runHUOBI() {
	huobi := btc_arbitrage.HUOBI{}
	huobi.StartTickEngine()
}
