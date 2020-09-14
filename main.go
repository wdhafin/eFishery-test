package main

import (
	"os"

	appMiddleware "github.com/wdhafin/eFishery-test/middleware"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	appInit "github.com/wdhafin/eFishery-test/init"

	_auth "github.com/wdhafin/eFishery-test/module/auth/usecase"

	_authHttpHandler "github.com/wdhafin/eFishery-test/module/auth/handler/http"

	_authRepo "github.com/wdhafin/eFishery-test/module/auth/store"
)

func init() {
	// Start pre-requisite app dependencies
	appInit.StartAppInit()
}

func main() {
	// Get PG Conn Instance
	pgDb, err := appInit.ConnectToPGServer()
	if err != nil {
		os.Exit(1)
	}

	// DI: Repository & Usecase
	authRepo := _authRepo.NewPgRepository(*pgDb)

	authUc := _auth.NewAuthUsecase(authRepo)
	// End of DI Steps

	// http handler
	runHTTPHandler(authUc)
	// end of http handler
}

func runHTTPHandler(authUc _auth.AuthUsecase) {
	// initiating echo web handler instance
	e := echo.New()

	//// Defining echo middlewares, e.g. JWT, CORS, http logger, etc

	// JWT middleware, if the path contains restricted then requires valid jwt
	e.Use(appMiddleware.EchoJWTAccessAuth())

	// Http logger
	e.Use(appMiddleware.EchoHTTPLogger())

	// Custom Middleware
	e.Use(appMiddleware.EchoCustomMiddleware)

	// usecase injection
	_authHttpHandler.NewAuthHandler(e, authUc)
	// end of usecase injection

	// Start the echo web handler framework
	e.Logger.Fatal(e.Start(viper.GetString("server.port")))
}
