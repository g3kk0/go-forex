package goforex

import (
	"errors"
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

	switch {
	case params[0]["from"] == "":
		err := errors.New("missing 'from' parameter")
		return convert, err
	case params[0]["to"] == "":
		err := errors.New("missing 'to' parameter")
		return convert, err
	case params[0]["amount"] == "":
		err := errors.New("missing 'amount' parameter")
		return convert, err
	}

	amount, err := strconv.ParseFloat(params[0]["amount"], 64)
	if err != nil {
		return convert, err
	}

	convert.From = params[0]["from"]
	convert.To = params[0]["to"]
	convert.Amount = amount

	p := map[string]string{"base": convert.From, "symbols": convert.To}
	rates, err := c.Latest(p)
	if err != nil {
		return convert, err
	}

	rateInt, rateDp := Ftoi(rates.Rates[strings.ToUpper(convert.To)])
	amountInt, amountDp := Ftoi(convert.Amount)

	resultInt := rateInt * amountInt

	dp := rateDp + amountDp
	result := Itof(resultInt, dp)

	convert.Result = result

	return convert, nil
}
