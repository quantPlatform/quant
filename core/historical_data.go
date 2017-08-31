package core

import (
	"time"
	agent "github.com/gofinance/ib"
)

type HistoricalDataItem = agent.HistoricalDataItem
type HistoricalData  []HistoricalDataItem

func (data HistoricalData) FilterClose() ([]time.Time, []float64) {
	size := len(data)
	date := make([]time.Time, size, size)
	close := make([]float64, size, size)

	for i := 0; i < size; i++ {
		close[i] = data[i].Close
		date[i] = data[i].Date
	}

	return date, close
}