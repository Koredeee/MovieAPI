package controllers

import (
	"MovieAPI/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AgeRatingCategoryInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Get All Rating godoc
// @Summary Get All Age Rating Category
// @Description Get List of Age Rating Category
// @Tags AgeRatingCategory
// @Produce json
// @Success 200 {object} []models.AgeRatingCategory
// @Router /age-rating-categories [get]
func GetAllRating(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var ratings []models.AgeRatingCategory

	db.Find(&ratings)

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": ratings})
}

// Create a Rating godoc
// @Summary Create Age Rating Category
// @Description Create new Age Rating Category
// @Tags AgeRatingCategory
// @Param Body body AgeRatingCategoryInput true "the body to create new age rating category"
// @Produce json
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories [post]
func CreateRating(c *gin.Context) {
	var input AgeRatingCategoryInput

	// untuk mengecek apakah json sesuai inputan atau tidak
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rating := models.AgeRatingCategory{Name: input.Name, Description: input.Description}

	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&rating)

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": rating})

}

// Get a Rating godoc
// @Summary Get an Age Rating Category by Id
// @Description Get one Age Rating category by Id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Eating Category Id"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [get]
func GetRatingById(c *gin.Context) {

	var rating models.AgeRatingCategory
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// mengecek hasil db. harus masukkin parameter id juga
	// First -> mengeluarkan data pertama
	if err := db.Where("age_rating_category_id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Get movies from one Rating godoc
// @Summary Get movies by Age Rating Category by id
// @Description Get all movies of Age Rating Category by id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Eating Category Id"
// @Success 200 {object} []models.Movie
// @Router /age-rating-categories/{id}/movies [get]
func GetMoviesByAgeRatingCategoryId(c *gin.Context) {

	// kalo outputnya list brarti plural
	var movies []models.Movie
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// sama kaya di atas cuma bedanya -> .Find(&movies)
	if err := db.Where("age_rating_category_id = ?", c.Param("id")).Find(&movies).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// Update Rating godoc
// @Summary Update an Age Rating Category by Id
// @Description Update one Age Rating category by Id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Eating Category Id"
// @Param Body body AgeRatingCategoryInput true "the body to create new age rating category"
// @Success 200 {object} models.AgeRatingCategory
// @Router /age-rating-categories/{id} [patch]
func UpdateRating(c *gin.Context) {

	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// get rating if exist
	var rating models.AgeRatingCategory
	if err := db.Where("age_rating_category_id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// validasi input
	var input AgeRatingCategoryInput

	// untuk mengecek apakah json sesuai inputan atau tidak
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// bedanya dengan Create
	// di sini ada validasi update
	var updatedInputRating models.AgeRatingCategory

	updatedInputRating.Name = input.Name
	updatedInputRating.Description = input.Description
	updatedInputRating.UpdatedAt = time.Now()

	// masukkin ke db nya
	// .Updates() itu bawaan gorm
	db.Model(&rating).Updates(updatedInputRating)

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": rating})
}

// Delete a Rating godoc
// @Summary Delete an Age Rating Category by Id
// @Description Delete one Age Rating category by Id
// @Tags AgeRatingCategory
// @Produce json
// @Param id path string true "Age Eating Category Id"
// @Success 200 {object} map[string]boolean
// @Router /age-rating-categories/{id} [delete]
func DeleteRating(c *gin.Context) {

	var rating models.AgeRatingCategory
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// mengecek hasil db. harus masukkin parameter id juga
	// First -> mengeluarkan data pertama
	if err := db.Where("age_rating_category_id = ?", c.Param("id")).First(&rating).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// bedanya dari GetRatingById
	// Delete cuma nambah .Delete()
	db.Delete(&rating)

	// return statusOK dan return true apabila berhasil di-delete
	c.JSON(http.StatusOK, gin.H{"data": true})
}
