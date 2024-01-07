package models

type Food struct {
	IdFood       uint64      `gorm:"column:idFood;primaryKey;AutoIncrement" json:"id"`
	Name         string      `gorm:"column:name;not null" json:"name"`
	Description  string      `gorm:"column:description;not null" json:"description"`
	RestaurantID uint        `gorm:"column:restaurant_id;not null" json:"restaurant_id"`
	Restaurant   *Restaurant `gorm:"foreignKey:RestaurantID"`
}
