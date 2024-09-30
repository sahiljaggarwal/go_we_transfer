package controllers

import (
	"log"
	"super_crud/src/common"
	"super_crud/src/common/cloudinary"
	"super_crud/src/common/dto"
	"super_crud/src/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FileController struct {
	DB *gorm.DB
}

func (fc *FileController) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error":"No File Provides",
		})
	}

	f, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":"failed to open file",
		})
	}
	defer f.Close()

	resultChan := make(chan *models.File, 1)
	errorChan := make(chan error, 1)

	go func (){
	resp, err := cloudinary.UploadToCloudinary(f, file.Filename)
	if err != nil {
		errorChan <- err
		return
	}

	// id := uuid.New()
	shortId := common.GenerateShortID(8)
	downloadLink := shortId
	fileRecord := models.File {
		// ID: id,
		FileName:     file.Filename,
		FileURL:      resp.SecureURL,
		FileType:     file.Header.Get("Content-Type"),
		Size:         file.Size,
		ExpiresAt:    time.Now().Add(2 * time.Hour), // File will expire after 2 hours
		DownloadLink: downloadLink,
		PublicID:     resp.PublicID,
		ShortID: shortId,
	}
	err = fc.DB.Create(&fileRecord).Error
	if err != nil {
		errorChan <- err
		return
	}
	resultChan <- &fileRecord
}()

select {
case fileRecord := <-resultChan:
	response  := dto.ToFilePrivateResponse(fileRecord)
		return c.Status(200).JSON(fiber.Map{
		"message":"File Uploaded successfully",
		"data":response,
	})
case err := <-errorChan:
	log.Fatal("Error during file upload: ", err)
	return c.Status(500).JSON(fiber.Map{
			"error": "Failed to upload file and save details",
		})
} 
}

func (fc *FileController) GetFileByShortId (c *fiber.Ctx) error {
	shortId := c.Params("id")
	log.Print(shortId)

	var fileRecord models.File
	err := fc.DB.Where("short_id = ?", shortId).First(&fileRecord).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":"File not found",
		})
	}
	if time.Now().After(fileRecord.ExpiresAt) {
		return c.Status(410).JSON(fiber.Map{
			"error":"file has expired",
		})
	}
	return c.Redirect(fileRecord.FileURL, 302)
}

func (fc *FileController) DeleteFileById (c *fiber.Ctx) error {
	fileId := c.Params("id")

	var fileRecord models.File

	err := fc.DB.First(&fileRecord, "id = ?", fileId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error":"File not found",
		})
	}

	cloudinaryErrorChan := make(chan error, 1)
	databaseErrorChan := make(chan error, 1)
	successChan := make(chan bool, 2)

	go func(){
		err := cloudinary.DeleteFromCloudinary(fileRecord.PublicID)
		if err != nil {
			cloudinaryErrorChan <-err
			return 
		}
		successChan <- true
	}()

	go func(){
		err := fc.DB.Delete(&fileRecord).Error
		if err != nil {
			databaseErrorChan <- err
			return
		}
		successChan <- true
	}()

	select {
	case err := <-cloudinaryErrorChan:
		log.Fatal("Error during Cloudinary file deleteion: ", err)
		return c.Status(500).JSON(fiber.Map{
			"error":"Failed  to delete file from cloudinary",
		})
	case err := <-databaseErrorChan:
		log.Fatal("Error during database file deleteion: ",err)
		return c.Status(500).JSON(fiber.Map{
			"error":"Failed to delete file from the database",
		})
	case <-successChan:
		return c.Status(200).JSON(fiber.Map{
			"message":"File deleted successfully from Cloudinary and the database",
		})
	}
}

func (fc *FileController) GetAllFiles (c *fiber.Ctx) error {
	var files []models.File

	err := fc.DB.Find(&files).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":"Failed to fetch files from the database",
		})
	}

	if len(files) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message":"No Files Found",
		})
	}

	var fileResponses []any
	for _, file := range files {
		fileResponse  := dto.ToFilePrivateResponse(&file)
		fileResponses = append(fileResponses, *&fileResponse)

	}

	return c.Status(200).JSON(fiber.Map{
		"message":"Files retrieved successfully",
		"data":fileResponses,
		"count":len(fileResponses),
	})

}

