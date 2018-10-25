package goforex

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Convert struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"` // make this a string in json
	Result string `json:"result"`
}

func (c *Client) Convert(params ...map[string]string) (Convert, error) {
	var convert Convert

	from := params[0]["from"]
	to := params[0]["to"]
	amount := params[0]["amount"]

	if len(params) < 1 {
		err := errors.New("missing parameters")
		return convert, err
	}

	if from == "" {
		err := errors.New("missing 'from' parameter")
		return convert, err
	}

	if to == "" {
		err := errors.New("missing 'to' parameter")
		return convert, err
	}

	if amount == "" {
		err := errors.New("missing 'amount' parameter")
		return convert, err
	}

	p := map[string]string{"base": from, "symbols": to}
	rates, err := c.Latest(p)
	if err != nil {
		return convert, err
	}

	// how many decimals?
	rate := fmt.Sprint(rates.Rates[strings.ToUpper(to)])
	rateMultiplier, err := getMultiplier(rate)
	if err != nil {
		return convert, err
	}

	fmt.Printf("rate = %T %+v\n", rate, rate)
	fmt.Printf("rateMultiplier = %+v\n", rateMultiplier)

	amountMultiplier, err := getMultiplier(amount)
	if err != nil {
		return convert, err
	}

	fmt.Printf("amount = %T %+v\n", amount, amount)
	fmt.Printf("amountMultiplier = %+v\n", amountMultiplier)

	// convert to int64

	rateInt := int64(rates.Rates[strings.ToUpper(to)] * float64(rateMultiplier))

	fmt.Printf("rateInt = %+v\n", rateInt)

	amountFloat, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return convert, err
	}

	amountInt := int64(amountFloat * float64(amountMultiplier))

	fmt.Printf("amountInt = %+v\n", amountInt)

	// boo, err := strconv.ParseInt(amount, 10, 64)
	// if err != nil {
	// 	fmt.Printf("err = %+v\n", err)
	// 	return convert, err
	// }

	//	fmt.Printf("x = %T %+v\n", boo, boo)

	// perform calcs

	result := rateInt * amountInt

	fmt.Printf("result = %T %+v\n", result, result)

	result2 := float64(result) / float64(rateMultiplier)

	fmt.Printf("result2 = %T %+v\n", result2, result2)

	// back to floats

	//	toInt := rates.Rates[strings.ToUpper(to)] * float64(toMultiplier)
	//amountInt := strconv.Atoi(amount) * amountMultiplier

	//	fmt.Printf("toMultiplier = %+v\n", toMultiplier)
	//	fmt.Printf("amountMultiplier = %+v\n", amountMultiplier)
	//	fmt.Printf("toInt = %+v\n", toInt)
	//fmt.Printf("amountInt = %+v\n", amountInt)

	//toMultiplier := getMultiplier("10.00")
	//amountMultiplier := getMultiplier("10.00")

	//multiplier := getMultiplier()

	//	fmt.Printf("amount = %+v\n", amount)

	//	fmt.Printf("rates = %+v\n", rates)

	convert.From = from
	convert.To = to
	convert.Amount = amount

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
