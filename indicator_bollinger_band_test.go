package techan

import (
	"strings"
	"testing"

	"github.com/algo-boyz/decimal"
	"github.com/stretchr/testify/assert"
)

func decimalAlmostEquals(t *testing.T, expected, actual decimal.Decimal, epsilon float64) {
	assert.InEpsilon(t, expected.InexactFloat64(), actual.InexactFloat64(), epsilon)
}

// number data from https://school.stockcharts.com/doku.php?id=technical_indicators:bollinger_bands
func TestBollingerBandIndicator(t *testing.T) {
	ts := mockTimeSeriesFl(
		86.16,
		89.09,
		88.78,
		90.32,
		89.07,
		91.15,
		89.44,
		89.18,
		86.93,
		87.68,
		86.96,
		89.43,
		89.32,
		88.72,
		87.45,
		87.26,
		89.50,
		87.90,
		89.13,

		90.70,
		92.90,
		92.98,
		91.80,
		92.66,
		92.68,
		92.30,
		92.77,
		92.54,
		92.95,
		93.20,
		91.07,
		89.83,
		89.74,
		90.40,
		90.74,
		88.02,
		88.09,
		88.84,
		90.78,
		90.54,
		91.39,
		90.65,
	)

	expectedLines := strings.Split(`
88.71	1.29	91.29	86.12	5.17
89.05	1.45	91.95	86.14	5.81
89.24	1.69	92.61	85.87	6.75
89.39	1.77	92.93	85.85	7.09
89.51	1.90	93.31	85.70	7.61
89.69	2.02	93.73	85.65	8.08
89.75	2.08	93.90	85.59	8.31
89.91	2.18	94.27	85.56	8.71
90.08	2.24	94.57	85.60	8.97
90.38	2.20	94.79	85.98	8.81
90.66	2.19	95.04	86.27	8.77
90.86	2.02	94.91	86.82	8.09
90.88	2.01	94.90	86.87	8.04
90.91	2.00	94.90	86.91	7.98
90.99	1.94	94.86	87.12	7.74
91.15	1.76	94.67	87.63	7.04
91.19	1.68	94.56	87.83	6.73
91.12	1.78	94.68	87.56	7.12
91.17	1.70	94.58	87.76	6.82
91.25	1.64	94.53	87.97	6.57
91.24	1.65	94.53	87.95	6.58
91.17	1.60	94.37	87.96	6.41
91.05	1.55	94.15	87.95	6.20`, "\n")

	SMAs := []string{}
	STDEVs := []string{}
	BBUPs := []string{}
	BBLOs := []string{}
	BBWs := []string{}
	for _, line := range expectedLines {
		tokens := strings.Split(line, "\t")
		if len(tokens) == 5 {
			SMAs = append(SMAs, tokens[0])
			STDEVs = append(STDEVs, tokens[1])
			BBUPs = append(BBUPs, tokens[2])
			BBLOs = append(BBLOs, tokens[3])
			BBWs = append(BBWs, tokens[4])
		}
	}
	assert.Equal(t, len(SMAs), 23)

	window := 20
	sigma := 2.0
	src := NewCloseIndicator(ts)
	sma := NewSimpleMovingAverage(src, window)
	wstd := NewWindowedStandardDeviationIndicator(src, window)
	bbUP := NewBollingerUpperBandIndicator(src, window, sigma)
	bbLO := NewBollingerLowerBandIndicator(src, window, sigma)

	for i := window - 1; i < len(ts.Candles); i++ {
		j := i - (window - 1)
		smaj, _ := decimal.NewFromString(SMAs[j])
		stdevj, _ := decimal.NewFromString(STDEVs[j])
		bbupj, _ := decimal.NewFromString(BBUPs[j])
		bbloj, _ := decimal.NewFromString(BBLOs[j])
		bbwj, _ := decimal.NewFromString(BBWs[j])
		decimalAlmostEquals(t, smaj, sma.Calculate(i), 0.01)
		decimalAlmostEquals(t, stdevj, wstd.Calculate(i), 0.01)
		decimalAlmostEquals(t, bbupj, bbUP.Calculate(i), 0.01)
		decimalAlmostEquals(t, bbloj, bbLO.Calculate(i), 0.01)
		decimalAlmostEquals(t, bbwj, bbUP.Calculate(i).Sub(bbLO.Calculate((i))), 0.01)
	}
}
