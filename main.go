package main

import "github.com/labstack/echo"

func main() {
	srv := echo.New()

	apiGroup := srv.Group("/api")
	apiV1Group := apiGroup.Group("/v1")
	apiV1Group.GET("test", func(context echo.Context) error {
		return context.JSON(200, "Ok")
	})

	srv.Start(":8149")
}
