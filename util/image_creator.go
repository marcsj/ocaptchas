package util

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"image/color"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func CreateTextImage(text string, textSize int) (image.Image, error) {
	ctx := gg.NewContext((textSize*len(text))+(textSize/4),textSize+(textSize/2))
	ctx.DrawRectangle(0, 0, float64(ctx.Width()), float64(ctx.Height()))
	ctx.SetRGB(0, 0, 0)
	ctx.Stroke()
	ctx.Clear()
	face, err := getFontFace(textSize)
	if err != nil {
		return nil, err
	}
	ctx.SetFontFace(face)
	for i, char := range text {
		ctx.Push()
		rot := rand.Intn(360)
		drawCharacter(char, textSize, color.White, float64(rot), i, ctx)
		ctx.Pop()
	}
	return ctx.Image(), nil
}

func getFontFace(textSize int) (font.Face, error) {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(font, &truetype.Options{Size: float64(textSize)})
	return face, nil
}

func drawCharacter(
	character int32,
	textSize int,
	textColor color.Color,
	rotation float64,
	position int,
	ctx *gg.Context) {
	ctx.SetRGB(
		float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255)))
	ctx.RotateAbout(
		gg.Degrees(rotation),
		float64(position*textSize+(textSize/2)),
		float64(ctx.Height()-(textSize/2)))
	ctx.DrawString(
		string(character),
		float64(position*textSize+(textSize/3)),
		float64(ctx.Height()-(textSize/3)))
}