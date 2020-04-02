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

func ReadImage(path string) (*challenge.ImageData, error) {
	imageType := strings.SplitAfter(path, ".")
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	img, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}
	resizedImg := resize.Resize(256, 256, img, resize.Lanczos3)

	buffer := new(bytes.Buffer)
	err = jpeg.Encode(buffer, resizedImg, nil)
	if err != nil {
		return nil, err
	}
	return &challenge.ImageData{
		Type: imageType[len(imageType)-1],
		Data: buffer.Bytes(),
	}, nil
}

func GetUUID() string {
	return uuid.New().String()
}
