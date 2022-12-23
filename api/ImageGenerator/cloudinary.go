package imagegenerator

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/joho/godotenv"
)

func Cloudinary(url string) string {
	fmt.Printf("cloudinary start....%v\n", url)
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	// token := config.Config.ApiKey
	cloudName := os.Getenv("CLOUD_NAME")
	cloudinaryApiKey := os.Getenv("CLOUDINARY_KEY")
	cloudinarySecret := os.Getenv("CLOUDINARY_SECRET")
	// fmt.Println(cloudName, cloudinaryApiKey, cloudinarySecret)
	cld, err := cloudinary.NewFromParams(cloudName, cloudinaryApiKey, cloudinarySecret)
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}

	var ctx = context.Background()

	// Uploading
	uploadResult, err := cld.Upload.Upload(
		ctx,
		url,
		uploader.UploadParams{PublicID: "logo"})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}

	log.Println(uploadResult.SecureURL)
	fmt.Printf("cloudinary generated URL....%v\n", uploadResult.SecureURL)
	return uploadResult.SecureURL
}
