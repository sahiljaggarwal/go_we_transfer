package routes

import (
	"super_crud/src/controllers"
	"super_crud/src/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func FileRoutes(router fiber.Router, db *gorm.DB){
	fileController := &controllers.FileController{DB:db}
	router.Post("file/upload",middlewares.FileSizeLimit(5 * 1024 * 1024), fileController.UploadFile )
	router.Get("file/:id", fileController.GetFileByShortId)
	router.Get("files", fileController.GetAllFiles)
	router.Delete("file/:id", fileController.DeleteFileById)
}