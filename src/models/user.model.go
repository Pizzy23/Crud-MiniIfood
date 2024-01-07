package models

type User struct {
	IdUser  uint   `gorm:"column:idUser;primaryKey;autoIncrement" json:"id"`
	Name    string `gorm:"column:name;not null" json:"name"`
	OrderID uint   `gorm:"column:order_id;null"`
	Order   Order  `gorm:"foreignKey:OrderID"`
	LoginID uint   `gorm:"column:login_id;null"`
	Login   Login  `gorm:"foreignKey:LoginID"`
}
