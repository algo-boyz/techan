package techan

import (
	"math"

	"github.com/algo-boyz/decimal"
)

type relativeStrengthIndexIndicator struct {
	rsIndicator Indicator
	oneHundred  decimal.Decimal
}

// NewRelativeStrengthIndexIndicator returns a derivative Indicator which returns the relative strength index of the base indicator
// in a given time frame. A more in-depth explanation of relative strength index can be found here:
// https://www.investopedia.com/terms/r/rsi.asp
func NewRelativeStrengthIndexIndicator(indicator Indicator, timeframe int) Indicator {
	return relativeStrengthIndexIndicator{
		rsIndicator: NewRelativeStrengthIndicator(indicator, timeframe),
		oneHundred:  decimal.NewFromInt(100),
	}
}

func (rsi relativeStrengthIndexIndicator) Calculate(index int) decimal.Decimal {
	relativeStrength := rsi.rsIndicator.Calculate(index)

	return rsi.oneHundred.Sub(rsi.oneHundred.Div(decimal.NewFromInt(1).Add(relativeStrength)))
}

type relativeStrengthIndicator struct {
	avgGain Indicator
	avgLoss Indicator
	window  int
}

// NewRelativeStrengthIndicator returns a derivative Indicator which returns the relative strength of the base indicator
// in a given time frame. Relative strength is the average again of up periods during the time frame divided by the
// average loss of down period during the same time frame
func NewRelativeStrengthIndicator(indicator Indicator, timeframe int) Indicator {
	return relativeStrengthIndicator{
		avgGain: NewMMAIndicator(NewGainIndicator(indicator), timeframe),
		avgLoss: NewMMAIndicator(NewLossIndicator(indicator), timeframe),
		window:  timeframe,
	}
}

func (rs relativeStrengthIndicator) Calculate(index int) decimal.Decimal {
	if index < rs.window-1 {
		return decimal.Zero
	}

	avgGain := rs.avgGain.Calculate(index)
	avgLoss := rs.avgLoss.Calculate(index)

	if avgLoss.Equal(decimal.Zero) {
		return decimal.NewFromFloat(math.MaxFloat64)
	}

	return avgGain.Div(avgLoss)
}
