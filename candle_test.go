package techan

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/algo-boyz/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCandle_AddTrade(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	})

	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(2)) // Open
	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(5)) // High
	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(1)) // Low
	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(3)) // No Diff
	candle.AddTrade(decimal.NewFromInt(1), decimal.NewFromInt(3)) // Close

	assert.EqualValues(t, 2, candle.Open.InexactFloat64())
	assert.EqualValues(t, 5, candle.High.InexactFloat64())
	assert.EqualValues(t, 1, candle.Low.InexactFloat64())
	assert.EqualValues(t, 3, candle.Close.InexactFloat64())
	assert.EqualValues(t, 5, candle.Volume.InexactFloat64())
	assert.EqualValues(t, 5, candle.TradeCount)
}

func TestCandle_String(t *testing.T) {
	now := time.Now()
	candle := NewCandle(TimePeriod{
		Start: now,
		End:   now.Add(time.Minute),
	})

	candle.Close, _ = decimal.NewFromString("1")
	candle.Open, _ = decimal.NewFromString("2")
	candle.High, _ = decimal.NewFromString("3")
	candle.Low, _ = decimal.NewFromString("0")
	candle.Volume, _ = decimal.NewFromString("10")

	expected := strings.TrimSpace(fmt.Sprintf(`
Time:	%s
Open:	2.00
Close:	1.00
High:	3.00
Low:	0.00
Volume:	10.00
`, candle.Period))

	assert.EqualValues(t, expected, candle.String())
}
