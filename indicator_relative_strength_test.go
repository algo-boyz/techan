package techan

import (
	"math"
	"testing"

	"github.com/algo-boyz/decimal"
	"github.com/stretchr/testify/assert"
)

func TestRelativeStrengthIndexIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndexIndicator(NewCloseIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{
		0,
		0,
		0,
		0,
		0,
		0,
		57.9952,
		54.0751,
		21.451,
		44.7739,
		14.1542,
		21.2794,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndicator(t *testing.T) {
	indicator := NewRelativeStrengthIndicator(NewCloseIndicator(mockedTimeSeries), 3)

	expectedValues := []float64{
		0,
		0,
		0,
		0,
		0,
		0,
		1.3807,
		1.1775,
		0.2731,
		0.8107,
		0.1649,
		0.2703,
	}

	indicatorEquals(t, expectedValues, indicator)
}

func TestRelativeStrengthIndicatorNoPriceChange(t *testing.T) {
	close := NewCloseIndicator(mockTimeSeries("42.0", "42.0"))
	rsInd := NewRelativeStrengthIndicator(close, 2)
	assert.Equal(t, decimal.NewFromFloat(math.MaxFloat64).StringFixed(2), rsInd.Calculate(1).StringFixed(2))
}
