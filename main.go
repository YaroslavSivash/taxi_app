package main

import (
	"net/http"
	"taxi_app/models"

	"github.com/labstack/echo/v4"
)

func main() {
	app := models.NewApplications() //создание пула заявок перед стартом приложения
	e := echo.New()
	e.GET("/request", func(c echo.Context) error {
		return c.String(http.StatusOK, app.GetApp()) // получение одной заявки для таксиста, метод который возвращает строку

	})
	e.GET("/admin/request", func(c echo.Context) error {
		return c.JSON(http.StatusOK, app.GetAllApps()) // получение всех заявок для админов, которые были показаны минимум один разб метод который возвращает слайс application

	})

	e.Logger.Fatal(e.Start(":9000"))
}
