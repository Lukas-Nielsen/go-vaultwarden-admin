package vaultwarden

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	base   string
	token  string
	debug  bool
	client *resty.Client
}

func NewClient(base string, token string, debug bool) (*Client, error) {
	c := Client{
		base:  strings.TrimSuffix(base, "/"),
		token: token,
		debug: debug,
	}

	c, err := c.setup()

	return &c, err
}

func (c *Client) setup() (Client, error) {
	c.client = resty.
		New().
		SetDebug(c.debug).
		SetBaseURL(fmt.Sprintf("%s/admin", c.base)).
		SetHeader("User-Agent", "github.com/lukas-nielsen/go-vaultwarden-admin")

	resp, err := c.client.R().
		SetFormData(map[string]string{"token": c.token}).
		Post("")

	if err != nil {
		return *c, err
	}

	if resp.IsError() {
		return *c, fmt.Errorf("%s", resp.String())
	}

	return *c, nil
}
