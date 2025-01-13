package helpers

import (
	"context"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/labstack/echo/v4"
)

type CloudinaryInterface interface {
	CloudinaryUpload(c echo.Context, fileheader string) string
}

type cloudinaryM struct {
	cloudinaryKey string
}

func NewCloudninary(cloudinaryKey string) CloudinaryInterface {
	return &cloudinaryM{
		cloudinaryKey: cloudinaryKey,
	}
}


func(cloud *cloudinaryM)CloudinaryUpload(c echo.Context, fileheader string) string {
	fileHeader, _ := c.FormFile(fileheader)
	file, _ := fileHeader.Open()
	ctx := context.Background()
	urlCloudinary := cloud.cloudinaryKey
	cldService, _ := cloudinary.NewFromURL(urlCloudinary)
	response, _ := cldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: "Photo"})
	return response.SecureURL
}
