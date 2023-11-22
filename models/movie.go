package models

import "gorm.io/gorm"

type (
	Movie struct {
		gorm.Model
		Title               string            `json:"title"`
		Year                int               `json:"year"`
		AgeRatingCategoryID uint              `json:"ageRatingCategoryID"`
		AgeRatingCategory   AgeRatingCategory `json:"-"`
	}
)
