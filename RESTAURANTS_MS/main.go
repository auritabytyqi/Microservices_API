package main

import (
	"Microservices_API/RESTAURANTS_MS/controller"
	"Microservices_API/storage"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.NewDB()

	e.GET("/", start)
	e.GET("request/api/restaurants", controller.GetRestaurant)
	e.GET("request/api/restaurants/:id", controller.GetRestaurants)
	e.GET("request/api/AddRestaurant", controller.AddRestaurant)
	e.GET("request/api/DeleteRestaurant/:id", controller.DeleteRestaurant)
	e.GET("request/api/UpdateRestaurant/values", controller.UpdateRestaurant)
	// Start server
	e.Logger.Fatal(e.Start(":1324"))
}

func start(c echo.Context) error {
	return c.String(http.StatusOK, "Starting app........ ")
}
