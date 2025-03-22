# go-vaultwarden-admin

package to wrap the api of [vaultwarden server admin](https://github.com/dani-garcia/vaultwarden)

## used packages

- [Resty](https://github.com/go-resty/resty)

## getting started

```sh
go get github.com/Lukas-Nielsen/go-vaultwarden-admin
```

```go
import "github.com/Lukas-Nielsen/go-vaultwarden-admin"
```

## usage

### conf

you need the base url and the admin token

```go
c, err := vaultwarden.NewClient(baseUrl string, adminToken string, debug bool);

// example
c, err := vaultwarden.NewClient("https://bw.example.com", "top-secret", false)
```

### functions

#### invite user

```go
user, err := c.Invite(email string)

// example
user, err := c.Invite("user@example.com")
```

#### list all users

```go
users, err := c.UsersGetAll()
```

#### user disable, enable, delete, resend invite and deauth

```go
err := c.Users<Disable|Enable|Delete|InviteResend|Deauth>(userid string)

// example
err := c.UsersDelete("00000000-0000-0000-0000-000000000000")
```

#### logout

you should always perform a logout

```go
err := c.Logout()
```
