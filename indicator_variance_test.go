package techan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVarianceIndicator(t *testing.T) {
	t.Run("when index is less than 1, returns 0", func(t *testing.T) {
		series := mockTimeSeries("0", "10")
		stdDev := NewStandardDeviationIndicator(NewCloseIndicator(series))

		assert.EqualValues(t, "0", stdDev.Calculate(0).String())
	})

	t.Run("returns the standard deviation when index > 2", func(t *testing.T) {
		series := mockTimeSeriesFl(
			10,
			2,
			38,
			23,
			38,
			23,
			21)

		varInd := NewVarianceIndicator(NewCloseIndicator(series))

		assert.EqualValues(t, "16.00", varInd.Calculate(1).StringFixed(2))
		assert.EqualValues(t, "238.22", varInd.Calculate(2).StringFixed(2))
		assert.EqualValues(t, "186.19", varInd.Calculate(3).StringFixed(2))
		assert.EqualValues(t, "211.36", varInd.Calculate(4).StringFixed(2))
		assert.EqualValues(t, "176.22", varInd.Calculate(5).StringFixed(2))
		assert.EqualValues(t, "151.27", varInd.Calculate(6).StringFixed(2))
	})
}
