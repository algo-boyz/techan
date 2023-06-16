package techan

import "github.com/algo-boyz/decimal"

type smaIndicator struct {
	indicator Indicator
	window    int
}

// NewSimpleMovingAverage returns a derivative Indicator which returns the average of the current value and preceding
// values in the given windowSize.
func NewSimpleMovingAverage(indicator Indicator, window int) Indicator {
	return smaIndicator{indicator, window}
}

func (sma smaIndicator) Calculate(index int) decimal.Decimal {
	if index < sma.window-1 {
		return decimal.Zero
	}

	sum := decimal.Zero
	for i := index; i > index-sma.window; i-- {
		sum = sum.Add(sma.indicator.Calculate(i))
	}

	period := decimal.NewFromInt(int64(sma.window))
	result := sum.Div(period)

	return result
}
