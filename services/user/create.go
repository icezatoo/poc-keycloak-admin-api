package user

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
)

func (userService *Service) SaveUser(user gocloak.User, accessToken string, realm string) (string, error) {
	return userService.Client.CreateUser(context.Background(), accessToken, realm, user)
}
