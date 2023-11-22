package controllers

import (
	"MovieAPI/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieInput struct {
	Title               string `json:"title"`
	Year                int    `json:"year"`
	AgeRatingCategoryID int    `json:"age_rating_category_id"`
}

// Get All Rating godoc
// @Summary Get All Movie
// @Description Get List of Movie
// @Tags Movie
// @Produce json
// @Success 200 {object} []models.Movie
// @Router /movies [get]
func GetAllMovie(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	var movies []models.Movie

	db.Find(&movies)

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": movies})
}

// Create a Rating godoc
// @Summary Create Movie
// @Description Create new Movie
// @Tags Movie
// @Param Body body MovieInput true "the body to create new Movie"
// @Produce json
// @Success 200 {object} models.Movie
// @Router /movies [post]
func CreateMovie(c *gin.Context) {
	var input MovieInput

	// untuk mengecek apakah json sesuai inputan atau tidak
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{Title: input.Title, Year: input.Year, AgeRatingCategoryID: uint(input.AgeRatingCategoryID)}

	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	db.Create(&movie)

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": movie})

}

// Get a Rating godoc
// @Summary Get an Movie by Id
// @Description Get one Movie by Id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie Id"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [get]
func GetMovieById(c *gin.Context) {

	var movie models.Movie
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// mengecek hasil db. harus masukkin parameter id juga
	// First -> mengeluarkan data pertama
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Update Rating godoc
// @Summary Update an Movie by Id
// @Description Update one Movie by Id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie Id"
// @Param Body body MovieInput true "the body to create new Movie"
// @Success 200 {object} models.Movie
// @Router /movies/{id} [patch]
func UpdateMovie(c *gin.Context) {

	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// get movie if exist
	var movie models.Movie
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// validasi input
	var input MovieInput

	// untuk mengecek apakah json sesuai inputan atau tidak
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// bedanya dengan Create
	// di sini ada validasi update
	var updatedInputMovie models.Movie

	updatedInputMovie.Title = input.Title
	updatedInputMovie.Year = input.Year
	updatedInputMovie.AgeRatingCategoryID = uint(input.AgeRatingCategoryID)
	updatedInputMovie.UpdatedAt = time.Now()

	// masukkin ke db nya
	// .Updates() itu bawaan gorm
	db.Model(&movie).Updates(updatedInputMovie)

	// return statusOK dan return bentuk map
	c.JSON(http.StatusOK, gin.H{"data": movie})
}

// Delete a Rating godoc
// @Summary Delete an Movie by Id
// @Description Delete one Movie by Id
// @Tags Movie
// @Produce json
// @Param id path string true "Movie Id"
// @Success 200 {object} map[string]boolean
// @Router /movies/{id} [delete]
func DeleteMovie(c *gin.Context) {

	var movie models.Movie
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)

	// mengecek hasil db. harus masukkin parameter id juga
	// First -> mengeluarkan data pertama
	if err := db.Where("id = ?", c.Param("id")).First(&movie).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// bedanya dari GetRatingById
	// Delete cuma nambah .Delete()
	db.Delete(&movie)

	// return statusOK dan return true apabila berhasil di-delete
	c.JSON(http.StatusOK, gin.H{"data": true})
}
