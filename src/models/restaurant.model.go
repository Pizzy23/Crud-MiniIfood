package models

type Restaurant struct {
	IdRestaurant uint   `gorm:"column:idRestaurant;primary_key;AutoIncrement" json:"id"`
	Name         string `gorm:"column:name;not null" json:"name"`
	Score        int    `gorm:"column:score;not null" json:"score"`
}
