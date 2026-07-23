package style

import "fmt"

type BoxShadow struct {
	X      Size
	Y      Size
	Blur   Size
	Spread Size
	Color  Color
}

func Shadow(
	x Size,
	y Size,
	blur Size,
	spread Size,
	color Color,
) BoxShadow {
	return BoxShadow{
		X:      x,
		Y:      y,
		Blur:   blur,
		Spread: spread,
		Color:  color,
	}
}

func (s BoxShadow) String() string {
	if s.X == "" &&
		s.Y == "" &&
		s.Blur == "" &&
		s.Spread == "" &&
		s.Color == "" {
		return ""
	}

	return fmt.Sprintf(
		"%s %s %s %s %s",
		s.X.String(),
		s.Y.String(),
		s.Blur.String(),
		s.Spread.String(),
		s.Color.string(),
	)
}
