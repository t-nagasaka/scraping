package main

import (
	"github.com/labstack/echo/v4"
	"main/api"
	"net/http"
)

func main() {
	// Echoのインスタンスを作成
	e := echo.New()

	// ルートハンドラ
	e.GET("/coconala", func(c echo.Context) error {
		// Coconalaオブジェクトを作成
		coconala := api.NewCoconala()
		// JSONレスポンスを返す
		title := coconala.FetchRootPage()

		return c.JSON(http.StatusOK, map[string]string{
			"title": title, // Limit output for readability
		})
	})

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
