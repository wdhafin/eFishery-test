package middleware

import (
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"github.com/wdhafin/eFishery-test/pkg/helper"
)

// EchoCustomMiddleware is
func EchoCustomMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}

// EchoJWTAccessAuth is
func EchoJWTAccessAuth() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper: func(c echo.Context) bool {
			return !strings.Contains(c.Path(), viper.GetString("route.restricted"))
		},
		SigningKey: []byte(viper.GetString("auth.accessSecret")),
		SuccessHandler: func(c echo.Context) {
			userJWT := helper.GetAuthenticatedUser(c.Get("user").(*jwt.Token))
			c.Set("Name", userJWT.Name)
		},
	})
}

// EchoHTTPLogger is
// todo: if the env is production then log them using file instead of stdout
func EchoHTTPLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n",
	})
}
