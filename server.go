package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/cms/banners/", func(c echo.Context) error {
		language := c.QueryParam("language")
		pageId := c.QueryParam("pageId")
		deviceId := c.QueryParam("deviceId")

		return c.String(http.StatusOK, GetBanners(language, pageId, deviceId))
	})

	e.Logger.Fatal(e.Start(":1323"))
}
