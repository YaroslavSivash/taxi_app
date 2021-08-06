package main

import (
		"github.com/labstack/echo"
)

func main(){
	e := echo.New()

	e.GET("/request", func(c echo.Context) error { return c.JSON(200, "GET REQUEST")})
	e.PUT("/request", func(c echo.Context) error { return c.JSON(200, " PUT REQUEST")})

	e.GET("/admin/request", func(c echo.Context) error { return c.JSON(200, "GET ADMIN REQUEST")})
	e.PUT("/admin/request", func(c echo.Context) error { return c.JSON(200, "PUT ADMIN REQUEST")})


	e.Logger.Fatal(e.Start(":8000"))

}