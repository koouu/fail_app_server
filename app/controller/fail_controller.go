package controller

import (
	"net/http"
	"strconv"
	"fail_app_server/model"
	"github.com/labstack/echo/v4"
)

func GetFail(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	fail := model.Fail{}
	fail.FirstByFailId(id)

	return c.JSON(http.StatusOK, fail)
}

func CreateFail(c echo.Context) error {
	content := c.FormValue("content")
	u, _ := strconv.Atoi(c.Param("user_id"))
	user_id := uint(u)

	fail := model.Fail{
		Content:	content,
		User_id:	user_id,
	}
	fail.CreateFail()

	return c.JSON(http.StatusOK, fail)
}

func UpdateFail(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	content := c.FormValue("content")
	u, _ := strconv.Atoi(c.Param("user_id"))
	user_id := uint(u)

	fail := model.Fail{
		Id:	id,
		Content:	content,
		User_id:	user_id,
	}
	fail.UpdatesFail()

	return c.JSON(http.StatusOK, fail)
}

func DeleteFail(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	fail := model.Fail{}
	fail.DeleteByFailId(id)

	return c.JSON(http.StatusOK, fail)
}