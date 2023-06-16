package techan

import "github.com/algo-boyz/decimal"

var mul = decimal.NewFromInt(2)

type demaIndicator struct {
	ema  Indicator
	ema2 Indicator
}

func NewDemaIndicator(indicator Indicator, window int) Indicator {
	ema := NewEMAIndicator(indicator, window)
	dema := demaIndicator{
		ema:  ema,
		ema2: NewEMAIndicator(ema, window),
	}

	return dema
}

// DEMA = (2 * EMA(values)) - EMA(EMA(values))
func (dema demaIndicator) Calculate(index int) decimal.Decimal {
	return mul.Mul(dema.ema.Calculate(index)).Sub(dema.ema2.Calculate(index))
}
