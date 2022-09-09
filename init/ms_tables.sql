CREATE TABLE Foods(
  FoodId int NOT NULL PRIMARY KEY,
  FoodName nvarchar(255),
  FoodDescription nvarchar(255)
);
CREATE TABLE Restaurants (
  RestaurantId int NOT NULL PRIMARY KEY,
  RestaurantName nvarchar(255),
  RestaurantDescription nvarchar(255),
  FoodId int NOT NULL,
  CONSTRAINT FK_Restaurants_Foods FOREIGN KEY (FoodId)     
    REFERENCES Foods (FoodId)
);