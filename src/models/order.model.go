package models

type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Confirmed OrderStatus = "confirmed"
	Shipped   OrderStatus = "shipped"
	Delivered OrderStatus = "delivered"
	Cancelled OrderStatus = "cancelled"
)

type Order struct {
	ID           uint        `gorm:"column:idOrder;primaryKey;autoIncrement" json:"id"`
	OrderNumber  int         `gorm:"column:orderNumber;not null" json:"orderNumber"`
	Status       OrderStatus `gorm:"column:status" json:"status"`
	Items        []string    `gorm:"column:items" json:"items"`
	Total        int         `gorm:"column:total json:total"`
	RestaurantID uint        `gorm:"column:restaurant_id" json:"restaurant_id"`
	Restaurant   *Restaurant `gorm:"foreignKey:RestaurantID"`
}
