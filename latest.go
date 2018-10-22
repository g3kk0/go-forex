package goforex

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strings"
)

type Latest struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func (c *Client) Latest(params ...map[string]string) (Latest, error) {
	var latest Latest

	u, err := url.Parse(BaseUrl)
	if err != nil {
		return latest, err
	}

	u.Path = "latest"
	if len(params) > 0 {
		q := u.Query()
		if params[0]["base"] != "" {
			q.Set("base", strings.ToUpper(params[0]["base"]))
		}
		if params[0]["symbols"] != "" {
			q.Set("symbols", strings.ToUpper(params[0]["symbols"]))
		}
		u.RawQuery = q.Encode()
	}

	resp, err := c.Conn.Get(u.String())
	if err != nil {
		return latest, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return latest, err
	}

	err = json.Unmarshal(body, &latest)
	if err != nil {
		return latest, err
	}

	return latest, nil
}
