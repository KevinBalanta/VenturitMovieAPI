package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title         string  `json:"title"`
	Released_year uint    `json:"released_year"`
	Rating        float64 `json:"rating"`
	Genres        string  `json:"genres"`
}
