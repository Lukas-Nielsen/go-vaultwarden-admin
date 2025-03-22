package vaultwarden

import (
	"fmt"
)

func (c *Client) Logout() error {
	resp, err := c.client.
		R().
		Get("/logout")

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("%s", resp.String())
	}

	return nil
}
