package main

import (
	"fmt"
	"time"

	"github.com/algo-boyz/decimal"
	"github.com/algo-boyz/techan"

	"github.com/algo-boyz/techan/ichimoku"
)

var interval = time.Minute * 30

func d(v float64) decimal.Decimal {
	return decimal.NewFromFloat(v)
}

func p(v int64) techan.TimePeriod {
	d := time.UnixMilli(v).UTC()
	return techan.NewTimePeriod(d, interval)
}

func main() {

	// sort bar_h1 [old date to new date] ascending
	bar_h1 := []*techan.Candle{
		{Low: d(9450.00), High: d(9570.00), Close: d(9490.00), Open: d(9530.00), Volume: d(1158096.00), Period: p(1662525000000)},
		{Low: d(9450.00), High: d(9550.00), Close: d(9480.00), Open: d(9530.00), Volume: d(1041077.00), Period: p(1662528600000)},
		{Low: d(9420.00), High: d(9520.00), Close: d(9500.00), Open: d(9500.00), Volume: d(1966815.00), Period: p(1662532200000)},
		{Low: d(9480.00), High: d(9520.00), Close: d(9510.00), Open: d(9480.00), Volume: d(748430.00), Period: p(1662535800000)},
		{Low: d(9250.00), High: d(9500.00), Close: d(9390.00), Open: d(9500.00), Volume: d(1961310.00), Period: p(1662784200000)},
		{Low: d(9340.00), High: d(9570.00), Close: d(9520.00), Open: d(9350.00), Volume: d(1881148.00), Period: p(1662787800000)},
		{Low: d(9300.00), High: d(9520.00), Close: d(9320.00), Open: d(9520.00), Volume: d(2174744.00), Period: p(1662791400000)},
		{Low: d(9320.00), High: d(9440.00), Close: d(9440.00), Open: d(9320.00), Volume: d(428193.00), Period: p(1662795000000)},
		{Low: d(9280.00), High: d(9500.00), Close: d(9470.00), Open: d(9440.00), Volume: d(1174617.00), Period: p(1662870600000)},
		{Low: d(9300.00), High: d(9490.00), Close: d(9330.00), Open: d(9480.00), Volume: d(564590.00), Period: p(1662874200000)},
		{Low: d(9260.00), High: d(9360.00), Close: d(9300.00), Open: d(9330.00), Volume: d(1325195.00), Period: p(1662877800000)},
		{Low: d(9280.00), High: d(9360.00), Close: d(9360.00), Open: d(9290.00), Volume: d(725716.00), Period: p(1662881400000)},
		{Low: d(9290.00), High: d(9470.00), Close: d(9310.00), Open: d(9310.00), Volume: d(1156762.00), Period: p(1662957000000)},
		{Low: d(9210.00), High: d(9330.00), Close: d(9250.00), Open: d(9310.00), Volume: d(1463096.00), Period: p(1662960600000)},
		{Low: d(9240.00), High: d(9440.00), Close: d(9380.00), Open: d(9240.00), Volume: d(1350073.00), Period: p(1662964200000)},
		{Low: d(9360.00), High: d(9400.00), Close: d(9390.00), Open: d(9380.00), Volume: d(526677.00), Period: p(1662967800000)},
		{Low: d(9260.00), High: d(9450.00), Close: d(9280.00), Open: d(9300.00), Volume: d(685445.00), Period: p(1663043400000)},
		{Low: d(9270.00), High: d(9450.00), Close: d(9440.00), Open: d(9280.00), Volume: d(957947.00), Period: p(1663047000000)},
		{Low: d(9340.00), High: d(9440.00), Close: d(9340.00), Open: d(9440.00), Volume: d(716902.00), Period: p(1663050600000)},
		{Low: d(9300.00), High: d(9380.00), Close: d(9380.00), Open: d(9350.00), Volume: d(860525.00), Period: p(1663054200000)},
		{Low: d(9230.00), High: d(9360.00), Close: d(9250.00), Open: d(9300.00), Volume: d(1844161.00), Period: p(1663129800000)},
		{Low: d(9230.00), High: d(9290.00), Close: d(9240.00), Open: d(9250.00), Volume: d(876713.00), Period: p(1663133400000)},
		{Low: d(9210.00), High: d(9280.00), Close: d(9260.00), Open: d(9240.00), Volume: d(909086.00), Period: p(1663137000000)},
		{Low: d(9200.00), High: d(9270.00), Close: d(9240.00), Open: d(9260.00), Volume: d(1347750.00), Period: p(1663140600000)},
		{Low: d(9060.00), High: d(9240.00), Close: d(9100.00), Open: d(9240.00), Volume: d(1513067.00), Period: p(1663475400000)},
		{Low: d(9060.00), High: d(9150.00), Close: d(9080.00), Open: d(9080.00), Volume: d(1770320.00), Period: p(1663479000000)},
		{Low: d(8950.00), High: d(9100.00), Close: d(8970.00), Open: d(9080.00), Volume: d(2172300.00), Period: p(1663482600000)},
		{Low: d(8950.00), High: d(9120.00), Close: d(9070.00), Open: d(8980.00), Volume: d(2417598.00), Period: p(1663486200000)},
		{Low: d(8960.00), High: d(9160.00), Close: d(9040.00), Open: d(8960.00), Volume: d(1154404.00), Period: p(1663561800000)},
		{Low: d(8980.00), High: d(9070.00), Close: d(8990.00), Open: d(9040.00), Volume: d(544846.00), Period: p(1663565400000)},
		{Low: d(8930.00), High: d(9020.00), Close: d(8960.00), Open: d(8990.00), Volume: d(901781.00), Period: p(1663569000000)},
		{Low: d(8910.00), High: d(8970.00), Close: d(8920.00), Open: d(8970.00), Volume: d(1021775.00), Period: p(1663572600000)},
		{Low: d(8900.00), High: d(9060.00), Close: d(8960.00), Open: d(8900.00), Volume: d(1163197.00), Period: p(1663648200000)},
		{Low: d(8860.00), High: d(9000.00), Close: d(8870.00), Open: d(8950.00), Volume: d(932584.00), Period: p(1663651800000)},
		{Low: d(8820.00), High: d(8880.00), Close: d(8830.00), Open: d(8880.00), Volume: d(1458995.00), Period: p(1663655400000)},
		{Low: d(8840.00), High: d(8880.00), Close: d(8880.00), Open: d(8840.00), Volume: d(462403.00), Period: p(1663659000000)},
		{Low: d(8800.00), High: d(8900.00), Close: d(8820.00), Open: d(8810.00), Volume: d(787614.00), Period: p(1663734600000)},
		{Low: d(8800.00), High: d(8870.00), Close: d(8810.00), Open: d(8810.00), Volume: d(677890.00), Period: p(1663738200000)},
		{Low: d(8800.00), High: d(8850.00), Close: d(8810.00), Open: d(8810.00), Volume: d(1599221.00), Period: p(1663741800000)},
		{Low: d(8800.00), High: d(8950.00), Close: d(8930.00), Open: d(8810.00), Volume: d(1817537.00), Period: p(1663745400000)},
		{Low: d(8560.00), High: d(8920.00), Close: d(8590.00), Open: d(8850.00), Volume: d(1956165.00), Period: p(1663997400000)},
		{Low: d(8530.00), High: d(8630.00), Close: d(8570.00), Open: d(8580.00), Volume: d(963507.00), Period: p(1664001000000)},
		{Low: d(8560.00), High: d(8650.00), Close: d(8560.00), Open: d(8570.00), Volume: d(1202470.00), Period: p(1664004600000)},
		{Low: d(8470.00), High: d(8570.00), Close: d(8490.00), Open: d(8560.00), Volume: d(1795527.00), Period: p(1664008200000)},
		{Low: d(8360.00), High: d(8900.00), Close: d(8860.00), Open: d(8360.00), Volume: d(2232707.00), Period: p(1664170200000)},
		{Low: d(8700.00), High: d(8850.00), Close: d(8710.00), Open: d(8850.00), Volume: d(804025.00), Period: p(1664173800000)},
		{Low: d(8710.00), High: d(8750.00), Close: d(8740.00), Open: d(8720.00), Volume: d(490415.00), Period: p(1664177400000)},
		{Low: d(8700.00), High: d(8740.00), Close: d(8740.00), Open: d(8710.00), Volume: d(354268.00), Period: p(1664181000000)},
		{Low: d(8720.00), High: d(9000.00), Close: d(8960.00), Open: d(8720.00), Volume: d(1527278.00), Period: p(1664343000000)},
		{Low: d(8910.00), High: d(9060.00), Close: d(9000.00), Open: d(8960.00), Volume: d(1085106.00), Period: p(1664346600000)},
		{Low: d(8980.00), High: d(9030.00), Close: d(9020.00), Open: d(9000.00), Volume: d(473696.00), Period: p(1664350200000)},
		{Low: d(8980.00), High: d(9020.00), Close: d(8990.00), Open: d(9010.00), Volume: d(699219.00), Period: p(1664353800000)},
		{Low: d(8570.00), High: d(8990.00), Close: d(8650.00), Open: d(8950.00), Volume: d(2525245.00), Period: p(1664602200000)},
		{Low: d(8540.00), High: d(8750.00), Close: d(8580.00), Open: d(8670.00), Volume: d(862745.00), Period: p(1664605800000)},
		{Low: d(8400.00), High: d(8590.00), Close: d(8450.00), Open: d(8570.00), Volume: d(2648031.00), Period: p(1664609400000)},
		{Low: d(8400.00), High: d(8490.00), Close: d(8430.00), Open: d(8470.00), Volume: d(880884.00), Period: p(1664613000000)},
		{Low: d(8380.00), High: d(8610.00), Close: d(8450.00), Open: d(8460.00), Volume: d(2047200.00), Period: p(1664688600000)},
		{Low: d(8410.00), High: d(8490.00), Close: d(8460.00), Open: d(8470.00), Volume: d(508220.00), Period: p(1664692200000)},
		{Low: d(8430.00), High: d(8480.00), Close: d(8450.00), Open: d(8460.00), Volume: d(652416.00), Period: p(1664695800000)},
		{Low: d(8400.00), High: d(8460.00), Close: d(8440.00), Open: d(8440.00), Volume: d(906352.00), Period: p(1664699400000)},
		{Low: d(8420.00), High: d(8600.00), Close: d(8480.00), Open: d(8490.00), Volume: d(826543.00), Period: p(1664775000000)},
		{Low: d(8430.00), High: d(8570.00), Close: d(8450.00), Open: d(8480.00), Volume: d(867974.00), Period: p(1664778600000)},
		{Low: d(8430.00), High: d(8480.00), Close: d(8450.00), Open: d(8450.00), Volume: d(580218.00), Period: p(1664782200000)},
		{Low: d(8440.00), High: d(8540.00), Close: d(8500.00), Open: d(8450.00), Volume: d(957990.00), Period: p(1664785800000)},
		{Low: d(8400.00), High: d(8560.00), Close: d(8500.00), Open: d(8400.00), Volume: d(659743.00), Period: p(1664861400000)},
		{Low: d(8450.00), High: d(8500.00), Close: d(8480.00), Open: d(8500.00), Volume: d(1096047.00), Period: p(1664865000000)},
		{Low: d(8490.00), High: d(8520.00), Close: d(8500.00), Open: d(8490.00), Volume: d(390241.00), Period: p(1664868600000)},
		{Low: d(8480.00), High: d(8530.00), Close: d(8520.00), Open: d(8510.00), Volume: d(1052288.00), Period: p(1664872200000)},
		{Low: d(8400.00), High: d(8650.00), Close: d(8620.00), Open: d(8400.00), Volume: d(1420025.00), Period: p(1665207000000)},
		{Low: d(8580.00), High: d(8630.00), Close: d(8580.00), Open: d(8630.00), Volume: d(664976.00), Period: p(1665210600000)},
		{Low: d(8520.00), High: d(8580.00), Close: d(8520.00), Open: d(8570.00), Volume: d(811777.00), Period: p(1665214200000)},
		{Low: d(8500.00), High: d(8640.00), Close: d(8620.00), Open: d(8520.00), Volume: d(1135548.00), Period: p(1665217800000)},
		{Low: d(8360.00), High: d(8620.00), Close: d(8400.00), Open: d(8460.00), Volume: d(1757266.00), Period: p(1665293400000)},
		{Low: d(8390.00), High: d(8540.00), Close: d(8410.00), Open: d(8400.00), Volume: d(2069165.00), Period: p(1665297000000)},
		{Low: d(8370.00), High: d(8460.00), Close: d(8450.00), Open: d(8420.00), Volume: d(1283825.00), Period: p(1665300600000)},
		{Low: d(8450.00), High: d(8550.00), Close: d(8550.00), Open: d(8450.00), Volume: d(1443903.00), Period: p(1665304200000)},
		{Low: d(8460.00), High: d(8590.00), Close: d(8500.00), Open: d(8500.00), Volume: d(792803.00), Period: p(1665379800000)},
		{Low: d(8420.00), High: d(8520.00), Close: d(8470.00), Open: d(8510.00), Volume: d(1158445.00), Period: p(1665383400000)},
		{Low: d(8460.00), High: d(8500.00), Close: d(8490.00), Open: d(8480.00), Volume: d(806689.00), Period: p(1665387000000)},
		{Low: d(8460.00), High: d(8540.00), Close: d(8540.00), Open: d(8490.00), Volume: d(1334939.00), Period: p(1665390600000)},
		{Low: d(8490.00), High: d(8900.00), Close: d(8800.00), Open: d(8490.00), Volume: d(2464838.00), Period: p(1665466200000)},
		{Low: d(8700.00), High: d(8900.00), Close: d(8740.00), Open: d(8810.00), Volume: d(2086815.00), Period: p(1665469800000)},
		{Low: d(8730.00), High: d(8820.00), Close: d(8750.00), Open: d(8740.00), Volume: d(911276.00), Period: p(1665473400000)},
		{Low: d(8740.00), High: d(8890.00), Close: d(8890.00), Open: d(8760.00), Volume: d(1329047.00), Period: p(1665477000000)},
		{Low: d(8720.00), High: d(8960.00), Close: d(8800.00), Open: d(8900.00), Volume: d(1709787.00), Period: p(1665552600000)},
		{Low: d(8790.00), High: d(8850.00), Close: d(8840.00), Open: d(8800.00), Volume: d(1069264.00), Period: p(1665556200000)},
		{Low: d(8840.00), High: d(8930.00), Close: d(8900.00), Open: d(8840.00), Volume: d(1340324.00), Period: p(1665559800000)},
		{Low: d(8880.00), High: d(8950.00), Close: d(8950.00), Open: d(8900.00), Volume: d(1694492.00), Period: p(1665563400000)},
		{Low: d(8750.00), High: d(8960.00), Close: d(8900.00), Open: d(8750.00), Volume: d(1787429.00), Period: p(1665811800000)},
		{Low: d(8780.00), High: d(8910.00), Close: d(8780.00), Open: d(8910.00), Volume: d(1494451.00), Period: p(1665815400000)},
		{Low: d(8730.00), High: d(8880.00), Close: d(8820.00), Open: d(8780.00), Volume: d(924913.00), Period: p(1665819000000)},
		{Low: d(8790.00), High: d(8900.00), Close: d(8860.00), Open: d(8820.00), Volume: d(1028523.00), Period: p(1665822600000)},
		{Low: d(8640.00), High: d(8880.00), Close: d(8650.00), Open: d(8870.00), Volume: d(1524800.00), Period: p(1665898200000)},
		{Low: d(8620.00), High: d(8700.00), Close: d(8690.00), Open: d(8680.00), Volume: d(1513406.00), Period: p(1665901800000)},
		{Low: d(8570.00), High: d(8690.00), Close: d(8600.00), Open: d(8640.00), Volume: d(1445787.00), Period: p(1665905400000)},
		{Low: d(8490.00), High: d(8690.00), Close: d(8500.00), Open: d(8590.00), Volume: d(1722207.00), Period: p(1665909000000)},
		{Low: d(8500.00), High: d(8650.00), Close: d(8500.00), Open: d(8500.00), Volume: d(795140.00), Period: p(1665984600000)},
		{Low: d(8510.00), High: d(8580.00), Close: d(8550.00), Open: d(8510.00), Volume: d(488220.00), Period: p(1665988200000)},
		{Low: d(8530.00), High: d(8610.00), Close: d(8600.00), Open: d(8550.00), Volume: d(501777.00), Period: p(1665991800000)},
		{Low: d(8540.00), High: d(8700.00), Close: d(8690.00), Open: d(8590.00), Volume: d(1029969.00), Period: p(1665995400000)},
		{Low: d(8600.00), High: d(8750.00), Close: d(8680.00), Open: d(8600.00), Volume: d(911243.00), Period: p(1666071000000)},
		{Low: d(8630.00), High: d(8700.00), Close: d(8650.00), Open: d(8680.00), Volume: d(757086.00), Period: p(1666074600000)},
		{Low: d(8600.00), High: d(8700.00), Close: d(8700.00), Open: d(8640.00), Volume: d(425422.00), Period: p(1666078200000)},
		{Low: d(8630.00), High: d(8730.00), Close: d(8700.00), Open: d(8640.00), Volume: d(660782.00), Period: p(1666081800000)},
		{Low: d(8600.00), High: d(8720.00), Close: d(8670.00), Open: d(8620.00), Volume: d(469671.00), Period: p(1666157400000)},
		{Low: d(8650.00), High: d(8700.00), Close: d(8660.00), Open: d(8670.00), Volume: d(546293.00), Period: p(1666161000000)},
		{Low: d(8620.00), High: d(8670.00), Close: d(8670.00), Open: d(8650.00), Volume: d(804228.00), Period: p(1666164600000)},
		{Low: d(8660.00), High: d(8730.00), Close: d(8700.00), Open: d(8670.00), Volume: d(1012749.00), Period: p(1666168200000)},
		{Low: d(8580.00), High: d(8840.00), Close: d(8720.00), Open: d(8580.00), Volume: d(939987.00), Period: p(1666416600000)},
		{Low: d(8700.00), High: d(8830.00), Close: d(8720.00), Open: d(8720.00), Volume: d(1003220.00), Period: p(1666420200000)},
		{Low: d(8680.00), High: d(8780.00), Close: d(8680.00), Open: d(8730.00), Volume: d(767169.00), Period: p(1666423800000)},
		{Low: d(8610.00), High: d(8770.00), Close: d(8610.00), Open: d(8680.00), Volume: d(472382.00), Period: p(1666427400000)},
		{Low: d(8500.00), High: d(8610.00), Close: d(8500.00), Open: d(8610.00), Volume: d(785573.00), Period: p(1666503000000)},
		{Low: d(8490.00), High: d(8540.00), Close: d(8490.00), Open: d(8500.00), Volume: d(1158605.00), Period: p(1666506600000)},
		{Low: d(8450.00), High: d(8520.00), Close: d(8520.00), Open: d(8490.00), Volume: d(797369.00), Period: p(1666510200000)},
		{Low: d(8450.00), High: d(8530.00), Close: d(8450.00), Open: d(8510.00), Volume: d(816922.00), Period: p(1666513800000)},
		{Low: d(8490.00), High: d(8580.00), Close: d(8520.00), Open: d(8580.00), Volume: d(965156.00), Period: p(1666589400000)},
		{Low: d(8380.00), High: d(8520.00), Close: d(8380.00), Open: d(8520.00), Volume: d(1881966.00), Period: p(1666593000000)},
		{Low: d(8300.00), High: d(8430.00), Close: d(8310.00), Open: d(8380.00), Volume: d(1552299.00), Period: p(1666596600000)},
		{Low: d(8290.00), High: d(8360.00), Close: d(8340.00), Open: d(8310.00), Volume: d(1326790.00), Period: p(1666600200000)},
		{Low: d(8280.00), High: d(8450.00), Close: d(8320.00), Open: d(8450.00), Volume: d(1639128.00), Period: p(1666675800000)},
		{Low: d(8160.00), High: d(8320.00), Close: d(8180.00), Open: d(8310.00), Volume: d(2019753.00), Period: p(1666679400000)},
		{Low: d(8170.00), High: d(8320.00), Close: d(8200.00), Open: d(8180.00), Volume: d(1416805.00), Period: p(1666683000000)},
		{Low: d(8150.00), High: d(8230.00), Close: d(8190.00), Open: d(8200.00), Volume: d(1475601.00), Period: p(1666686600000)},
		{Low: d(8010.00), High: d(8260.00), Close: d(8010.00), Open: d(8150.00), Volume: d(1740023.00), Period: p(1666762200000)},
		{Low: d(7920.00), High: d(8040.00), Close: d(7940.00), Open: d(8010.00), Volume: d(4274066.00), Period: p(1666765800000)},
		{Low: d(7920.00), High: d(7950.00), Close: d(7920.00), Open: d(7940.00), Volume: d(3708253.00), Period: p(1666769400000)},
		{Low: d(7920.00), High: d(8040.00), Close: d(8020.00), Open: d(7920.00), Volume: d(2467113.00), Period: p(1666773000000)},
		{Low: d(7990.00), High: d(8230.00), Close: d(7990.00), Open: d(8230.00), Volume: d(3782570.00), Period: p(1667021400000)},
		{Low: d(7900.00), High: d(8010.00), Close: d(7900.00), Open: d(7990.00), Volume: d(4110182.00), Period: p(1667025000000)},
		{Low: d(7740.00), High: d(7970.00), Close: d(7750.00), Open: d(7910.00), Volume: d(1834825.00), Period: p(1667028600000)},
		{Low: d(7690.00), High: d(7800.00), Close: d(7760.00), Open: d(7760.00), Volume: d(2606139.00), Period: p(1667032200000)},
		{Low: d(7820.00), High: d(7990.00), Close: d(7950.00), Open: d(7950.00), Volume: d(4859163.00), Period: p(1667107800000)},
		{Low: d(7920.00), High: d(8030.00), Close: d(8010.00), Open: d(7960.00), Volume: d(4220234.00), Period: p(1667111400000)},
		{Low: d(8000.00), High: d(8080.00), Close: d(8050.00), Open: d(8010.00), Volume: d(1975520.00), Period: p(1667115000000)},
		{Low: d(8030.00), High: d(8210.00), Close: d(8200.00), Open: d(8040.00), Volume: d(2780194.00), Period: p(1667118600000)},
		{Low: d(8070.00), High: d(8250.00), Close: d(8210.00), Open: d(8070.00), Volume: d(1003287.00), Period: p(1667194200000)},
		{Low: d(8100.00), High: d(8210.00), Close: d(8110.00), Open: d(8210.00), Volume: d(642473.00), Period: p(1667197800000)},
		{Low: d(8110.00), High: d(8180.00), Close: d(8160.00), Open: d(8110.00), Volume: d(664372.00), Period: p(1667201400000)},
		{Low: d(8100.00), High: d(8260.00), Close: d(8200.00), Open: d(8150.00), Volume: d(1241301.00), Period: p(1667205000000)},
		{Low: d(8110.00), High: d(8450.00), Close: d(8440.00), Open: d(8170.00), Volume: d(2909458.00), Period: p(1667280600000)},
		{Low: d(8310.00), High: d(8450.00), Close: d(8360.00), Open: d(8440.00), Volume: d(778238.00), Period: p(1667284200000)},
		{Low: d(8240.00), High: d(8370.00), Close: d(8260.00), Open: d(8360.00), Volume: d(658420.00), Period: p(1667287800000)},
		{Low: d(8240.00), High: d(8450.00), Close: d(8440.00), Open: d(8260.00), Volume: d(1814124.00), Period: p(1667291400000)},
		{Low: d(8270.00), High: d(8440.00), Close: d(8300.00), Open: d(8440.00), Volume: d(1267103.00), Period: p(1667367000000)},
		{Low: d(8270.00), High: d(8510.00), Close: d(8510.00), Open: d(8300.00), Volume: d(1821017.00), Period: p(1667370600000)},
		{Low: d(8430.00), High: d(8540.00), Close: d(8440.00), Open: d(8510.00), Volume: d(559250.00), Period: p(1667374200000)},
		{Low: d(8420.00), High: d(8470.00), Close: d(8440.00), Open: d(8440.00), Volume: d(544851.00), Period: p(1667377800000)},
		{Low: d(8480.00), High: d(8730.00), Close: d(8730.00), Open: d(8550.00), Volume: d(4284720.00), Period: p(1667626200000)},
		{Low: d(8730.00), High: d(8730.00), Close: d(8730.00), Open: d(8730.00), Volume: d(1382828.00), Period: p(1667629800000)},
		{Low: d(8730.00), High: d(8730.00), Close: d(8730.00), Open: d(8730.00), Volume: d(1678201.00), Period: p(1667633400000)},
		{Low: d(8730.00), High: d(8730.00), Close: d(8730.00), Open: d(8730.00), Volume: d(549277.00), Period: p(1667637000000)},
		{Low: d(8800.00), High: d(9070.00), Close: d(9060.00), Open: d(8800.00), Volume: d(5342062.00), Period: p(1667712600000)},
		{Low: d(9040.00), High: d(9070.00), Close: d(9070.00), Open: d(9060.00), Volume: d(8126959.00), Period: p(1667716200000)},
		{Low: d(9070.00), High: d(9070.00), Close: d(9070.00), Open: d(9070.00), Volume: d(527101.00), Period: p(1667719800000)},
		{Low: d(9070.00), High: d(9070.00), Close: d(9070.00), Open: d(9070.00), Volume: d(702521.00), Period: p(1667723400000)},
		{Low: d(9160.00), High: d(9440.00), Close: d(9430.00), Open: d(9290.00), Volume: d(4409696.00), Period: p(1667799000000)},
		{Low: d(9410.00), High: d(9490.00), Close: d(9490.00), Open: d(9420.00), Volume: d(7522839.00), Period: p(1667802600000)},
		{Low: d(9490.00), High: d(9490.00), Close: d(9490.00), Open: d(9490.00), Volume: d(777299.00), Period: p(1667806200000)},
		{Low: d(9490.00), High: d(9490.00), Close: d(9490.00), Open: d(9490.00), Volume: d(405416.00), Period: p(1667809800000)},
		{Low: d(9300.00), High: d(9890.00), Close: d(9530.00), Open: d(9890.00), Volume: d(7097789.00), Period: p(1667885400000)},
		{Low: d(9460.00), High: d(9570.00), Close: d(9470.00), Open: d(9520.00), Volume: d(3033312.00), Period: p(1667889000000)},
		{Low: d(9380.00), High: d(9490.00), Close: d(9410.00), Open: d(9470.00), Volume: d(2714433.00), Period: p(1667892600000)},
		{Low: d(9390.00), High: d(9490.00), Close: d(9450.00), Open: d(9420.00), Volume: d(3876877.00), Period: p(1667896200000)},
		{Low: d(9250.00), High: d(9540.00), Close: d(9410.00), Open: d(9350.00), Volume: d(3448605.00), Period: p(1667971800000)},
		{Low: d(9400.00), High: d(9840.00), Close: d(9800.00), Open: d(9410.00), Volume: d(6547559.00), Period: p(1667975400000)},
		{Low: d(9640.00), High: d(9830.00), Close: d(9650.00), Open: d(9800.00), Volume: d(2416825.00), Period: p(1667979000000)},
		{Low: d(9650.00), High: d(9860.00), Close: d(9680.00), Open: d(9700.00), Volume: d(2463503.00), Period: p(1667982600000)},
		{Low: d(9640.00), High: d(9870.00), Close: d(9800.00), Open: d(9750.00), Volume: d(2000789.00), Period: p(1668231000000)},
		{Low: d(9520.00), High: d(9800.00), Close: d(9520.00), Open: d(9780.00), Volume: d(3214849.00), Period: p(1668234600000)},
		{Low: d(9520.00), High: d(9680.00), Close: d(9620.00), Open: d(9550.00), Volume: d(3019512.00), Period: p(1668238200000)},
		{Low: d(9610.00), High: d(9810.00), Close: d(9740.00), Open: d(9640.00), Volume: d(2473212.00), Period: p(1668241800000)},
		{Low: d(9450.00), High: d(9710.00), Close: d(9530.00), Open: d(9710.00), Volume: d(1455003.00), Period: p(1668317400000)},
		{Low: d(9510.00), High: d(9700.00), Close: d(9700.00), Open: d(9520.00), Volume: d(1341450.00), Period: p(1668321000000)},
		{Low: d(9520.00), High: d(9720.00), Close: d(9650.00), Open: d(9700.00), Volume: d(2922575.00), Period: p(1668324600000)},
		{Low: d(9470.00), High: d(9650.00), Close: d(9470.00), Open: d(9650.00), Volume: d(907574.00), Period: p(1668328200000)},
		{Low: d(9250.00), High: d(9620.00), Close: d(9250.00), Open: d(9510.00), Volume: d(1573592.00), Period: p(1668403800000)},
		{Low: d(9220.00), High: d(9420.00), Close: d(9380.00), Open: d(9270.00), Volume: d(1372258.00), Period: p(1668407400000)},
		{Low: d(9340.00), High: d(9530.00), Close: d(9490.00), Open: d(9380.00), Volume: d(3147032.00), Period: p(1668411000000)},
		{Low: d(9370.00), High: d(9550.00), Close: d(9370.00), Open: d(9490.00), Volume: d(2153637.00), Period: p(1668414600000)},
		{Low: d(9380.00), High: d(9750.00), Close: d(9670.00), Open: d(9450.00), Volume: d(1861478.00), Period: p(1668490200000)},
		{Low: d(9580.00), High: d(9700.00), Close: d(9650.00), Open: d(9670.00), Volume: d(2890813.00), Period: p(1668493800000)},
		{Low: d(9610.00), High: d(9700.00), Close: d(9670.00), Open: d(9610.00), Volume: d(1288957.00), Period: p(1668497400000)},
		{Low: d(9630.00), High: d(9800.00), Close: d(9730.00), Open: d(9650.00), Volume: d(2413843.00), Period: p(1668501000000)},
		{Low: d(9580.00), High: d(9780.00), Close: d(9630.00), Open: d(9750.00), Volume: d(803830.00), Period: p(1668576600000)},
		{Low: d(9630.00), High: d(9720.00), Close: d(9670.00), Open: d(9650.00), Volume: d(699785.00), Period: p(1668580200000)},
		{Low: d(9640.00), High: d(9700.00), Close: d(9640.00), Open: d(9700.00), Volume: d(393592.00), Period: p(1668583800000)},
		{Low: d(9580.00), High: d(9660.00), Close: d(9630.00), Open: d(9640.00), Volume: d(1443871.00), Period: p(1668587400000)},
		{Low: d(9300.00), High: d(9600.00), Close: d(9370.00), Open: d(9510.00), Volume: d(3845936.00), Period: p(1668835800000)},
		{Low: d(9310.00), High: d(9380.00), Close: d(9330.00), Open: d(9380.00), Volume: d(1380628.00), Period: p(1668839400000)},
	}

	series := techan.NewTimeSeries()
	series.Candles = bar_h1
	indicator := ichimoku.NewIndicator()

	err := indicator.MakeIchimokuInPast(series, 135)
	if err != nil {
		fmt.Println("error :", err)
	}

	lines_result := make([]*ichimoku.IchimokuStatus, 2)
	arr := indicator.GetList()
	for i := len(arr) - 2; i > 0; i-- {
		lines_result[0] = arr[i]   //current
		lines_result[1] = arr[i+1] // previous

		a, e := indicator.PreAnalyseIchimoku(lines_result)

		if e != nil {
			fmt.Println("err", e)
		}
		if a != nil {
			fmt.Print(a.Print())
		}
	}
}
