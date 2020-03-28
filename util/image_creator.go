package util

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func CreateTextImage(text string, textSize int) (image.Image, error) {
	ctx := createAndFill(textSize, len(text))
	face, err := getFontFace(textSize)
	if err != nil {
		return nil, err
	}
	ctx.SetFontFace(face)

	// write characters to image
	for i, char := range text {
		ctx.Push()
		drawCharacter(char, textSize, i, ctx)
		ctx.Pop()
	}
	// draw obstacles
	for i := 0; i < ctx.Width(); i += textSize / 2 {
		ctx.Push()
		drawCircles(ctx, i, textSize)
		ctx.Pop()
		ctx.Push()
		drawLines(ctx, i, textSize)
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
	position int,
	ctx *gg.Context) {
	ctx.SetRGB(
		float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255)))
	ctx.RotateAbout(
		gg.Degrees(float64(rand.Intn(360))),
		float64(position*textSize+(textSize/2)),
		float64(ctx.Height()-(textSize/2)))
	ctx.DrawString(
		string(character),
		float64(position*textSize+(textSize/3)),
		float64(ctx.Height()-(textSize/3)))
}

func drawCircles(ctx *gg.Context, i int, textSize int) {
	ctx.SetRGBA(
		float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(180)),
	)
	ctx.DrawCircle(
		float64(i),
		float64(rand.Intn(ctx.Height())),
		float64(rand.Intn(textSize-(textSize/4))))
	ctx.Fill()
}

func drawLines(ctx *gg.Context, i int, textSize int) {
	ctx.SetRGBA(
		float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255)),
		float64(rand.Intn(255)),
	)
	x1 := i-5
	if x1 < 0 {
		x1 = 0
	}
	x2 := rand.Intn(i+5)
	if x2 > ctx.Width() {
		x2 = ctx.Width()
	}
	ctx.SetLineWidth(float64(rand.Intn(6)))
	ctx.DrawLine(
		float64(x1),
		float64(rand.Intn(ctx.Height())),
		float64(x2),
		float64(rand.Intn(ctx.Height())),
	)
	ctx.Stroke()
}


func createAndFill(textSize int, length int) *gg.Context {
	ctx := gg.NewContext((textSize*length)+(textSize/4),textSize+(textSize/2))
	ctx.DrawRectangle(0, 0, float64(ctx.Width()), float64(ctx.Height()))
	ctx.SetRGB(0, 0, 0)
	ctx.Stroke()
	ctx.Clear()
	return ctx
}