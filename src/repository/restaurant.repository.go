package repository

import (
	"icomida/src/models"

	"gorm.io/gorm"
)

func FindRestaurants(q *gorm.DB, data map[string]interface{}) (map[string]interface{}, error) {
	restaurant := models.Restaurant{
		Name: data["name"].(string),
	}
	result := map[string]interface{}{}
	err := q.Model(&restaurant).First(&result).Error
	return result, err
}

func CreateRestaurant(q *gorm.DB, data map[string]interface{}) (*gorm.DB, error) {
	restaurant := models.Restaurant{
		Name:  data["name"].(string),
		Score: data["score"].(int),
	}
	result := q.Create(&restaurant)

	return result, result.Error

}
func AllRestaurants(q *gorm.DB) (*gorm.DB, error) {
	var restaurants []models.Restaurant
	result := q.Find(&restaurants)

	return result, result.Error

}
