package models

type Login struct {
	IdLogin  uint64 `gorm:"column:idLogin;primary_key;AutoIncrement" json:"id"`
	IsLogged bool   `gorm:"column:isLogged;not null" json:"isLogged"`
}
