package main

import "MovieAPI/config"

func main() {
	db := config.ConnectDatabase()
	sqlDB, _ := db.DB()

	defer sqlDB.Close()
}
