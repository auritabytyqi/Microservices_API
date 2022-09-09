package controller

import (
	"log"
	"net/http"

	"Microservices_API/FOODS_MS/controller"
	"Microservices_API/RESTAURANTS_MS/model"
	"Microservices_API/storage"

	"github.com/labstack/echo/v4"
)

func GetRestaurant(c echo.Context) error {
	id := c.Param("id")
	restaurants, _ := GetRepoRestaurants()
	for i := 0; i < len(restaurants); i++ {
		if restaurants[i].RestaurantId == id {
			return c.JSON(http.StatusOK, restaurants[i])
		}
	}
	return c.JSON(http.StatusOK, "Food doesn't exist")
}

func GetRestaurants(c echo.Context) error {
	restaurants, _ := GetRepoRestaurants()
	return c.JSON(http.StatusOK, restaurants)
}

func GetRepoRestaurants() ([]model.Restaurant, error) {
	db := storage.GetDBInstance()
	log.Println(db)
	defer db.Close()
	var restaurants []model.Restaurant
	if err := db.Raw("SELECT Restaurants.RestaurantId, Restaurants.RestaurantName, Restaurants.RestaurantDescription, Foods.FoodName, Foods.FoodDescription FROM Restaurants INNER JOIN Foods ON Restaurants.FoodId=Foods.FoodId").Scan(&restaurants).Error; err != nil {
		return nil, nil
	}
	log.Println(restaurants)
	return restaurants, nil
}

func AddRestaurant(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	description := c.QueryParam("description")
	restaurant_exists := RestaurantExists(id)
	if restaurant_exists {
		return c.JSON(http.StatusOK, "Restaurant exists with this id. Try another one...")

	}
	food_exists := controller.FoodExists(id)
	if !food_exists {
		return c.JSON(http.StatusOK, "Food product doesn't exists with this id. Try another one...")

	}
	restaurant := model.Restaurant{RestaurantId: id, RestaurantName: name, RestaurantDescription: description}
	storage.AddRestaurantRecord(restaurant)
	return c.JSON(http.StatusOK, "Restaurant is created")
}

func DeleteRestaurant(c echo.Context) error {
	id := c.Param(("id"))
	id_exists := RestaurantExists(id)
	if id_exists {
		storage.DeleteRestaurant(id)
		return c.JSON(http.StatusOK, "Restaurant is deleted")
	}
	return c.JSON(http.StatusOK, "Restaurant doesn't exist")
}

func UpdateRestaurant(c echo.Context) error {
	id := c.QueryParam("id")
	name := c.QueryParam("name")
	description := c.QueryParam("description")
	foodId := c.QueryParam("foodId")
	id_exists := RestaurantExists(id)
	if id_exists {
		storage.UpdateRestaurant(id, name, description, foodId)
		return c.JSON(http.StatusOK, "Restaurant is updated")
	}
	return c.JSON(http.StatusOK, "Restaurant doesn't exist")
}

func RestaurantExists(id string) bool {
	restaurant, _ := GetRepoRestaurants()
	id_exists := false
	for i := 0; i < len(restaurant); i++ {
		if restaurant[i].RestaurantId == id {
			id_exists = true
		}
	}
	return id_exists
}
