package repository

import (
	"icomida/src/models"

	"gorm.io/gorm"
)

func FindFood(q *gorm.DB, data map[string]interface{}) (map[string]interface{}, error) {
	food := models.Food{
		Name: data["name"].(string),
	}
	result := map[string]interface{}{}
	err := q.Model(&food).First(&result).Error
	return result, err
}

func AddFood(q *gorm.DB, data map[string]interface{}) (*gorm.DB, error) {
	food := models.Food{
		Name:        data["name"].(string),
		Description: data["description"].(string),
	}
	result := q.Create(&food)

	return result, result.Error

}
