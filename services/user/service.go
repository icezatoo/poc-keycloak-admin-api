package user

import (
	"github.com/Nerzal/gocloak/v8"
	"poc-keycloak-admin-api/requests"
)

type ServiceWrapper interface {
	GetUserList(request *requests.GetKeycloakUserParamsRequest, accessToken string, realm string) ([]*gocloak.User, error)
	SaveUser(user gocloak.User, accessToken string, realm string) (string, error)
	GetUserInfo(accessToken string, realm string)
	DeleteUser(accessToken string, realm string, userId string) error
}

type Service struct {
	Client gocloak.GoCloak
}

func NewUserService(client gocloak.GoCloak) *Service {
	return &Service{Client: client}
}
