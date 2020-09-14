package http

import (
	"github.com/jinzhu/copier"
	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"github.com/wdhafin/eFishery-test/entity"
	"github.com/wdhafin/eFishery-test/module/auth/usecase"
	"github.com/wdhafin/eFishery-test/pkg/utils"
	"github.com/wdhafin/eFishery-test/schema"
)

// AuthHandler  represent the httphandler for article
type AuthHandler struct {
	cUsecase usecase.AuthUsecase
}

// NewAuthHandler will initialize the auth / resources endpoint
func NewAuthHandler(e *echo.Echo, us usecase.AuthUsecase) {
	handler := &AuthHandler{
		cUsecase: us,
	}
	g := e.Group("/" + viper.GetString("route.public") + "/auth")

	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)
}

// Register will
func (cHandler *AuthHandler) Register(c echo.Context) error {
	request := new(schema.RegisterRequest)

	//parsing and validate
	err := utils.ParsingAndValidateParameter(c, request)
	if err != nil {
		return utils.ErrorParsingValidate(c, err)
	}

	userRequest := new(entity.User)
	copier.Copy(&userRequest, &request)

	user, err := cHandler.cUsecase.Register(*userRequest)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	userResponse := new(schema.UserResponse)
	copier.Copy(&userResponse, &user)

	return utils.SuccessResponse(c, userResponse)
}

// Login will
func (cHandler *AuthHandler) Login(c echo.Context) error {
	request := new(schema.LoginRequest)

	//parsing and validate
	err := utils.ParsingAndValidateParameter(c, request)
	if err != nil {
		return utils.ErrorParsingValidate(c, err)
	}

	authRequest := new(entity.User)
	copier.Copy(&authRequest, &request)

	token, err := cHandler.cUsecase.Login(*authRequest)
	if err != nil {
		return utils.ErrorResponse(c, err)
	}

	loginResponse := new(schema.LoginResponse)
	copier.Copy(&loginResponse, &token)

	return utils.SuccessResponse(c, loginResponse)
}
