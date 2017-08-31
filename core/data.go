package core

import (
	agent "github.com/gofinance/ib"
)

// Hist Data
const (
	HistBarSize1Sec  = agent.HistBarSize1Sec
	HistBarSize5Sec  = agent.HistBarSize5Sec
	HistBarSize15Sec = agent.HistBarSize5Sec
	HistBarSize30Sec = agent.HistBarSize30Sec
	HistBarSize1Min  = agent.HistBarSize1Min
	HistBarSize2Min  = agent.HistBarSize2Min
	HistBarSize3Min  = agent.HistBarSize3Min
	HistBarSize5Min  = agent.HistBarSize5Min
	HistBarSize15Min = agent.HistBarSize15Min
	HistBarSize30Min = agent.HistBarSize30Min
	HistBarSize1Hour = agent.HistBarSize1Hour
	HistBarSize1Day  = agent.HistBarSize1Day

	HistTrades     = agent.HistTrades
	HistMidpoint   = agent.HistMidpoint
	HistBid        = agent.HistBid
	HistAsk        = agent.HistAsk
	HistBidAsk     = agent.HistBidAsk
	HistVolatility = agent.HistVolatility
	HistOptionIV   = agent.HistOptionIV
)

// TickType enum
const (
	TickBidSize               = agent.TickBidSize
	TickBid                   = agent.TickBid
	TickAsk                   = agent.TickAsk
	TickAskSize               = agent.TickAskSize
	TickLast                  = agent.TickLast
	TickLastSize              = agent.TickLastSize
	TickHigh                  = agent.TickHigh
	TickLow                   = agent.TickLow
	TickVolume                = agent.TickVolume
	TickClose                 = agent.TickClose
	TickBidOptionComputation  = agent.TickBidOptionComputation
	TickAskOptionComputation  = agent.TickAskOptionComputation
	TickLastOptionComputation = agent.TickLastOptionComputation
	TickModelOption           = agent.TickModelOption
	TickOpen                  = agent.TickOpen
	TickLow13Week             = agent.TickLow13Week
	TickHigh13Week            = agent.TickHigh13Week
	TickLow26Week             = agent.TickLow26Week
	TickHigh26Week            = agent.TickHigh26Week
	TickLow52Week             = agent.TickLow52Week
	TickHigh52Week            = agent.TickHigh52Week
	TickAverageVolume         = agent.TickAverageVolume
	TickOpenInterest          = agent.TickOpenInterest
	TickOptionHistoricalVol   = agent.TickOptionHistoricalVol
	TickOptionImpliedVol      = agent.TickOptionImpliedVol
	TickOptionBidExch         = agent.TickOptionBidExch
	TickOptionAskExch         = agent.TickOptionAskExch
	TickOptionCallOpenInt     = agent.TickOptionCallOpenInt
	TickOptionPutOpenInt      = agent.TickOptionPutOpenInt
	TickOptionCallVolume      = agent.TickOptionCallVolume
	TickOptionPutVolume       = agent.TickOptionPutVolume
	TickIndexFuturePremium    = agent.TickIndexFuturePremium
	TickBidExch               = agent.TickBidExch
	TickAskExch               = agent.TickAskExch
	TickAuctionVolume         = agent.TickAuctionVolume
	TickAuctionPrice          = agent.TickAuctionPrice
	TickAuctionImbalance      = agent.TickAuctionImbalance
	TickMarkPrice             = agent.TickMarkPrice
	TickBidEFPComputation     = agent.TickBidEFPComputation
	TickAskEFPComputation     = agent.TickAskEFPComputation
	TickLastEFPComputation    = agent.TickLastEFPComputation
	TickOpenEFPComputation    = agent.TickOpenEFPComputation
	TickHighEFPComputation    = agent.TickHighEFPComputation
	TickLowEFPComputation     = agent.TickLowEFPComputation
	TickCloseEFPComputation   = agent.TickCloseEFPComputation
	TickLastTimestamp         = agent.TickLastTimestamp
	TickShortable             = agent.TickShortable
	TickFundamentalRations    = agent.TickFundamentalRations
	TickRTVolume              = agent.TickRTVolume
	TickHalted                = agent.TickHalted
	TickBidYield              = agent.TickBidYield
	TickAskYield              = agent.TickAskYield
	TickLastYield             = agent.TickLastYield
	TickCustOptionComputation = agent.TickCustOptionComputation
	TickTradeCount            = agent.TickTradeCount
	TickTradeRate             = agent.TickTradeRate
	TickVolumeRate            = agent.TickVolumeRate
	TickLastRTHTrade          = agent.TickLastRTHTrade
	TickNotSet                = agent.TickNotSet
	TickRegulatoryImbalance   = agent.TickRegulatoryImbalance
)

type HistDataBarSize = agent.HistDataBarSize
