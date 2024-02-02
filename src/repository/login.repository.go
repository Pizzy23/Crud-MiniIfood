package repository

import (
	"icomida/src/models"

	"gorm.io/gorm"
)

func ConnectUserInDB(q *gorm.DB, email string) (map[string]interface{}, error) {
	resultUser, errUser := FindUserByEmail(q, email)
	if errUser != nil {
		return nil, errUser
	}
	var login models.Login
	if err := q.Model(&resultUser).Association("Login").Find(&login); err != nil {
		return nil, err
	}
	login.IsLogged = true
	if err := q.Save(&login).Error; err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"message": "Connected successfully",
	}

	return result, nil
}
func DesconnectUserInDB(q *gorm.DB, email string) (map[string]interface{}, error) {

	resultUser, errUser := FindUserByEmail(q, email)
	if errUser != nil {
		return nil, errUser
	}
	var login models.Login
	if err := q.Model(&resultUser).Association("Login").Find(&login); err != nil {
		return nil, err
	}
	login.IsLogged = false
	if err := q.Save(&login).Error; err != nil {
		return nil, err
	}
	result := map[string]interface{}{
		"message": "Desconnected successfully",
	}

	return result, nil
}
