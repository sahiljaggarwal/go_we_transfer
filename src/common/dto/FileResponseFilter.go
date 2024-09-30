package dto

import (
	"super_crud/src/common"
	"super_crud/src/models"
)

type FileResponse struct {
	Id           uint `json:"id"`            // UUID or string ID
	FileName     string `json:"file_name"`     // Name of the file
	FileURL      string `json:"file_url"`      // URL to access the file
	FileType     string `json:"file_type"`     // MIME type of the file
	Size         int64  `json:"size"`          // Size of the file in bytes
	ExpiresAt    string `json:"expires_at"`    // Expiration time of the file
	DownloadLink string `json:"download_link"` // Link for downloading the file
	PublicID     string `json:"public_id"`     // Public ID for the file (if applicable)
	ShortID      string `json:"short_id"`      // Shortened identifier for the file
}

func ToFilePublickResponse(file *models.File) FileResponse {
	return FileResponse{
		Id:          file.ID,
		FileName:     file.FileName,
		FileURL:      file.FileURL,
		FileType:     file.FileType,
		Size:         file.Size,
		ExpiresAt:    file.ExpiresAt.Format("2006-01-02 15:04:05"), // Format as needed
		DownloadLink: common.GenerateHostURL(file.ShortID),
		PublicID:     file.PublicID,
		ShortID:      file.ShortID,
	}
}

func ToFilePrivateResponse(file *models.File) FileResponse {
	return FileResponse{
		Id:          file.ID,
		FileName:    file.FileName,
		FileType:    file.FileType,
		Size:        file.Size,
		DownloadLink: common.GenerateHostURL(file.ShortID),
		ExpiresAt:   file.ExpiresAt.Format("2006-01-02 15:04:05"),
		PublicID:     file.PublicID,
		ShortID:      file.ShortID,
	}
}