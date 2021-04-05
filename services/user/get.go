package user

import (
	"context"
	"github.com/Nerzal/gocloak/v8"
	"poc-keycloak-admin-api/requests"
)

func (userService *Service) GetUserList(request *requests.GetKeycloakUserParamsRequest, accessToken string, realm string) ([]*gocloak.User, error) {
	return userService.Client.GetUsers(context.Background(), accessToken, realm, request.GetUsersParams)
}

func (userService *Service) GetUserInfo(accessToken string, realm string) (*gocloak.UserInfo, error) {
	return userService.Client.GetUserInfo(context.Background(), accessToken, realm)
}
