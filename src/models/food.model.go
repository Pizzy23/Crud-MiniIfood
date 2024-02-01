package models

type Food struct {
	IdFood       uint64      `gorm:"column:idFood;primaryKey;AutoIncrement" json:"id"`
	Name         string      `gorm:"column:name;not null" json:"name"`
	Description  string      `gorm:"column:description;not null" json:"description"`
	Price        float64     `gorm:"column:price"`
	RestaurantID uint        `gorm:"column:restaurant_id;not null" json:"restaurant_id"`
	Restaurant   *Restaurant `gorm:"foreignKey:RestaurantID"`
}
