# TrueFX

The TrueFX API client for Go. Read [TrueFX](https://www.truefx.com/) for more information.

## Installation

```go
go get github.com/tonkla/gotruefx
```

## Usage

TrueFX provides a price feed of these ten paris by default,  
EUR/USD, USD/JPY, GBP/USD, EUR/GBP, USD/CHF, EUR/JPY, EUR/CHF, USD/CAD, AUD/USD, GBP/JPY

* get all: ```feed.Get()```
* get by the specific symbol: ```feed.GetBySymbol("EUR/USD")```
* get more, separated by comma: ```feed.GetBySymbol("EUR/USD,USD/JPY")```

```go
package main

import (
  "fmt"
  "encoding/json"

  "github.com/tonkla/gotruefx"
)

func main() {
  var ticks []truefx.Tick
  feed := truefx.NewFeed()
  result := feed.GetBySymbol("EUR/USD")
  err := json.Unmarshal(result, &ticks)
  if err != nil {
      panic(err)
  }
  tick := ticks[0]
  fmt.Printf("%v\n", tick)
  fmt.Printf("Symbol: %s\n", tick.Symbol)
  fmt.Printf("Timestamp: %d\n", tick.Timestamp)
  fmt.Printf("Bid: %.5f\n", tick.Bid)
  fmt.Printf("Offer: %.5f\n", tick.Offer)
  fmt.Printf("High: %.5f\n", tick.High)
  fmt.Printf("Low: %.5f\n", tick.Low)
  fmt.Printf("Open: %.5f\n", tick.Open)
  fmt.Printf("Spread: %.1f", tick.Spread)
}
```

Results

```
{EUR/USD 1485553500705 1.0697 1.06978 1.0658 1.07253 1.06816 0.8}
Symbol: EUR/USD
Timestamp: 1485553500705
Bid: 1.06970
Offer: 1.06978
High: 1.06580
Low: 1.07253
Open: 1.06816
Spread: 0.8
```

Authorized session can access to more minor pairs. [Register](https://www.truefx.com)

AUD/CAD, AUD/CHF, AUD/JPY, AUD/NZD, CAD/CHF, CAD/JPY, CHF/JPY, EUR/AUD, EUR/CAD,  
EUR/NOK, EUR/NZD, GBP/CAD, GBP/CHF, NZD/JPY, NZD/USD, USD/NOK, USD/SEK

```go
feed := truefx.NewFeedAuthorize("USERNAME", "PASSWORD")
feed.GetBySymbol("AUD/JPY")
```

**Issue:** Getting a tick data by authorized session got an empty result `[]`; authorization was succeeded but something went wrong with the session. Even unauthorized session (request by fake id like ```&id=user:passwd:session:1```) can get a tick data of the above minor pairs, if you change the session_id everytime you request. Weird!!

## Contributing

1. Fork it ( https://github.com/tonkla/gotruefx/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [tonkla](https://github.com/tonkla) Surakarn Samkaew - creator, maintainer
