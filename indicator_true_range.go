package techan

import "github.com/algo-boyz/decimal"

type trueRangeIndicator struct {
	series *TimeSeries
}

// NewTrueRangeIndicator returns a base indicator
// which calculates the true range at the current point in time for a series
// https://www.investopedia.com/terms/a/atr.asp
func NewTrueRangeIndicator(series *TimeSeries) Indicator {
	return trueRangeIndicator{
		series: series,
	}
}

func (tri trueRangeIndicator) Calculate(index int) decimal.Decimal {
	if index-1 < 0 {
		return decimal.Zero
	}

	candle := tri.series.Candles[index]
	previousClose := tri.series.Candles[index-1].Close

	trueHigh := decimal.Max(candle.High, previousClose)
	trueLow := decimal.Min(candle.Low, previousClose)

	return trueHigh.Sub(trueLow)
}
