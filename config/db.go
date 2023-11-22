package config

import (
	"MovieAPI/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// digunakan untuk config (kalo mau konek harus panggil ini dulu)
func ConnectDatabase() *gorm.DB {
	username := "root"
	password := "root123"
	host := "tcp(127.0.0.1:3306)"
	database := "db_movieAPI"

	dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic((err.Error()))
	}

	// connect database by the models
	db.AutoMigrate(&models.Movie{}, &models.AgeRatingCategory{})

	return db
}
