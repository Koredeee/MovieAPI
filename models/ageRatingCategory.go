package models

import "time"

type (
	AgeRatingCategory struct {
		// gorm.Model
		ID          int       `gorm:"primary_key" json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`

		//Ngebuka listnya
		// untuk menunjukkan adanya relasi
		// agar si migrator tidak bingung bahwa AgeRatingCategory have relationship with Movies
		Movies []Movie `json:"-"`
	}
)
