package models

import (
	"time"

	"gorm.io/gorm"
)

type File struct {
	
	gorm.Model
	FileName 	string `json:"file_name"`
	FileURL 	string `json:"file_url"`
	FileType 	string `json:"file_type"`
	Size 		int64 `json:"size"`
	ExpiresAt 	time.Time `json:"expires_at"`
	DownloadLink string `json:"download_link"`
	PublicID 	string `json:"public_id"`
	ShortID     string    `json:"short_id"`
}