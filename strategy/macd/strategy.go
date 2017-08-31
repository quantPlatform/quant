package macd

import (
	"time"

	ta "github.com/markcheno/go-talib"
	"github.com/quantPlatform/quant/core"
	log "github.com/sirupsen/logrus"
)

const strategyName = "macd"

const (
	macdFastPeriod   = 12
	macdSlowPeriod   = 26
	macdSignalPeriod = 9
)

type macd struct {
	date []time.Time
	dif  []float64
	dea  []float64
	macd []float64
}

const (
	positionLow   = 0
	postionMiddle = 1
	positionHigh  = 2
)

type MACDStrategy struct {
	core.Strategy

	dataLevels []core.HistDataBarSize

	bars30Min core.HistoricalData
	macd30Min *macd
	macd1Day  *macd

	tradeSignalIndex []int
}

func (strategy *MACDStrategy) Init() {
	strategy.dataLevels = []core.HistDataBarSize{core.HistBarSize30Min, core.HistBarSize1Day}
	const reserved = 1000
	strategy.tradeSignalIndex = make([]int, 0, 100)
	//strategy.bars30Min = make(core.HistoricalData, reserved)
}

func (strategy *MACDStrategy) Register(registry core.StrategyRegistry) {
	if registry == nil {
		return
	}

	registry.Register(strategy)
}

func (strategy *MACDStrategy) Name() string {
	return strategyName
}

func (strategy *MACDStrategy) GetDataLevels() []core.HistDataBarSize {
	return strategy.dataLevels
}

func (strategy *MACDStrategy) Condition() {
	log.Debug("run condition for ", strategy.Name())

	const unstableCount = 100
	const validCount = 10

	if len(strategy.macd1Day.date) < unstableCount+validCount {
		log.Error("macd1day is too short")
		return
	}

	if len(strategy.macd30Min.date) == unstableCount+validCount {
		log.Error("macd30min is too short")
		return
	}

	for dayIndex := unstableCount; dayIndex < len(strategy.macd1Day.date); dayIndex++ {
		preDifItem := strategy.macd1Day.dif[dayIndex-1]
		difItem := strategy.macd1Day.dif[dayIndex]

		if difItem < 0 {
			log.WithFields(log.Fields{"dif": difItem})
			continue
		}

		macdItem := strategy.macd1Day.macd[dayIndex]
		if macdItem > 0 && preDifItem <= difItem {
			log.WithFields(log.Fields{"macd": macdItem, "preDif": preDifItem, "dif": difItem})
			continue
		}

		yearDay := strategy.macd1Day.date[dayIndex].YearDay()
		min30Index := strategy.find30MinGoldCross(yearDay)
		if min30Index < 0 {
			log.WithFields(log.Fields{"30MinGoldCross": min30Index})
			continue
		}

		strategy.tradeSignalIndex = append(strategy.tradeSignalIndex, min30Index)
	}

	log.Debug("trade signals: %#v", strategy.tradeSignalIndex)
}

func (strategy *MACDStrategy) find30MinGoldCross(yearDay int) int {
	const unstableCount = 100

	for min30Index := unstableCount; min30Index < len(strategy.macd30Min.date); min30Index++ {
		dateItem := &strategy.macd30Min.date[min30Index]
		if dateItem.YearDay() != yearDay {
			continue
		}

		difItem := strategy.macd30Min.dif[min30Index]
		if difItem < 0 {
			continue
		}

		deaItem := strategy.macd30Min.dea[min30Index]
		preDifItem := strategy.macd30Min.dif[min30Index-1]
		preDeaItem := strategy.macd30Min.dea[min30Index-1]
		if preDifItem < preDeaItem && difItem >= deaItem {
			return min30Index
		}
	}

	return -1
}

func (strategy *MACDStrategy) OnTick(ask, bid, last float64) {
	log.WithFields(log.Fields{"ask": ask, "bid": bid, "last": last})
}

func (strategy *MACDStrategy) OnHistBarSize30Min(data core.HistoricalData) {
	log.Debug("[OnMinute] data : %#v", data)
	strategy.bars30Min = data
	strategy.macd30Min = strategy.calculateMACD(data)
}

func (strategy *MACDStrategy) OnHistBarSize1Day(data core.HistoricalData) {
	log.WithFields(log.Fields{"func": "OnHistBarSize1Day", "data": data})
	strategy.macd1Day = strategy.calculateMACD(data)
}

func (strategy *MACDStrategy) calculateMACD(data core.HistoricalData) *macd {
	date, close := data.FilterClose()

	macdResult := &macd{date: date}
	macdResult.dif, macdResult.dea, macdResult.macd =
		ta.Macd(close, macdFastPeriod, macdSlowPeriod, macdSignalPeriod)
	return macdResult
}

func (strategy *MACDStrategy) Order() {
	for i := 0; i < len(strategy.tradeSignalIndex); i++ {
		signalIndex := strategy.tradeSignalIndex[i]
		price := strategy.bars30Min[signalIndex].Close
		log.WithFields(log.Fields{"buyPrice": price})
	}
}
