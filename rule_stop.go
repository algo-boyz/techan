package techan

import "github.com/algo-boyz/decimal"

type stopLossRule struct {
	Indicator
	tolerance decimal.Decimal
}

// NewStopLossRule returns a new rule that is satisfied when the given loss tolerance (a percentage) is met or exceeded.
// Loss tolerance should be a value between -1 and 1.
func NewStopLossRule(series *TimeSeries, lossTolerance float64) Rule {
	return stopLossRule{
		Indicator: NewCloseIndicator(series),
		tolerance: decimal.NewFromFloat(lossTolerance),
	}
}

func (slr stopLossRule) IsSatisfied(index int, record *TradingRecord) bool {
	if !record.CurrentPosition().IsOpen() {
		return false
	}

	Open := record.CurrentPosition().CostBasis()
	loss := slr.Indicator.Calculate(index).Div(Open).Sub(decimal.NewFromInt(1))
	return loss.LessThanOrEqual(slr.tolerance)
}
