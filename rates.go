package goforex

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const BaseUrl = "https://api.exchangeratesapi.io/"

type Client struct {
	Conn *http.Client
	//	Key  string
}

type Latest struct {
	Base string `json:"base"`
	Date string `json:"date"`
	//	Rates       string `json:"base"`
}

func NewClient() *Client {
	var c Client
	c.Conn = &http.Client{}
	//	c.Key = key
	return &c
}

func (c *Client) Latest() Latest {
	u, err := url.Parse(BaseUrl)
	if err != nil {
		log.Println(err)
	}
	u.Path = "latest"
	//	q := u.Query()
	//	q.Set("access_key", c.Key)
	//	u.RawQuery = q.Encode()

	resp, err := c.Conn.Get(u.String())
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var latest Latest
	err = json.Unmarshal(body, &latest)
	if err != nil {
		log.Println(err)
	}

	return latest
}
