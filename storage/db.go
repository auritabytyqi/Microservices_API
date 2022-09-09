package storage

import (
	"log"

	"Microservices_API/FOODS_MS/model"
	model1 "Microservices_API/RESTAURANTS_MS/model"
	config "Microservices_API/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	var err error
	conString := config.GetMySQLConnectionString()
	DB, err = gorm.Open(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}
	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}

func AddFoodRecord(food model.Food) error {
	DB.Table("Foods").Select("FoodId", "FoodName", "FoodDescription").Create(&food)
	return nil
}

func AddRestaurantRecord(restaurant model1.Restaurant) error {
	DB.Table("Restaurants").Select("RestaurantId", "RestaurantName", "RestaurantDescription", "FoodIdW").Create(&restaurant)
	return nil
}

func DeleteFood(id string) error {
	DB.Table("Foods").Delete(&model.Food{}, id)
	return nil
}

func DeleteRestaurant(id string) error {
	DB.Table("Restaurants").Delete(&model1.Restaurant{}, id)
	return nil
}

func UpdateFood(id, name, description string) error {
	DB.Table("Foods").Model(&model.Food{}).Where("id=?", id).Updates(model.Food{FoodName: name, FoodDescription: description})
	return nil
}

func UpdateRestaurant(id, name, description, foodId string) error {
	DB.Table("Restaurants").Model(&model1.Restaurant{}).Where("id=?", id).Updates(model1.Restaurant{RestaurantName: name, RestaurantDescription: description, FoodId: foodId})
	return nil
}
