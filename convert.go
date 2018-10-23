package goforex

import (
	"errors"
	"fmt"
)

type Convert struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Amount string `json:"amount"`
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

	convert.From = from
	convert.To = to
	convert.Amount = amount

	// how many decimals?
	multiplier :=

		fmt.Printf("amount = %+v\n", amount)

	fmt.Printf("rates = %+v\n", rates)

	// u, err := url.Parse(BaseUrl)
	// if err != nil {
	// 	return latest, err
	// }

	// u.Path = "latest"
	// if len(params) > 0 {
	// 	q := u.Query()
	// 	if params[0]["base"] != "" {
	// 		q.Set("base", strings.ToUpper(params[0]["base"]))
	// 	}
	// 	if params[0]["symbols"] != "" {
	// 		q.Set("symbols", strings.ToUpper(params[0]["symbols"]))
	// 	}
	// 	u.RawQuery = q.Encode()
	// }

	// resp, err := c.Conn.Get(u.String())
	// if err != nil {
	// 	return latest, err
	// }
	// defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return latest, err
	// }

	// err = json.Unmarshal(body, &latest)
	// if err != nil {
	// 	return latest, err
	// }

	return convert, nil
}
