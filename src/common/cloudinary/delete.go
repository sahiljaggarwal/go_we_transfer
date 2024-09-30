package cloudinary

import (
	"context"
	"log"
	"super_crud/src/config"

	"github.com/cloudinary/cloudinary-go/v2"
	// "github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)


func DeleteFromCloudinary(publicId string) error {
	cld, err := cloudinary.NewFromParams(config.CLOUDINARY_CLOUD_NAME, config.CLOUDINARY_API_KEY, config.CLOUDINARY_API_SECRET)
	if err != nil {
		log.Println("Cloudinary error: ", err)
		return err
	}

	_, err = cld.Upload.Destroy(context.Background(), uploader.DestroyParams{
		PublicID: publicId,
	})

	if err != nil {
		log.Fatal("Error on deleting from cloudinary", err)
		return err
	}

	return nil
	
}