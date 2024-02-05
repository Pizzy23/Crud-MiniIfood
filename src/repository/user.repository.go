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
func CreateUSer(q *gorm.DB, filters map[string]interface{}) (models.User, error) {
	newUser := models.User{
		Name:     filters["name"].(string),
		Email:    filters["email"].(string),
		Adress:   filters["address"].(string),
		Password: filters["password"].(string),
	}
	if err := q.Create(&newUser).Error; err != nil {
		return newUser, err
	}

	return newUser, nil
}

func EditName(q *gorm.DB, filters map[string]interface{}) error {
	resultUser, errUser := FindUserByEmail(q, filters["email"].(string))
	if errUser != nil {
		return errUser
	}
	result := q.Model(&models.User{}).Where("id = ?", resultUser.IdUser).Update("username", filters["name"].(string))
	if result.Error != nil {
		println("Erro ao atualizar o valor:", result.Error)
		return result.Error
	}
	return nil
}
func EditAdress(q *gorm.DB, filters map[string]interface{}) error {
	resultUser, errUser := FindUserByEmail(q, filters["email"].(string))
	if errUser != nil {
		return errUser
	}
	result := q.Model(&models.User{}).Where("id = ?", resultUser.IdUser).Update("username", filters["adress"].(string))
	if result.Error != nil {
		println("Erro ao atualizar o valor:", result.Error)
		return result.Error
	}
	return nil
}
func EditEmail(q *gorm.DB, email string) error {
	resultUser, errUser := FindUserByEmail(q, email)
	if errUser != nil {
		return errUser
	}
	result := q.Model(&models.User{}).Where("id = ?", resultUser.IdUser).Update("username", email)
	if result.Error != nil {
		println("Erro ao atualizar o valor:", result.Error)
		return result.Error
	}
	return nil
}
func EditPassword(q *gorm.DB, filters map[string]interface{}) error {
	resultUser, errUser := FindUserByEmail(q, filters["email"].(string))
	if errUser != nil {
		return errUser
	}
	result := q.Model(&models.User{}).Where("id = ?", resultUser.IdUser).Update("username", filters["password"].(string))
	if result.Error != nil {
		println("Erro ao atualizar o valor:", result.Error)
		return result.Error
	}
	return nil
}
