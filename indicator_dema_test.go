package techan

import (
	"testing"
)

var demaMock = mockTimeSeriesFl(0.73, 0.72, 0.86, 0.72, 0.62, 0.76, 0.84, 0.69, 0.65, 0.71,
	0.53, 0.73, 0.77, 0.67, 0.68)

func TestDemaIndicator(t *testing.T) {
	// t.Run("DemaLine", func(t *testing.T) {
	// 	expectedValues := []float64{
	// 		0,
	// 		0,
	// 		0,
	// 		0,
	// 		0.73,
	// 		0.7244,
	// 		0.7992,
	// 		0.7858,
	// 		0.7374,
	// 		0.6884,
	// 		0.7184,
	// 		0.6939,
	// 		0.6859,
	// 	}

	// 	closeIndicator := NewCloseIndicator(demaMock)
	// 	indicatorEquals(t, expectedValues, NewDemaIndicator(closeIndicator, 5))
	// })

}
