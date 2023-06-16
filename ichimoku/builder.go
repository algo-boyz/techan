package ichimoku

import (
	"github.com/algo-boyz/decimal"
	"github.com/algo-boyz/techan"
)

func BuildIchimokuStatus(bars []*techan.Candle) (*IchimokuStatus, error) {
	if len(bars) == 0 {
		return nil, ErrDataNotFill
	}

	if len(bars) < 52 {
		return nil, ErrNotEnoughData
	}

	tenkenLine := calcLine(Line_Tenkan_sen, bars)
	kijonLine := calcLine(Line_kijon_sen, bars)

	span_a := calculate_span_a(tenkenLine, kijonLine)
	span_b := calcLine(Line_spanPeriod, bars)
	chiko_index := (len(bars) - int(Line_chikoPeriod)) - 1
	cheko_span := bars[chiko_index]

	var latestPrice *techan.Candle
	latestPriceIndex := (len(bars) - 1)
	if (len(bars) - 1) >= latestPriceIndex {
		latestPrice = bars[latestPriceIndex]
	}

	if !tenkenLine.isNil && !kijonLine.isNil && !span_a.isNil && !span_b.isNil {
		ichi := NewIchimokuStatus(tenkenLine, kijonLine, span_a, span_b, cheko_span, latestPrice)
		return ichi, nil
	}

	return nil, ErrBuildFailed
}

func calcLine(line_type ELine, bars []*techan.Candle) ValueLine {
	high := NewValueLineNil()
	low := NewValueLineNil()
	l := len(bars)
	from := l - 1 - int(line_type)
	if from == -1 {
		from = 0
	}
	bars_tmp := bars[from : l-1]
	if len(bars) < int(line_type) {
		return NewValueLineNil()
	}
	for _, v := range bars_tmp {

		if high.isNil {
			high.SetValue(v.High)
		}
		if low.isNil {
			low.SetValue(v.Low)
		}

		if v.High.GreaterThan(high.valLine) {
			high.SetValue(v.High)
		}

		if v.Low.LessThan(low.valLine) {
			low.SetValue(v.Low)
		}

	}
	line := low.valLine.Add(high.valLine).Div(decimal.NewFromInt(2))
	return NewValue(line)
}

func calculate_span_a(tenken ValueLine, kijon ValueLine) ValueLine {
	if !tenken.isNil && !kijon.isNil {
		v := tenken.valLine.Add(kijon.valLine).Div(decimal.NewFromInt(2))
		return NewValue(v)
	}

	return NewValueLineNil()
}
