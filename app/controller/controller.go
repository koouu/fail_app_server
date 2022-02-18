package controller

import (
	"net/http"
	"strconv"
	"fail_app_server/model"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := model.User{}
	user.FirstById(id)

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")

	user := model.User{
		Name:	name,
		Password:	password,
	}
	user.Create()

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	name := c.FormValue("name")
	password := c.FormValue("password")

	user := model.User{
		Id:	id,
		Name:	name,
		Password:	password,
	}
	user.Updates()

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := model.User{}
	user.DeleteById(id)

	return c.JSON(http.StatusOK, user)
}