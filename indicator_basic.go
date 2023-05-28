package techan

import "github.com/algo-boyz/decimal"

type volumeIndicator struct {
	*TimeSeries
}

// NewVolumeIndicator returns an indicator which returns the volume of a candle for a given index
func NewVolumeIndicator(series *TimeSeries) Indicator {
	return volumeIndicator{series}
}

func (vi volumeIndicator) Calculate(index int) decimal.Decimal {
	return vi.Candles[index].Volume
}

type CloseIndicator struct {
	*TimeSeries
}

// NewCloseIndicator returns an Indicator which returns the close price of a candle for a given index
func NewCloseIndicator(series *TimeSeries) Indicator {
	return CloseIndicator{series}
}

func (cpi CloseIndicator) Calculate(index int) decimal.Decimal {
	return cpi.Candles[index].Close
}

type highPriceIndicator struct {
	*TimeSeries
}

// NewHighPriceIndicator returns an Indicator which returns the high price of a candle for a given index
func NewHighPriceIndicator(series *TimeSeries) Indicator {
	return highPriceIndicator{
		series,
	}
}

func (hpi highPriceIndicator) Calculate(index int) decimal.Decimal {
	return hpi.Candles[index].High
}

type lowPriceIndicator struct {
	*TimeSeries
}

// NewLowPriceIndicator returns an Indicator which returns the low price of a candle for a given index
func NewLowPriceIndicator(series *TimeSeries) Indicator {
	return lowPriceIndicator{
		series,
	}
}

func (lpi lowPriceIndicator) Calculate(index int) decimal.Decimal {
	return lpi.Candles[index].Low
}

type OpenIndicator struct {
	*TimeSeries
}

// NewOpenIndicator returns an Indicator which returns the open price of a candle for a given index
func NewOpenIndicator(series *TimeSeries) Indicator {
	return OpenIndicator{
		series,
	}
}

func (opi OpenIndicator) Calculate(index int) decimal.Decimal {
	return opi.Candles[index].Open
}

type typicalPriceIndicator struct {
	*TimeSeries
}

// NewTypicalPriceIndicator returns an Indicator which returns the typical price of a candle for a given index.
// The typical price is an average of the high, low, and close prices for a given candle.
func NewTypicalPriceIndicator(series *TimeSeries) Indicator {
	return typicalPriceIndicator{series}
}

func (tpi typicalPriceIndicator) Calculate(index int) decimal.Decimal {
	numerator := tpi.Candles[index].High.Add(tpi.Candles[index].Low).Add(tpi.Candles[index].Close)
	return numerator.Div(decimal.NewFromInt(3))
}
