package common

import (
	"context"
	"fmt"

	"qShell/services"
)

type Token struct {
	Username string
	Password string
}

func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"username": t.Username, "password": t.Password}, nil
}

func (t *Token) RequireTransportSecurity() bool {
	return false
}

func CheckUser(username, password string) error {
	userData, err := services.LoadUserConfig()
	if err != nil {
		return err
	}

	if pass, ok := userData[username]; ok {
		if pass == password {
			return nil
		}
	}
	return fmt.Errorf("auth error")
}
