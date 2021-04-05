package utils

import (
	"github.com/labstack/echo/v4"
	"strings"
)

func extractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func GetAccessTokenFromHeader(c echo.Context) string {
	token := c.Request().Header.Get("Authorization")
	return extractBearerToken(token)
}
