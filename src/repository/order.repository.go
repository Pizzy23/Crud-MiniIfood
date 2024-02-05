package repository

import (
	"gorm.io/gorm"
)

func FindAllOrderUser(q *gorm.DB)       {}
func FindAllOrderRestaurant(q *gorm.DB) {}
func FindOrder(q *gorm.DB)              {}
func CreateOrder(q *gorm.DB)            {}
func ChangeStatusOrder(q *gorm.DB)      {}
