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

func GetAllFail(c echo.Context) error {
	fails:=model.GetAllFail()
	return c.JSON(http.StatusOK, fails)
}

func GetUserFail(c echo.Context) error {
	u, _ := strconv.Atoi(c.Param("user_id"))
	user_id := uint(u)
	fails:=model.GetUserFail(user_id)
	return c.JSON(http.StatusOK, fails)
}

func CreateFail(c echo.Context) error {
	content := c.FormValue("content")
	print("start!!")
	u := c.FormValue("user_id")
	id,_:=strconv.Atoi(u)
	user_id := uint(id)
	print(user_id)
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