package responses

import "github.com/Nerzal/gocloak/v8"

type UserListResponse struct {
	Items    []*gocloak.User `json:"items"`
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
}

func NewUserResponse(users []*gocloak.User) *UserListResponse {
	userResponse := UserListResponse{
		Items:    users,
		Page:     1,
		PageSize: 10,
	}
	return &userResponse
}
