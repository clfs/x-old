package main

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/clfs/x/rt2"
)

const (
	imageWidth  = 128
	imageHeight = 72
)

type Render struct {
	w io.Writer
}

func (r *Render) Run() error {
	img := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))
	bounds := img.Bounds()

	world := rt2.NewWorld(color.RGBA{R: 0xEE, G: 0xB4, B: 0xFF, A: 0xFF})
	world.AddLight(rt2.Light{Point: rt2.NewVec(1, 1)})

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			v := rt2.NewVec(float64(x), float64(y))
			img.Set(x, y, world.At(v))
		}
	}

	return png.Encode(r.w, img)
}
