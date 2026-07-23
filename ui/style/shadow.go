package style

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
