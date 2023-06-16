package tests

import (
	"testing"
	"time"

	"github.com/algo-boyz/techan"
	"github.com/stretchr/testify/assert"

	"github.com/algo-boyz/techan/ichimoku"
)

var (
	interval = time.Minute * 30
	//52 bar
	bar_h1 = []*techan.Candle{
		{Low: d(8110), High: d(8180), Close: d(8160), Open: d(8110), Volume: d(664372.00), Period: p(1667201400000)},
		{Low: d(8100), High: d(8260), Close: d(8200), Open: d(8150), Volume: d(1241301.00), Period: p(1667205000000)},
		{Low: d(8110), High: d(8450), Close: d(8440), Open: d(8170), Volume: d(2909458.00), Period: p(1667280600000)},
		{Low: d(8310), High: d(8450), Close: d(8360), Open: d(8440), Volume: d(778238.00), Period: p(1667284200000)},
		{Low: d(8240), High: d(8370), Close: d(8260), Open: d(8360), Volume: d(658420.00), Period: p(1667287800000)},
		{Low: d(8240), High: d(8450), Close: d(8440), Open: d(8260), Volume: d(1814124.00), Period: p(1667291400000)},
		{Low: d(8270), High: d(8440), Close: d(8300), Open: d(8440), Volume: d(1267103.00), Period: p(1667367000000)},
		{Low: d(8270), High: d(8510), Close: d(8510), Open: d(8300), Volume: d(1821017.00), Period: p(1667370600000)},
		{Low: d(8430), High: d(8540), Close: d(8440), Open: d(8510), Volume: d(559250.00), Period: p(1667374200000)},
		{Low: d(8420), High: d(8470), Close: d(8440), Open: d(8440), Volume: d(544851.00), Period: p(1667377800000)},
		{Low: d(8480), High: d(8730), Close: d(8730), Open: d(8550), Volume: d(4284720.00), Period: p(1667626200000)},
		{Low: d(8730), High: d(8730), Close: d(8730), Open: d(8730), Volume: d(1382828.00), Period: p(1667629800000)},
		{Low: d(8730), High: d(8730), Close: d(8730), Open: d(8730), Volume: d(1678201.00), Period: p(1667633400000)},
		{Low: d(8730), High: d(8730), Close: d(8730), Open: d(8730), Volume: d(549277.00), Period: p(1667637000000)},
		{Low: d(8800), High: d(9070), Close: d(9060), Open: d(8800), Volume: d(5342062.00), Period: p(1667712600000)},
		{Low: d(9040), High: d(9070), Close: d(9070), Open: d(9060), Volume: d(8126959.00), Period: p(1667716200000)},
		{Low: d(9070), High: d(9070), Close: d(9070), Open: d(9070), Volume: d(527101.00), Period: p(1667719800000)},
		{Low: d(9070), High: d(9070), Close: d(9070), Open: d(9070), Volume: d(702521.00), Period: p(1667723400000)},
		{Low: d(9160), High: d(9440), Close: d(9430), Open: d(9290), Volume: d(4409696.00), Period: p(1667799000000)},
		{Low: d(9410), High: d(9490), Close: d(9490), Open: d(9420), Volume: d(7522839.00), Period: p(1667802600000)},
		{Low: d(9490), High: d(9490), Close: d(9490), Open: d(9490), Volume: d(777299.00), Period: p(1667806200000)},
		{Low: d(9490), High: d(9490), Close: d(9490), Open: d(9490), Volume: d(405416.00), Period: p(1667809800000)},
		{Low: d(9300), High: d(9890), Close: d(9530), Open: d(9890), Volume: d(7097789.00), Period: p(1667885400000)},
		{Low: d(9460), High: d(9570), Close: d(9470), Open: d(9520), Volume: d(3033312.00), Period: p(1667889000000)},
		{Low: d(9380), High: d(9490), Close: d(9410), Open: d(9470), Volume: d(2714433.00), Period: p(1667892600000)},
		{Low: d(9390), High: d(9490), Close: d(9450), Open: d(9420), Volume: d(3876877.00), Period: p(1667896200000)},
		{Low: d(9250), High: d(9540), Close: d(9410), Open: d(9350), Volume: d(3448605.00), Period: p(1667971800000)},
		{Low: d(9400), High: d(9840), Close: d(9800), Open: d(9410), Volume: d(6547559.00), Period: p(1667975400000)},
		{Low: d(9640), High: d(9830), Close: d(9650), Open: d(9800), Volume: d(2416825.00), Period: p(1667979000000)},
		{Low: d(9650), High: d(9860), Close: d(9680), Open: d(9700), Volume: d(2463503.00), Period: p(1667982600000)},
		{Low: d(9640), High: d(9870), Close: d(9800), Open: d(9750), Volume: d(2000789.00), Period: p(1668231000000)},
		{Low: d(9520), High: d(9800), Close: d(9520), Open: d(9780), Volume: d(3214849.00), Period: p(1668234600000)},
		{Low: d(9520), High: d(9680), Close: d(9620), Open: d(9550), Volume: d(3019512.00), Period: p(1668238200000)},
		{Low: d(9610), High: d(9810), Close: d(9740), Open: d(9640), Volume: d(2473212.00), Period: p(1668241800000)},
		{Low: d(9450), High: d(9710), Close: d(9530), Open: d(9710), Volume: d(1455003.00), Period: p(1668317400000)},
		{Low: d(9510), High: d(9700), Close: d(9700), Open: d(9520), Volume: d(1341450.00), Period: p(1668321000000)},
		{Low: d(9520), High: d(9720), Close: d(9650), Open: d(9700), Volume: d(2922575.00), Period: p(1668324600000)},
		{Low: d(9470), High: d(9650), Close: d(9470), Open: d(9650), Volume: d(907574.00), Period: p(1668328200000)},
		{Low: d(9250), High: d(9620), Close: d(9250), Open: d(9510), Volume: d(1573592.00), Period: p(1668403800000)},
		{Low: d(9220), High: d(9420), Close: d(9380), Open: d(9270), Volume: d(1372258.00), Period: p(1668407400000)},
		{Low: d(9340), High: d(9530), Close: d(9490), Open: d(9380), Volume: d(3147032.00), Period: p(1668411000000)},
		{Low: d(9370), High: d(9550), Close: d(9370), Open: d(9490), Volume: d(2153637.00), Period: p(1668414600000)},
		{Low: d(9380), High: d(9750), Close: d(9670), Open: d(9450), Volume: d(1861478.00), Period: p(1668490200000)},
		{Low: d(9580), High: d(9700), Close: d(9650), Open: d(9670), Volume: d(2890813.00), Period: p(1668493800000)},
		{Low: d(9610), High: d(9700), Close: d(9670), Open: d(9610), Volume: d(1288957.00), Period: p(1668497400000)},
		{Low: d(9630), High: d(9800), Close: d(9730), Open: d(9650), Volume: d(2413843.00), Period: p(1668501000000)},
		{Low: d(9580), High: d(9780), Close: d(9630), Open: d(9750), Volume: d(803830.00), Period: p(1668576600000)},
		{Low: d(9630), High: d(9720), Close: d(9670), Open: d(9650), Volume: d(699785.00), Period: p(1668580200000)},
		{Low: d(9640), High: d(9700), Close: d(9640), Open: d(9700), Volume: d(393592.00), Period: p(1668583800000)},
		{Low: d(9580), High: d(9660), Close: d(9630), Open: d(9640), Volume: d(1443871.00), Period: p(1668587400000)},
		{Low: d(9300), High: d(9600), Close: d(9370), Open: d(9510), Volume: d(3845936.00), Period: p(1668835800000)},
		{Low: d(9310), High: d(9380), Close: d(9330), Open: d(9380), Volume: d(1380628.00), Period: p(1668839400000)},
	}
)

