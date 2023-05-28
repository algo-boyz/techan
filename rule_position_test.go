package techan

import (
	"testing"

	"github.com/algo-boyz/decimal"
	"github.com/stretchr/testify/assert"
)

func TestPositionNewRule(t *testing.T) {
	t.Run("returns true when position new", func(t *testing.T) {
		record := NewTradingRecord()
		rule := PositionNewRule{}

		assert.True(t, rule.IsSatisfied(0, record))
	})

	t.Run("returns false when position open", func(t *testing.T) {
		record := NewTradingRecord()

		record.Operate(Order{
			Side:   BUY,
			Amount: decimal.NewFromInt(1),
			Price:  decimal.NewFromInt(1),
		})

		rule := PositionNewRule{}

		assert.False(t, rule.IsSatisfied(0, record))
	})
}

func TestPositionOpenRule(t *testing.T) {
	t.Run("returns false when position new", func(t *testing.T) {
		record := NewTradingRecord()

		rule := PositionOpenRule{}

		assert.False(t, rule.IsSatisfied(0, record))
	})

	t.Run("returns true when position open", func(t *testing.T) {
		record := NewTradingRecord()

		record.Operate(Order{
			Side:   BUY,
			Amount: decimal.NewFromInt(1),
			Price:  decimal.NewFromInt(1),
		})

		rule := PositionOpenRule{}

		assert.True(t, rule.IsSatisfied(0, record))
	})
}
