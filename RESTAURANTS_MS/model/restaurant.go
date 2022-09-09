package model

type Restaurant struct {
	RestaurantId          string `json:"RestaurantId"`
	RestaurantName        string `json:"RestaurantName"`
	RestaurantDescription string `json:"RestaurantDescription"`
	FoodId                string `json:"FoodId"`
}
