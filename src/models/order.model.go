package models

type Order struct {
	ID           uint        `gorm:"column:idOrder;primaryKey;autoIncrement" json:"id"`
	OrderNumber  int         `gorm:"column:orderNumber;not null" json:"orderNumber"`
	RestaurantID uint        `gorm:"column:restaurant_id" json:"restaurant_id"`
	Restaurant   *Restaurant `gorm:"foreignKey:RestaurantID"`
}