func Test_26DayinThePastNotExist(t *testing.T) {

	series := techan.NewTimeSeries()
	series.Candles = bar_h1
	indicator := ichimoku.NewIndicator()

	indicator.MakeIchimokuInPast(series, 1)

	_, e1 := indicator.AnalyseTriggerCross(indicator.GetLast(), bar_h1)

	assert.Equal(t, e1, ichimoku.ErrChikoStatus26InPastNotMade)

}
func TestInside(t *testing.T) {

	indicator := ichimoku.NewIndicator()

	today := ichimoku.NewIchimokuStatus(
		ichimoku.NewValue(d(8705)),
		ichimoku.NewValue(d(8710)),
		ichimoku.NewValue(d(8707)),
		ichimoku.NewValue(d(8930)),
		&techan.Candle{},
		&techan.Candle{Low: d(8400), High: d(8460), Close: d(8440), Open: d(8440), Volume: d(906352), Period: p(1664699400000)})

	yesterday := ichimoku.NewIchimokuStatus(
		ichimoku.NewValue(d(8720)),
		ichimoku.NewValue(d(8710)),
		ichimoku.NewValue(d(8715)),
		ichimoku.NewValue(d(8940)),
		&techan.Candle{},
		&techan.Candle{Low: d(8430), High: d(8480), Close: d(8450), Open: d(8460), Volume: d(652416), Period: p(1664695800000)})

	lines_result := make([]*ichimoku.IchimokuStatus, 2)
	lines_result[0] = today //today
	lines_result[1] = yesterday

	a, e := indicator.PreAnalyseIchimoku(lines_result)
	assert.Empty(t, e)
	assert.Equal(t, a.Status, ichimoku.IchimokuStatus_Cross_Above)

}

func p(v int64) techan.TimePeriod {
	d := time.UnixMilli(v).UTC()
	return techan.NewTimePeriod(d, interval)
}
