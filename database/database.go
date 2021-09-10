package database

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating string `json:"rating"`
}

var (
	Conn *gorm.DB
)
