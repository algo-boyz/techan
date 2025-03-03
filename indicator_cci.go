package techan

import "github.com/algo-boyz/decimal"

type commidityChannelIndexIndicator struct {
	series *TimeSeries
	window int
}

// NewCCIIndicator Returns a new Commodity Channel Index Indicator
// http://stockcharts.com/school/doku.php?id=chart_school:technical_indicators:commodity_channel_index_cci
func NewCCIIndicator(ts *TimeSeries, window int) Indicator {
	return commidityChannelIndexIndicator{
		series: ts,
		window: window,
	}
}

func (ccii commidityChannelIndexIndicator) Calculate(index int) decimal.Decimal {
	typicalPrice := NewTypicalPriceIndicator(ccii.series)
	typicalPriceSma := NewSimpleMovingAverage(typicalPrice, ccii.window)
	meanDeviation := NewMeanDeviationIndicator(NewCloseIndicator(ccii.series), ccii.window)

	return typicalPrice.Calculate(index).Sub(typicalPriceSma.Calculate(index)).Div(meanDeviation.Calculate(index).Mul(decimal.NewFromFloat(0.015)))
}
