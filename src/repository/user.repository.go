package repository

import (
	"icomida/src/models"

	"gorm.io/gorm"
)

func FindUserByEmail(q *gorm.DB, email string) (models.User, error) {
	var user models.User
	if err := q.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
