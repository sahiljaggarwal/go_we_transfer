package database

import (
	"fmt"
	"log"
	"super_crud/src/config"
	"super_crud/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_Host,
		config.DB_User,
		config.DB_Password,
		config.DB_Name,
		config.DB_Port,
	)
	var err error
	DB,err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error on connecting database: ", err)
		return
	}
	err = DB.AutoMigrate(&models.User{}, &models.File{})
	if err != nil {
		log.Fatal("Failed to auto migrate the user model: ", err)
	}
	fmt.Println("Database connected successfully.")

}