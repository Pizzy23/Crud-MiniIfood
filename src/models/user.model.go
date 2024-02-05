package models

type User struct {
	IdUser   uint   `gorm:"column:idUser;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:name;not null" json:"name"`
	Email    string `gorm:"column:email;not null" json:"email"`
	Adress   string `gorm:"column:address;not null" json:"address"`
	Password string `gorm:"column:password;not null" json:"password"`
	OrderID  uint   `gorm:"column:order_id;null"`
	Order    Order  `gorm:"foreignKey:OrderID"`
	LoginID  uint   `gorm:"column:login_id;null"`
	Login    Login  `gorm:"foreignKey:LoginID"`
}
