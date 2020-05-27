package util

import (
	"image/png"
	"log"
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

func TestCreateTextImageRandom(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	txt := RandStringRunes(10)
	log.Println(txt)
	img, err := CreateTextImage(txt, 64)
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
