package techan

import (
	"fmt"
	"strings"

	"github.com/algo-boyz/decimal"
)

// Candle represents basic market information for a security over a given time period
type Candle struct {
	Period     TimePeriod
	Open       decimal.Decimal
	Close      decimal.Decimal
	High       decimal.Decimal
	Low        decimal.Decimal
	Volume     decimal.Decimal
	TradeCount uint
}

// NewCandle returns a new *Candle for a given time period
func NewCandle(period TimePeriod) (c *Candle) {
	return &Candle{
		Period: period,
		Open:   decimal.Zero,
		Close:  decimal.Zero,
		High:   decimal.Zero,
		Low:    decimal.Zero,
		Volume: decimal.Zero,
	}
}

// AddTrade adds a trade to this candle. It will determine if the current price is higher or lower than the min or max
// price and increment the tradecount.
func (c *Candle) AddTrade(tradeAmount, tradePrice decimal.Decimal) {
	if c.Open.IsZero() {
		c.Open = tradePrice
	}
	c.Close = tradePrice

	if c.High.IsZero() {
		c.High = tradePrice
	} else if tradePrice.GreaterThan(c.High) {
		c.High = tradePrice
	}

	if c.Low.IsZero() {
		c.Low = tradePrice
	} else if tradePrice.LessThan(c.Low) {
		c.Low = tradePrice
	}

	if c.Volume.IsZero() {
		c.Volume = tradeAmount
	} else {
		c.Volume = c.Volume.Add(tradeAmount)
	}

	c.TradeCount++
}

func (c *Candle) String() string {
	return strings.TrimSpace(fmt.Sprintf(
		`
Time:	%s
Open:	%s
Close:	%s
High:	%s
Low:	%s
Volume:	%s
	`,
		c.Period,
		c.Open.StringFixed(2),
		c.Close.StringFixed(2),
		c.High.StringFixed(2),
		c.Low.StringFixed(2),
		c.Volume.StringFixed(2),
	))
}
