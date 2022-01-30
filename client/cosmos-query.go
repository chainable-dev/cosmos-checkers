package client

import (
	"github.com/go-resty/resty/v2"
)

func Auth(address string) *resty.Response {
	client := resty.New()
	resp, _ := client.R().
		EnableTrace().
		Get("https://api.cosmos.network/auth/accounts/" + address)
	return resp
}
