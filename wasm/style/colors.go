package style

import "fmt"

type Color string

type colors struct {
	Transparent Color
	White       Color
	Black       Color
}

var Colors = colors{
	Transparent: "transparent",
	White:       "#ffffff",
	Black:       "#000000",
}

func (c Color) String() string {
	return string(c)
}

func (c *colors) Hex(value string) Color {
	return Color(value)
}

func (c *colors) RGB(r, g, b uint8) Color {
	return Color(
		fmt.Sprintf("rgb(%d, %d, %d)", r, g, b),
	)
}

func (c *colors) RGBA(r, g, b uint8, a float64) Color {
	return Color(
		fmt.Sprintf("rgba(%d, %d, %d, %g)", r, g, b, a),
	)
}
