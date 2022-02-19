package router

import (
	"net/http"

	"fail_app_server/controller"

	"github.com/labstack/echo/v4"
)



func Init() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/user", controller.GetAllUser)
	e.GET("/user/:id", controller.GetUser)
	e.POST("/user", controller.CreateUser)
	e.PUT("/user/:id", controller.UpdateUser)
	e.DELETE("/user/:id", controller.DeleteUser)
	e.GET("/fail/:id", controller.GetFail)
	e.POST("/fail", controller.CreateFail)
	e.PUT("/fail/:id", controller.UpdateFail)
	e.DELETE("/fail/:id", controller.DeleteFail)
	//e.GET("/fail", controller.GetAllFail)
	//e.GET("/fail/:user_id", controller.GetUserFail)

	e.Logger.Fatal(e.Start(":1323"))
}