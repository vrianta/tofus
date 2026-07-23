package style

import "fmt"

type Color string

type colors struct{}

var Colors colors

const (
	colorTransparent Color = "transparent"
	colorWhite       Color = "#ffffff"
	colorBlack       Color = "#000000"
)

func (colors) Transparent() Color {
	return colorTransparent
}

func (colors) White() Color {
	return colorWhite
}

func (colors) Black() Color {
	return colorBlack
}

func (colors) Hex(value string) Color {
	return Color(value)
}

func (colors) RGB(r, g, b uint8) Color {
	return Color(
		fmt.Sprintf("rgb(%d, %d, %d)", r, g, b),
	)
}

func (colors) RGBA(r, g, b uint8, a float64) Color {
	return Color(
		fmt.Sprintf("rgba(%d, %d, %d, %g)", r, g, b, a),
	)
}
