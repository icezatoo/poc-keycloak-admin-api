package user

import (
	"context"
)

func (userService *Service) DeleteUser(userID string, accessToken string, realm string) error {
	return userService.Client.DeleteUser(context.Background(), accessToken, realm, userID)
}
