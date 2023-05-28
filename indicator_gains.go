package techan

import "github.com/algo-boyz/decimal"

type gainLossIndicator struct {
	Indicator
	coefficient decimal.Decimal
}

// NewGainIndicator returns a derivative indicator that returns the gains in the underlying indicator in the last bar,
// if any. If the delta is negative, zero is returned
func NewGainIndicator(indicator Indicator) Indicator {
	return gainLossIndicator{
		Indicator:   indicator,
		coefficient: decimal.NewFromInt(1),
	}
}

// NewLossIndicator returns a derivative indicator that returns the losses in the underlying indicator in the last bar,
// if any. If the delta is positive, zero is returned
func NewLossIndicator(indicator Indicator) Indicator {
	return gainLossIndicator{
		Indicator:   indicator,
		coefficient: decimal.NewFromInt(1).Neg(),
	}
}

func (gli gainLossIndicator) Calculate(index int) decimal.Decimal {
	if index == 0 {
		return decimal.Zero
	}

	delta := gli.Indicator.Calculate(index).Sub(gli.Indicator.Calculate(index - 1)).Mul(gli.coefficient)
	if delta.GreaterThan(decimal.Zero) {
		return delta
	}

	return decimal.Zero
}

type cumulativeIndicator struct {
	Indicator
	window int
	mult   decimal.Decimal
}

// NewCumulativeGainsIndicator returns a derivative indicator which returns all gains made in a base indicator for a given
// window.
func NewCumulativeGainsIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      decimal.NewFromInt(1),
	}
}

// NewCumulativeLossesIndicator returns a derivative indicator which returns all losses in a base indicator for a given
// window.
func NewCumulativeLossesIndicator(indicator Indicator, window int) Indicator {
	return cumulativeIndicator{
		Indicator: indicator,
		window:    window,
		mult:      decimal.NewFromInt(1).Neg(),
	}
}

func (ci cumulativeIndicator) Calculate(index int) decimal.Decimal {
	total := decimal.NewFromInt(0.0)

	for i := Max(1, index-(ci.window-1)); i <= index; i++ {
		diff := ci.Indicator.Calculate(i).Sub(ci.Indicator.Calculate(i - 1))
		if diff.Mul(ci.mult).GreaterThan(decimal.Zero) {
			total = total.Add(diff.Abs())
		}
	}

	return total
}

type percentChangeIndicator struct {
	Indicator
}

// NewPercentChangeIndicator returns a derivative indicator which returns the percent change (positive or negative)
// made in a base indicator up until the given indicator
func NewPercentChangeIndicator(indicator Indicator) Indicator {
	return percentChangeIndicator{indicator}
}

func (pgi percentChangeIndicator) Calculate(index int) decimal.Decimal {
	if index == 0 {
		return decimal.Zero
	}

	cp := pgi.Indicator.Calculate(index)
	cplast := pgi.Indicator.Calculate(index - 1)
	return cp.Div(cplast).Sub(decimal.NewFromInt(1))
}
