package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)



type UserKeycloak struct {
	FirstName    string `json:"firstName" validate:"required" example:"Test"`
	LastName string `json:"lastName" validate:"required" example:"lastName"`
	Email string `json:"email" validate:"required" example:"john.doe@example.com"`
	Enabled bool `json:"enabled" validate:"required" example:"true"`
	Username string `json:"username" validate:"required" example:"Hello"`
}

func (ba UserKeycloak) Validate() error {
	return validation.ValidateStruct(&ba,
		validation.Field(&ba.Email, is.Email),
		validation.Field(&ba.FirstName),
		validation.Field(&ba.Enabled),
		validation.Field(&ba.Username),
	)
}

type CreateUserRequest struct {
	UserKeycloak
}