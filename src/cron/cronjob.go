package cron

import (
	"log"
	"super_crud/src/common/cloudinary"
	"super_crud/src/models"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

func StartCronJob(db *gorm.DB) {
	go func() { // Run in a Goroutine
		c := cron.New()
		_, err := c.AddFunc("@every 1h", func() {
			log.Println("Cron job running...")
			deleteExpiredFiles(db)
		})
		if err != nil {
			log.Fatalf("Error scheduling cron job: %v", err)
		}

		c.Start()
	}()
}

func deleteExpiredFiles(db *gorm.DB) {
	var expiredFiles []models.File

	err := db.Where("expires_at < ?", time.Now()).Find(&expiredFiles).Error
	if err != nil {
		log.Printf("Error fetching expired files: %v", err)
		return
	}

	if len(expiredFiles) == 0 {
		log.Println("No Expired Files found")
		return
	}

	var wg sync.WaitGroup

	for _, file := range expiredFiles {
		wg.Add(1)

		go func(file models.File) {
			defer wg.Done()

			err := cloudinary.DeleteFromCloudinary(file.PublicID)
			if err != nil {
				log.Printf("Error deleting files from Cloudinary for file %s: %v", file.FileName, err)
				return
			}

			err = db.Delete(&file).Error
			if err != nil {
				log.Printf("Error deleting file %s from database: %v", file.FileName, err)
			} else {
				log.Printf("Deleted expired file: %s", file.FileName)
			}
		}(file)
	}

	wg.Wait() 
}
