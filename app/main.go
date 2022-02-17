package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)
func getHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello !World!")
}

func main() {
	e := echo.New()  // echo を利用する
    // GET リクエストでパスが `/` のとき第２引数の関数を実行する
	e.GET("/",getHello )

    // 1323 ポートでリッスンを開始。 start がエラーを起こしたら Fatal を起こしてログに記録する
	e.Logger.Fatal(e.Start(":1323"))
}