package models

import "time"

type (
	Movie struct {
		// gorm.Model
		ID                  int       `gorm:"primary_key" json:"id"` // menandakan ini primary key
		Title               string    `json:"title"`
		Year                int       `json:"year"`
		AgeRatingCategoryID uint      `json:"age_rating_category_id"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`

		// semacam embedded models tetapi ga perlu dimunculkan jsonnya biar ga numpuk
		AgeRatingCategory AgeRatingCategory `json:"-"`
	}
)
