package util

import (
	"github.com/google/uuid"
	"image"
	"os"
)

func ContainsUInt(a []uint, x uint) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func ReadImage(path string) (image.Image, error) {
	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer imageFile.Close()

	imageData, _, err := image.Decode(imageFile)
	if err != nil {
		return nil, err
	}
	return imageData, nil
}

func GetUUID() string {
	return uuid.New().String()
}
