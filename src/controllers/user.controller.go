package controllers

import (
	"net/http"
	"super_crud/src/common/dto"
	"super_crud/src/models"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

// CreateUser handles user registration
func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	var input dto.SignUpInputDTO

	// Parse the input
	if err := c.BodyParser(&input); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

		// Initialize the validator
	validate := validator.New()

	// Validate the input struct
	if err := validate.Struct(input); err != nil {
		// If validation fails, return a 400 Bad Request
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Validation failed",
			"details": err.Error(), // Optionally, you can return more detailed validation errors
		})
	}

	// Check if the username or email already exists
	var existingUser models.User
	if err := uc.DB.Where("email = ?", input.Email).First(&existingUser).Error; err == nil {
		return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Username or email already in use"})
	}

	// Create a new user instance
	user := models.User{
		Email:    input.Email,
		Password: input.Password, // Consider hashing the password before saving
	}

	// Save the user to the database
	if err := uc.DB.Create(&user).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "Failed to create account"})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"data":    user, // Optionally, you might want to format this data with a DTO
	})
}
