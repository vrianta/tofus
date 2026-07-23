package style

type Border struct {
	Width Size
	Style string
	Color Color
}

type borders struct{}

var Borders borders

// premade borders
var borderNone = Border{
	Width: "0px",
	Style: "none",
	Color: Colors.Black(),
}

func (borders) None() Border {
	return borderNone
}

func (borders) Solid(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "solid",
		Color: color,
	}
}

func (borders) Dashed(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "dashed",
		Color: color,
	}
}

func (borders) Dotted(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "dotted",
		Color: color,
	}
}

func (borders) Double(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "double",
		Color: color,
	}
}

func (borders) Groove(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "groove",
		Color: color,
	}
}

func (borders) Ridge(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "ridge",
		Color: color,
	}
}

func (borders) Inset(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "inset",
		Color: color,
	}
}

func (borders) Outset(width Size, color Color) Border {
	return Border{
		Width: width,
		Style: "outset",
		Color: color,
	}
}

type BorderRadius struct {
	TopLeft     Size
	TopRight    Size
	BottomLeft  Size
	BottomRight Size
}

var BorderRadiuses = BorderRadius{}

func (b BorderRadius) SetAll(v Size) BorderRadius {
	b.TopLeft = v
	b.TopRight = v
	b.BottomLeft = v
	b.BottomRight = v
	return b
}

func (b BorderRadius) SetTopLeft(v Size) BorderRadius {
	b.TopLeft = v
	return b
}

func (b BorderRadius) SetTopRight(v Size) BorderRadius {
	b.TopRight = v
	return b
}

func (b BorderRadius) SetBottomLeft(v Size) BorderRadius {
	b.BottomLeft = v
	return b
}

func (b BorderRadius) SetBottomRight(v Size) BorderRadius {
	b.BottomRight = v
	return b
}

func (b BorderRadius) SetTop(v Size) BorderRadius {
	b.TopLeft = v
	b.TopRight = v
	return b
}

func (b BorderRadius) SetBottom(v Size) BorderRadius {
	b.BottomLeft = v
	b.BottomRight = v
	return b
}

func (b BorderRadius) SetLeft(v Size) BorderRadius {
	b.TopLeft = v
	b.BottomLeft = v
	return b
}

func (b BorderRadius) SetRight(v Size) BorderRadius {
	b.TopRight = v
	b.BottomRight = v
	return b
}
