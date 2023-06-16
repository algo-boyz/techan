package ichimoku

import (
	"fmt"

	"github.com/algo-boyz/decimal"
	"github.com/algo-boyz/techan"
	"github.com/logrusorgru/aurora"
)

type IchimokuStatus struct {
	TenkenSen ValueLine

	//_______________

	KijonSen ValueLine

	//in the future
	SenKoA_Shifted26 ValueLine

	//in the future
	SenKoB_Shifted26 ValueLine

	//extract value sen A & B from 26 candle past (26 shift forward in calc ichimoku)
	//SencoA 26 candle in the past (for check)
	SencoA_Past Point
	//extract value sen A & B from 26 candle past (26 shift forward in calc ichimoku)
	//SencoB 26 candle in the past (for check)
	SencoB_Past Point

	ChikoSpan *techan.Candle //close bar

	bar *techan.Candle
	//-----
	Status         EIchimokuStatus
	cloudSwitching bool

	line_helper lineHelper
}

func NewIchimokuStatus(tenken ValueLine, kijon ValueLine, senKoA_Shifted26 ValueLine, senKoB_Shifted52 ValueLine, chiko_span, bar *techan.Candle) *IchimokuStatus {

	o := IchimokuStatus{}

	o.TenkenSen = tenken

	o.KijonSen = kijon

	o.SenKoA_Shifted26 = senKoA_Shifted26
	o.SenKoB_Shifted26 = senKoB_Shifted52

	o.ChikoSpan = chiko_span
	o.bar = bar
	o.Status = IchimokuStatus_NAN
	o.line_helper = NewLineHelper()

	return &o
}

func (o *IchimokuStatus) SetChikoSpan(v *techan.Candle) {
	o.ChikoSpan = v
}

func (o *IchimokuStatus) Set_SenCo_A_Past(p Point) {
	o.SencoA_Past = p
}

func (o *IchimokuStatus) Set_SenCo_B_Past(p Point) {
	o.SencoB_Past = p
}

func (o *IchimokuStatus) SetStatus(status EIchimokuStatus) {
	o.Status = status
}

func (o *IchimokuStatus) GetStatus() EIchimokuStatus {
	return o.Status
}

func (o *IchimokuStatus) SetCloudSwitching(v bool) {
	o.cloudSwitching = v
}

func (o *IchimokuStatus) GetCloudSwitching() bool {
	return o.cloudSwitching
}

func (o *IchimokuStatus) Is_cloud_green() bool {
	return o.SenKoA_Shifted26.valLine.GreaterThan(o.SenKoB_Shifted26.valLine)
}

func (o *IchimokuStatus) IsChikoAbovePrice() bool {
	return o.ChikoSpan.High.GreaterThan(o.bar.Close)
}

func (o *IchimokuStatus) CloudStatus(intersection decimal.Decimal) EIchimokuStatus {
	if o.SenKoA_Shifted26.isNil || o.SenKoB_Shifted26.isNil {
		return IchimokuStatus_NAN
	}
	if o.SencoA_Past.isNil || o.SencoB_Past.isNil {
		return IchimokuStatus_NAN
	}

	sen_B := o.SencoB_Past //Senko B in_26_candle_pass
	sen_A := o.SencoA_Past //Senko A in_26_candle_pass
	if sen_A.Y.GreaterThan(intersection) && sen_B.Y.GreaterThan(intersection) {
		return IchimokuStatus_Cross_Below
	} else if sen_A.Y.LessThan(intersection) && sen_B.Y.LessThan(intersection) {
		return IchimokuStatus_Cross_Above
	} else if sen_A.Y.LessThan(intersection) && sen_B.Y.GreaterThan(intersection) || sen_A.Y.GreaterThan(intersection) && sen_B.Y.LessThan(intersection) {
		return IchimokuStatus_Cross_Inside
	}

	return IchimokuStatus_NAN

}

func (o *IchimokuStatus) GetStatusString() string {
	result := ""
	switch o.Status {
	case IchimokuStatus_NAN:
		result = "nan"

	case IchimokuStatus_Cross_Below:
		result = "cross below"
	case IchimokuStatus_Cross_Above:
		result = "cross above"
	case IchimokuStatus_Cross_Inside:
		result = "cross inside"
	}

	return result
}

func (o *IchimokuStatus) Print() string {
	return fmt.Sprintf("ichimoku cloud %v|%v|%v|%v|%v| Green: %v, Chiko UP: %v | status: %v | %v\n",
		aurora.Green(o.TenkenSen.Value()),
		aurora.Green(o.KijonSen.Value()),
		aurora.Green(o.SenKoA_Shifted26.Value()),
		aurora.Green(o.SenKoB_Shifted26.Value()),
		aurora.Green(o.ChikoSpan.Close),
		aurora.Red(o.Is_cloud_green()),
		aurora.Red(o.IsChikoAbovePrice()),
		o.GetStatusString(),
		o.bar.Period.Start,
	)
}
