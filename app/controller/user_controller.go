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
	user.FirstByUserId(id)

	return c.JSON(http.StatusOK, user)
}

func GetAllUser(c echo.Context) error {
	users:=model.GetAllUser()
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	password := c.FormValue("password")

	user := model.User{
		Name:	name,
		Password:	password,
	}
	user.CreateUser()

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
	user.UpdatesUser()

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := model.User{}
	user.DeleteByUserId(id)

	return c.JSON(http.StatusOK, user)
}