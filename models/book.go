package models

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name"`
	Auther string `json:"auther"`
	City   string `json:"city"`
}
