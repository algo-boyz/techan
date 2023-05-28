package techan

import (
	"math"

	"github.com/algo-boyz/decimal"
)

type kIndicator struct {
	Close    Indicator
	minValue Indicator
	maxValue Indicator
	window   int
}

// NewFastStochasticIndicator returns a derivative Indicator which returns the fast stochastic indicator (%K) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewFastStochasticIndicator(series *TimeSeries, timeframe int) Indicator {
	return kIndicator{
		Close:    NewCloseIndicator(series),
		minValue: NewMinimumValueIndicator(NewLowPriceIndicator(series), timeframe),
		maxValue: NewMaximumValueIndicator(NewHighPriceIndicator(series), timeframe),
		window:   timeframe,
	}
}

func (k kIndicator) Calculate(index int) decimal.Decimal {
	closeVal := k.Close.Calculate(index)
	minVal := k.minValue.Calculate(index)
	maxVal := k.maxValue.Calculate(index)

	if minVal.Equal(maxVal) {
		return decimal.NewFromFloat(math.MaxFloat64)
	}

	return closeVal.Sub(minVal).Div(maxVal.Sub(minVal)).Mul(decimal.NewFromInt(100))
}

type dIndicator struct {
	k      Indicator
	window int
}

// NewSlowStochasticIndicator returns a derivative Indicator which returns the slow stochastic indicator (%D) for the
// given window.
// https://www.investopedia.com/terms/s/stochasticoscillator.asp
func NewSlowStochasticIndicator(k Indicator, window int) Indicator {
	return dIndicator{k, window}
}

func (d dIndicator) Calculate(index int) decimal.Decimal {
	return NewSimpleMovingAverage(d.k, d.window).Calculate(index)
}
