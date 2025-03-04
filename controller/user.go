package controller

import (
	"go-crud-mongo/dto/request"
	"go-crud-mongo/dto/response"
	"go-crud-mongo/helper"
	"go-crud-mongo/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddUser(c echo.Context) error {
	in, err := helper.ReadBody[request.AddUserIn](c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.ResponseCommon{
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		})
	}

	usr := model.NewUser()
	usr.Name = in.Name
	usr.Age = in.Age
	usr.Email = in.Email
	err = usr.Insert()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseCommon{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		})
	}

	out := response.AddUserOut{}
	out.Data = usr
	out.SetStatus(http.StatusCreated)
	out.SetMessage(`User created !!`)

	return c.JSON(http.StatusCreated, out)
}

func GetAllUsers(c echo.Context) error {
	usr := model.NewUser()
	users, err := usr.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.ResponseCommon{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		})
	}

	out := response.GetUsersOut{}
	out.Data = &users
	out.SetStatus(http.StatusOK)
	out.SetMessage(`Users obtained !!`)

	return c.JSON(http.StatusOK, out)
}
