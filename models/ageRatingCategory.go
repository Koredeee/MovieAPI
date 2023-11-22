package models

import (
	"gorm.io/gorm"
)

type (
	AgeRatingCategory struct {
		gorm.Model
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Movies      []Movie `json:"-"`
	}
)
