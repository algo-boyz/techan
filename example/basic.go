package example

import (
	"strconv"
	"time"

	"github.com/algo-boyz/decimal"
	"github.com/algo-boyz/techan"
)

// BasicEma is an example of how to create a basic Exponential moving average indicator
// based on the close prices of a timeseries from your exchange of choice.
func BasicEma() techan.Indicator {
	series := techan.NewTimeSeries()

	// fetch this from your preferred exchange
	dataset := [][]string{
		// Timestamp, Open, Close, High, Low, volume
		{"1234567", "1", "2", "3", "5", "6"},
	}

	for _, datum := range dataset {
		start, _ := strconv.ParseInt(datum[0], 10, 64)
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

		candle := techan.NewCandle(period)
		candle.Open, _ = decimal.NewFromString(datum[1])
		candle.Close, _ = decimal.NewFromString(datum[2])
		candle.High, _ = decimal.NewFromString(datum[3])
		candle.Low, _ = decimal.NewFromString(datum[4])

		series.AddCandle(candle)
	}

	Closes := techan.NewCloseIndicator(series)
	movingAverage := techan.NewEMAIndicator(Closes, 10) // Create an exponential moving average with a window of 10

	return movingAverage
}
