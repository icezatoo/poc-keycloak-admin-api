package handlers

import (
	"fmt"
	"github.com/Nerzal/gocloak/v8"
	"github.com/labstack/echo/v4"
	"net/http"
	"poc-keycloak-admin-api/requests"
	"poc-keycloak-admin-api/responses"
	s "poc-keycloak-admin-api/servers"
	"poc-keycloak-admin-api/utils"
)

type UserHandler struct {
	server *s.Server
}

func NewUserHandler(server *s.Server) *UserHandler {
	return &UserHandler{server: server}
}



// GetUsers godoc
// @Summary Get users list
// @Description Get the a user list
// @ID user-get
// @Tags User Actions
// @Produce json
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users [get]
func (u *UserHandler) GetUsers(c echo.Context) error {
	accessToken := utils.GetAccessTokenFromHeader(c)

	// get the userId  from access token
	//_ , jwtClaims , _ := u.server.GoCloakClient.DecodeAccessToken(u.server.Ctx,accessToken, u.server.Config.KeycloakConfig.Realm,"")
	//userId := (*jwtClaims)["sub"]
	//fmt.Println(userId)

	users ,_ :=	u.server.GoCloakClient.GetUsers(
		u.server.Ctx,
		accessToken,
		u.server.Config.KeycloakConfig.Realm ,
		gocloak.GetUsersParams{},
		)

	return responses.Response(c, http.StatusOK, users)
}


// GetUserDetail godoc
// @Summary Get user info
// @Description Get the a user detail
// @ID user-detail
// @Tags User Actions
// @Produce json
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Security ApiKeyAuth
// @Router /me [get]
func (u *UserHandler) GetUserDetail(c echo.Context) error  {
	accessToken := utils.GetAccessTokenFromHeader(c)

    userInfo , err :=u.server.GoCloakClient.GetUserInfo(u.server.Ctx, accessToken , u.server.Config.KeycloakConfig.Realm)

    if err !=nil {
		return responses.ErrorResponse(c, http.StatusNotFound, "Cannot get the userInfo")
	}

	return responses.Response(c, http.StatusOK, userInfo)
}


// CreateUser godoc
// @Summary Create user
// @Description  Create user
// @ID user-create
// @Tags User Actions
// @Accept json
// @Produce json
// @Param params body requests.CreateUserRequest true "UserInfo"
// @Success 200 {object} responses.Data
// @Failure 401 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users [post]
func (u *UserHandler) CreateUser(c echo.Context) error  {
	accessToken := utils.GetAccessTokenFromHeader(c)
	createUserRequest := new(requests.CreateUserRequest)

	fmt.Println(createUserRequest)

	if err := c.Bind(createUserRequest); err != nil {
		return err
	}

	if err := createUserRequest.Validate(); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "Required fields are empty")
	}

	user := gocloak.User{
		FirstName: gocloak.StringP(createUserRequest.FirstName),
		LastName:  gocloak.StringP(createUserRequest.LastName),
		Email:     gocloak.StringP(createUserRequest.Email),
		Enabled:   gocloak.BoolP(createUserRequest.Enabled),
		Username:  gocloak.StringP(createUserRequest.Username),
	}

	 _, err :=	u.server.GoCloakClient.CreateUser(u.server.Ctx,accessToken,u.server.Config.KeycloakConfig.Realm,user)

 	if err != nil {
 		return responses.ErrorResponse(c,http.StatusInternalServerError , "Cannot create user in keycloak")
	 }

    return responses.MessageResponse(c, http.StatusCreated, "User successfully created")
}

// DeleteUser godoc
// @Summary Delete user
// @Description  Delete user by id
// @ID user-delete
// @Tags User Actions
// @Param id path int true "User ID"
// @Success 204 {object} responses.Data
// @Failure 404 {object} responses.Error
// @Security ApiKeyAuth
// @Router /users/{id} [delete]
func (u *UserHandler) DeleteUser(c  echo.Context) error  {

	accessToken := utils.GetAccessTokenFromHeader(c)
	id := c.Param("id")

	if id == "" {
		return responses.ErrorResponse(c,http.StatusBadRequest , "The id is null or missing")
	}

   err := u.server.GoCloakClient.DeleteUser(u.server.Ctx,accessToken,u.server.Config.KeycloakConfig.Realm,id)

   if err != nil {
	 return responses.ErrorResponse(c,http.StatusInternalServerError , "Cannot delete user in keycloak")
   }

   return responses.MessageResponse(c, http.StatusNoContent, "User deleted successfully")

}

