package vaultwarden

import (
	"fmt"
	"time"
)

type User struct {
	Status                  int       `json:"_status"`
	AvatarColor             any       `json:"avatarColor"`
	CreatedAt               string    `json:"createdAt"`
	CreationDate            time.Time `json:"creationDate"`
	Culture                 string    `json:"culture"`
	Email                   string    `json:"email"`
	EmailVerified           bool      `json:"emailVerified"`
	ForcePasswordReset      bool      `json:"forcePasswordReset"`
	ID                      string    `json:"id"`
	Key                     string    `json:"key"`
	LastActive              string    `json:"lastActive"`
	MasterPasswordHint      string    `json:"masterPasswordHint"`
	Name                    string    `json:"name"`
	Object                  string    `json:"object"`
	Organizations           []any     `json:"organizations"`
	Premium                 bool      `json:"premium"`
	PremiumFromOrganization bool      `json:"premiumFromOrganization"`
	PrivateKey              string    `json:"privateKey"`
	ProviderOrganizations   []any     `json:"providerOrganizations"`
	Providers               []any     `json:"providers"`
	SecurityStamp           string    `json:"securityStamp"`
	TwoFactorEnabled        bool      `json:"twoFactorEnabled"`
	UserEnabled             bool      `json:"userEnabled"`
	UsesKeyConnector        bool      `json:"usesKeyConnector"`
}

func (c *Client) UsersGetAll() ([]User, error) {
	var res []User
	resp, err := c.client.
		R().
		SetResult(&res).
		Get("/users")

	if err != nil {
		return []User{}, err
	}

	if resp.IsError() {
		return []User{}, fmt.Errorf("%s", resp.String())
	}

	return res, nil
}

func (c *Client) UsersDisable(user string) error {
	var res []User
	resp, err := c.client.
		R().
		SetResult(&res).
		Post(fmt.Sprintf("/users/%s/disable", user))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("%s", resp.String())
	}

	return nil
}

func (c *Client) UsersEnable(user string) error {
	var res []User
	resp, err := c.client.
		R().
		SetResult(&res).
		Post(fmt.Sprintf("/users/%s/enable", user))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("%s", resp.String())
	}

	return nil
}

func (c *Client) UsersInviteResend(user string) error {
	var res []User
	resp, err := c.client.
		R().
		SetResult(&res).
		Post(fmt.Sprintf("/users/%s/invite/resend", user))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("%s", resp.String())
	}

	return nil
}

func (c *Client) UsersDeauth(user string) error {
	var res []User
	resp, err := c.client.
		R().
		SetResult(&res).
		Post(fmt.Sprintf("/users/%s/deauth", user))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("%s", resp.String())
	}

	return nil
}

func (c *Client) UsersDelete(user string) error {
	var res []User
	resp, err := c.client.
		R().
		SetResult(&res).
		Post(fmt.Sprintf("/users/%s/delete", user))

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("%s", resp.String())
	}

	return nil
}
