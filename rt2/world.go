package rt2

import "image/color"

type World struct {
	Color  color.RGBA
	Lights []Light
	Shapes []Shape
}

func NewWorld(c color.RGBA) World {
	return World{Color: c}
}

func (w *World) AddLight(l Light) {
	w.Lights = append(w.Lights, l)
}

func (w *World) AddShape(s Shape) {
	w.Shapes = append(w.Shapes, s)
}

func (w *World) At(v Vec) color.RGBA {
	for _, light := range w.Lights {
		if v.Dist(&light.Point) < 10 {
			return color.RGBA{255, 255, 255, 255}
		}
	}
	return w.Color
}
