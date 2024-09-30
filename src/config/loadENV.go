package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	CLOUDINARY_CLOUD_NAME  string
	CLOUDINARY_API_KEY string
	CLOUDINARY_API_SECRET string
	PORT 		string
	HOST string

	DB_Host                 string
	DB_User                 string
	DB_Password             string
	DB_Name                 string
	DB_Port                 string
)

func LoadConfig() { 
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error on loading .env file")
	}
	CLOUDINARY_CLOUD_NAME = os.Getenv("CLOUDINARY_CLOUD_NAME")
	CLOUDINARY_API_KEY = os.Getenv("CLOUDINARY_API_KEY")
	CLOUDINARY_API_SECRET = os.Getenv("CLOUDINARY_API_SECRET")

	PORT = os.Getenv("PORT")
	HOST = os.Getenv("HOST")
	DB_Host = os.Getenv("DB_HOST")
	DB_User = os.Getenv("DB_USER")
	DB_Password = os.Getenv("DB_PASSWORD")
	DB_Name = os.Getenv("DB_NAME")
	DB_Port = os.Getenv("DB_PORT")
}
