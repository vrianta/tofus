package style

import "fmt"

type Outline struct {
	Width Size
	Style string
	Color Color
}

func (o Outline) string() string {
	if o.Style == "" {
		return ""
	}

	if o.Style == "none" {
		return "none"
	}

	return fmt.Sprintf("%s %s %s",
		o.Width.string(),
		o.Style,
		o.Color.string(),
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
func (outlines) Dashed(width Size, color Color) Outline {
	return Outline{
		Width: width,
		Style: "dashed",
		Color: color,
	}
}

func (outlines) Dotted(width Size, color Color) Outline {
	return Outline{
		Width: width,
		Style: "dotted",
		Color: color,
	}
}

func (outlines) Double(width Size, color Color) Outline {
	return Outline{
		Width: width,
		Style: "double",
		Color: color,
	}
}

func (outlines) Custom(style string, width Size, color Color) Outline {
	return Outline{
		Width: width,
		Style: style,
		Color: color,
	}
}
