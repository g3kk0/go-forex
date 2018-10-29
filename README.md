# Go Forex

Provides access to foreign exchange rates and currency conversions. [exchangeratesapi.io](https://exchangeratesapi.io/)

## Installation

```sh
go get github.com/g3kk0/go-forex
```

## Usage

Import package and create client.

```go
import forex "github.com/g3kk0/go-forex"

fc := forex.NewClient()
```

### Get Latest Rates

```go
rates, err := fc.Latest()
if err != nil {
    log.Println(err)
}

fmt.Println(rates)

{Base:EUR Date:2018-10-29 Rates:map[ILS:4.217 RON:4.6647 NZD:1.7401 CZK:25.849 ...]}

fmt.Println(rates)
// {Base:EUR Date:2018-10-29 Rates:map[ILS:4.217 RON:4.6647 NZD:1.7401 CZK:25.849 ...]}
```

```
{Base:EUR Date:2018-10-29 Rates:map[ILS:4.217 RON:4.6647 NZD:1.7401 CZK:25.849 ...]}
```

### Get Latest Rates for a Specific Base



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

{From:eur To:gbp Amount:46.44 Result:41.2331472}

```
