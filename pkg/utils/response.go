package utils

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/wdhafin/eFishery-test/schema"
)

//ParsingError is
type ParsingError struct {
	msg string
}

func (re *ParsingError) Error() string { return re.msg }

// SuccessResponse returns
func SuccessResponse(ctx echo.Context, data interface{}) error {

	responseData := schema.Response{
		Success: true,
		Data:    data,
	}

	//log

	return ctx.JSON(http.StatusOK, responseData)
}

// ErrorResponse returns
func ErrorResponse(ctx echo.Context, err error) error {

	log.Println(err)
	responseData := schema.ResponseError{
		Success: false,
		Error:   err.Error(),
	}

	//log

	return ctx.JSON(http.StatusBadRequest, responseData)
}

// UnauthorizedResponse returns
func UnauthorizedResponse(ctx echo.Context) error {

	responseData := schema.ResponseError{
		Success: false,
		Error:   "unauthorized",
	}

	//log

	return ctx.JSON(http.StatusUnauthorized, responseData)

}

// ErrorValidate returns
func ErrorValidate(ctx echo.Context, err error) error {

	errorMap := []map[string]interface{}{}

	for _, err := range err.(validator.ValidationErrors) {
		errorMap = append(errorMap, map[string]interface{}{(ToSnakeCase(err.StructField())): err.Tag()})
	}

	responseData := schema.ResponseError{
		Success: false,
		Error:   errorMap,
	}

	//log

	return ctx.JSON(http.StatusUnprocessableEntity, responseData)
}

// ErrorParsing returns
func ErrorParsing(ctx echo.Context, err error) error {

	responseData := schema.ResponseError{
		Success: false,
		Error:   err.Error(),
	}

	//log

	return ctx.JSON(http.StatusUnprocessableEntity, responseData)
}

// ErrorParsingValidate returns
func ErrorParsingValidate(ctx echo.Context, err error) (errs error) {
	switch err.(type) {
	default:
		errs = ErrorValidate(ctx, err)
	case *ParsingError:
		errs = ErrorParsing(ctx, err)
	}

	return errs
}
