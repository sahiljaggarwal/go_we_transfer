package routes

import (
	"super_crud/src/controllers"

	// "github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoutes(router fiber.Router, db *gorm.DB){
	userController := &controllers.UserController{DB:db}
	router.Post("auth/signup", userController.CreateUser)
}