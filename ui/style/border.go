package style

import "fmt"

type Border struct {
	Width Size
	Style string
	Color Color
}

type BorderRadius struct {
	TopLeft     Size
	TopRight    Size
	BottomLeft  Size
	BottomRight Size
}

func (b Border) Solid(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "solid",
		Color: color,
	}
}

func (b BorderRadius) SetAll(s Size) BorderRadius {
	b.TopLeft = s
	b.TopRight = s
	b.BottomLeft = s
	b.BottomRight = s
	return b
}

func (b Border) None() Border {
	return Border{
		Width: Sizes.None,
		Style: "none",
		Color: Colors.White,
	}
}

func (b Border) String() string {
	if b.Width == "" && b.Style == "" && b.Color == "" {
		return ""
	}

	return fmt.Sprintf(
		"%s %s %s",
		b.Width.String(),
		b.Style,
		b.Color.String(),
	)
}

func (b BorderRadius) String() string {
	// all sides equal
	if b.TopLeft == b.TopRight &&
		b.TopLeft == b.BottomRight &&
		b.TopLeft == b.BottomLeft {
		return b.TopLeft.String()
	}

	// top-left/bottom-right and top-right/bottom-left
	if b.TopLeft == b.BottomRight &&
		b.TopRight == b.BottomLeft {
		return fmt.Sprintf(
			"%s %s",
			b.TopLeft.String(),
			b.TopRight.String(),
		)
	}

	return fmt.Sprintf(
		"%s %s %s %s",
		b.TopLeft.String(),
		b.TopRight.String(),
		b.BottomRight.String(),
		b.BottomLeft.String(),
	)
}
