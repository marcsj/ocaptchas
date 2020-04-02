package util

import (
	"bytes"
	"github.com/google/uuid"
	"github.com/marcsj/ocaptchas/challenge"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"os"
	"strings"
)

func ContainsUInt(a []uint, x uint) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ReadImage(path string) (img image.Image, imageType string, err error) {
	imageTypes := strings.SplitAfter(path, ".")
	imageType = imageTypes[len(imageType)-1]
	imageFile, err := os.Open(path)
	if err != nil {
		return
	}
	defer imageFile.Close()

	img, _, err = image.Decode(imageFile)
	return
}

func ConvertImage(img image.Image, imageType string) (*challenge.ImageData, error) {
	resizedImg := resize.Resize(256, 256, img, resize.Lanczos3)

	buffer := new(bytes.Buffer)
	err := jpeg.Encode(buffer, resizedImg, nil)
	if err != nil {
		return nil, err
	}
	return &challenge.ImageData{
		Type: imageType,
		Data: buffer.Bytes(),
	}, nil
}

func GetUUID() string {
	return uuid.New().String()
}
