package util

import (
	"image/png"
	"os"
	"testing"
)

func TestCreateImage(t *testing.T) {
	img, err := CreateTextImage("test 123 works", 16)
	if err != nil {
		t.Fatal(err)
	}
	f, err := os.Create("test.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		t.Fatal(err)
	}
}
