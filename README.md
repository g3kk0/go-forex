# Go Forex

Provides access to foreign exchange rates and currency conversions. [exchangeratesapi.io](https://exchangeratesapi.io/)

## Installation

```sh
go get github.com/g3kk0/go-forex
```

## Usage

```go
import forex "github.com/g3kk0/go-forex"

// create client
fc := forex.NewClient()

// request latest exchange rates
rates, err := fc.Latest()
if err != nil {
    log.Println(err)
}

// quote against a specific base (default EUR)
params := map[string]string{"base": "usd"}
rates, err := fc.Latest(params)
if err != nil {
    log.Println(err)
}

// quote against specific symbols
params := map[string]string{"base": "usd", "symbols": "gbp,eur"}
rate, err := fc.Latest(params)
if err != nil {
    log.Println(err)
}

// currency conversion
params := map[string]string{"from": "usd", "to": "gbp", "amount": "25.00"}
conversion, err := fc.Convert(params)
if err != nil {
    log.Println(err)
}

fmt.Println(conversion.Result)
```
