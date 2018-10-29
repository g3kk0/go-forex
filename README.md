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

fmt.Printf("rates = %+v", rates)
// {Base:EUR Date:2018-10-29 Rates:map[ILS:4.217 RON:4.6647 NZD:1.7401 ...]}
```

### Get Latest Rates for a Specific Base

*Defaults to `EUR`*

```go
params := map[string]string{"base": "usd"}
rates, err := fc.Latest(params)
if err != nil {
    log.Println(err)
}

fmt.Printf("rates = %+v", rates)
// {Base:USD Date:2018-10-29 Rates:map[HUF:285.0364642826 ILS:3.7052983042 ...]}
```

### Get Rates for Specific Symbols

```go
params := map[string]string{"base": "usd", "symbols": "gbp,eur"}
rates, err := fc.Latest(params)
if err != nil {
    log.Println(err)
}

fmt.Printf("rates = %+v", rates)
// {Base:USD Date:2018-10-29 Rates:map[GBP:0.7801423425 EUR:0.8786574115]}
```

### Currency Conversion

```go
// currency conversion
params := map[string]string{"from": "usd", "to": "gbp", "amount": "25.62"}
conversion, err := fc.Convert(params)
if err != nil {
    log.Println(err)
}

fmt.Printf("conversion = %+v", conversion)
// {From:EUR To:GBP Amount:25.62 Result:22.7474856}
```
