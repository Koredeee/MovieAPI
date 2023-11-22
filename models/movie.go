package models

import "time"

type (
	Movie struct {
		// gorm.Model
		ID                  int               `gorm:"primary_key" json:"id"`
		Title               string            `json:"title"`
		Year                int               `json:"year"`
		AgeRatingCategoryID uint              `json:"ageRatingCategoryID"`
		CreatedAt           time.Time         `json:"created_at"`
		UpdatedAt           time.Time         `json:"updated_at"`
		AgeRatingCategory   AgeRatingCategory `json:"-"`
	}
)
