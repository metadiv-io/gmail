# gmail

## Installation

```bash
go get -u github.com/metadiv-io/gmail
```

## Highlights

* gmail.NewAuth(host string, port int, user string, password string) *Auth

* gmail.SendEmail(auth *Auth, mail *Message) error
