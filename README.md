## Techan
![](https://travis-ci.org/algo-boyz/techan.svg?branch=master)

[![codecov](https://codecov.io/gh/algo-boyz/techan/branch/master/graph/badge.svg)](https://codecov.io/gh/algo-boyz/techan)

TechAn is a **tech**nical **an**alysis library for Go!

## Features 
* Basic and advanced technical analysis indicators
* Profit and trade analysis
* Strategy building

### Installation
```sh
$ go get github.com/algo-boyz/techan
```

### Quickstart
```go
series := techan.NewTimeSeries()

// fetch this from your preferred exchange
dataset := [][]string{
	// Timestamp, Open, Close, High, Low, volume
	{"1234567", "1", "2", "3", "5", "6"},
}

for _, datum := range dataset {
	start, _ := strconv.ParseInt(datum[0], 10, 64)
	period := techan.NewTimePeriod(time.Unix(start, 0), time.Hour*24)

	candle := techan.NewCandle(period)
	candle.Open = decimal.NewFromString(datum[1])
	candle.Close = decimal.NewFromString(datum[2])
	candle.High = decimal.NewFromString(datum[3])
	candle.Low = decimal.NewFromString(datum[4])

	series.AddCandle(candle)
}

Closes := techan.NewCloseIndicator(series)
movingAverage := techan.NewEMAIndicator(Closes, 10) // Create an exponential moving average with a window of 10

fmt.Println(movingAverage.Calculate(0).StringFixed(2))
```

### Create trading strategies
```go
indicator := techan.NewCloseIndicator(series)

// record trades on this object
record := techan.NewTradingRecord()

entryConstant := techan.NewConstantIndicator(30)
exitConstant := techan.NewConstantIndicator(10)

// Is satisfied when the price ema moves above 30 and the current position is new
entryRule := techan.And(
	techan.NewCrossUpIndicatorRule(entryConstant, indicator),
	techan.PositionNewRule{})
	
// Is satisfied when the price ema moves below 10 and the current position is open
exitRule := techan.And(
	techan.NewCrossDownIndicatorRule(indicator, exitConstant),
	techan.PositionOpenRule{})

strategy := techan.RuleStrategy{
	UnstablePeriod: 10, // Period before which ShouldEnter and ShouldExit will always return false
	EntryRule:      entryRule,
	ExitRule:       exitRule,
}

strategy.ShouldEnter(0, record) // returns false
```

### Credits
Techan is heavily influenced by [ta4j](https://github.com/ta4j/ta4j)

### License
Techan is released under the MIT license. See [LICENSE](./LICENSE) for details.
