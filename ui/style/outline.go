package style

import "fmt"

type Outline struct {
	Width Size
	Style string
	Color Color
}

func (o Outline) String() string {
	if o.Style == "" {
		return ""
	}

	if o.Style == "none" {
		return "none"
	}

	return fmt.Sprintf("%s %s %s",
		o.Width.String(),
		o.Style,
		o.Color.String(),
	)
}

type outlines struct{}

var Outlines outlines

func (outlines) Solid(width Size, color Color) Outline {
	return Outline{
		Width: width,
		Style: "solid",
		Color: color,
	}
}

func (outlines) None() Outline {
	return Outline{
		Style: "none",
	}
}
