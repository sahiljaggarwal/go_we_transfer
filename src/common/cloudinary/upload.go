package cloudinary

import (
	"context"
	"log"
	"mime/multipart"
	"super_crud/src/config"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadToCloudinary(file multipart.File, fileName string)(*uploader.UploadResult, error){
	cld, err := cloudinary.NewFromParams(config.CLOUDINARY_CLOUD_NAME, config.CLOUDINARY_API_KEY, config.CLOUDINARY_API_SECRET)
	if err != nil {
		log.Println("Cloudinary error: ", err)
		return nil, err
	}

	uploadParams := uploader.UploadParams{Folder:"share"}
	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		log.Println("Cloudinary upload error ", err)
		return nil, err
	}
	return uploadResult, nil
}