package helper

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImage(image []byte, id, fileName string) (string, error) {

	name := os.Getenv("CLOUD_NAME")
	key := os.Getenv("CLOUD_API_KEY")
	secret := os.Getenv("CLOUD_API_SECRET")

	cld, _ := cloudinary.NewFromParams(name, key, secret)
	ctx := context.Background()
	// Convert []byte to io.Reader
	imageReader := bytes.NewReader(image)
	publicId := fmt.Sprintf("%s/%s", fileName, id)
	resp, err := cld.Upload.Upload(ctx, imageReader, uploader.UploadParams{PublicID: publicId})
	if err != nil {
		return "", err
	}
	log.Println(resp)
	return resp.SecureURL, nil
}
