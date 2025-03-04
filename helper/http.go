package helper

import (
	"errors"

	"github.com/labstack/echo/v4"
)

func ReadBody[T any | struct{}](c echo.Context) (out T, err error) {
	errBind := c.Bind(&out)
	if errBind != nil {
		Log.Error(errBind, `failed to bind request to valid payload`)
		err = errors.New(`invalid payload, please check your request and try again`)
		return
	}

	err = ValidateStruct(out)
	if err != nil {
		return
	}

	return
}
