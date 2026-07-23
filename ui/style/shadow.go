package style

import "fmt"

type Shadow struct {
	Inset  bool
	X      Size
	Y      Size
	Blur   Size
	Spread Size
	Color  Color
}

type shadows struct{}

var Shadows shadows

func (shadows) Box(
	x Size,
	y Size,
	blur Size,
	spread Size,
	color Color,
) Shadow {
	return Shadow{
		X:      x,
		Y:      y,
		Blur:   blur,
		Spread: spread,
		Color:  color,
	}
}

func (shadows) None() Shadow {
	return Shadow{}
}

func (shadows) Inset(
	x, y, blur, spread Size,
	color Color,
) Shadow {
	return Shadow{
		Inset:  true,
		X:      x,
		Y:      y,
		Blur:   blur,
		Spread: spread,
		Color:  color,
	}
}

func (s Shadow) string() string {
	if s.X == "" &&
		s.Y == "" &&
		s.Blur == "" &&
		s.Spread == "" &&
		s.Color == "" {
		return ""
	}

	prefix := ""
	if s.Inset {
		prefix = "inset "
	}

	return fmt.Sprintf(
		"%s%s %s %s %s %s",
		prefix,
		s.X.string(),
		s.Y.string(),
		s.Blur.string(),
		s.Spread.string(),
		s.Color.string(),
	)
}
