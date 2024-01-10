package main

import (
	"stand_up/internal/core/domain"
	"stand_up/internal/database"
)

func Migrate() {
	db := database.ConnectDB()
	db.AutoMigrate(&domain.User{}, &domain.UserTask{}, &domain.StandUpdate{}, &domain.Sprint{},
		&domain.Task{})
}

func main() {
	Migrate()
}
