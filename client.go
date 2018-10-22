package goforex

import (
	"net/http"
)

const BaseUrl = "https://api.exchangeratesapi.io/"

type Client struct {
	Conn *http.Client
}

func NewClient() *Client {
	var c Client
	c.Conn = &http.Client{}
	return &c
}
