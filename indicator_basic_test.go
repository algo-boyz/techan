package techan

import (
	"testing"
	"time"

	"github.com/algo-boyz/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewVolumeIndicator(t *testing.T) {
	assert.NotNil(t, NewVolumeIndicator(NewTimeSeries()))
}

func TestVolumeIndicator_Calculate(t *testing.T) {
	series := NewTimeSeries()

	candle := NewCandle(TimePeriod{
		Start: time.Now(),
		End:   time.Now().Add(time.Minute),
	})
	candle.Volume = decimal.NewFromFloat(1.2080)

	series.AddCandle(candle)

	indicator := NewVolumeIndicator(series)
	assert.EqualValues(t, "1.208", indicator.Calculate(0).StringFixed(3))
}

func TestTypicalPriceIndicator_Calculate(t *testing.T) {
	series := NewTimeSeries()

	candle := NewCandle(TimePeriod{
		Start: time.Now(),
		End:   time.Now().Add(time.Minute),
	})
	candle.Low = decimal.NewFromFloat(1.2080)
	candle.High = decimal.NewFromFloat(1.22)
	candle.Close = decimal.NewFromFloat(1.215)

	series.AddCandle(candle)

	typicalPrice := NewTypicalPriceIndicator(series).Calculate(0)

	assert.EqualValues(t, "1.2143", typicalPrice.StringFixed(4))
}
