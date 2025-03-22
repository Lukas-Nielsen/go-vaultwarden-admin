package vaultwarden

import (
	"fmt"
)

func (c *Client) Invite(email string) (User, error) {
	var res User
	resp, err := c.client.
		R().
		SetResult(&res).
		SetBody(map[string]string{"email": email}).
		Post("/invite")

	if err != nil {
		return User{}, err
	}

	if resp.IsError() {
		return User{}, fmt.Errorf("%s", resp.String())
	}

	return res, nil
}
