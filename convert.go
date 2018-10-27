package goforex

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Convert struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount,string"`
	Result float64 `json:"result,string"`
}

func (c *Client) Convert(params ...map[string]string) (Convert, error) {
	var convert Convert

	if len(params) < 1 {
		err := errors.New("missing parameters")
		return convert, err
	}

	if params[0]["from"] == "" {
		err := errors.New("missing 'from' parameter")
		return convert, err
	}

	if params[0]["to"] == "" {
		err := errors.New("missing 'to' parameter")
		return convert, err
	}

	if params[0]["amount"] == "" {
		err := errors.New("missing 'amount' parameter")
		return convert, err
	}

	// missing zero!!!
	fmt.Printf("paramsAmount = %+v\n", params[0]["amount"])
	amount, err := strconv.ParseFloat(params[0]["amount"], 64)
	if err != nil {
		return convert, err
	}

	fmt.Printf("amount = %+v\n", amount)

	convert.From = params[0]["from"]
	convert.To = params[0]["to"]
	convert.Amount = amount

	p := map[string]string{"base": convert.From, "symbols": convert.To}
	rates, err := c.Latest(p)
	if err != nil {
		return convert, err
	}

	fmt.Printf("rates = %+v\n", rates)

	rateInt, rateDp := Ftoi(rates.Rates[strings.ToUpper(convert.To)])
	amountInt, amountDp := Ftoi(convert.Amount)

	fmt.Printf("rateInt = %+v\n", rateInt)
	fmt.Printf("rateDp = %+v\n", rateDp)
	fmt.Printf("amountInt = %+v\n", amountInt)
	fmt.Printf("amountDp = %+v\n", amountDp)
	//

	result := rateInt * amountInt

	fmt.Printf("result = %T %+v\n", result, result)

	result2 := Itof(result, rateDp)

	fmt.Printf("result2 = %T %+v\n", result2, result2)

	return convert, nil
}

func getMultiplier(in string) (int64, error) {
	s := strings.Split(in, ".")

	m := "1"
	for i := 0; i < len(s[1]); i++ {
		m = m + "0"
	}

	multiplier, err := strconv.Atoi(m)
	if err != nil {
		return 0, err
	}

	return int64(multiplier), nil
}
