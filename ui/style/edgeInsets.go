package style

import "fmt"

type edgeInsets struct {
	Top    Size
	Right  Size
	Bottom Size
	Left   Size
}

var EdgeInsets = edgeInsets{}

func (e edgeInsets) SetAll(v Size) edgeInsets {
	e.Top = v
	e.Bottom = v
	e.Left = v
	e.Right = v

	return e
}

func (e edgeInsets) String() string {
	if e.Top == e.Right && e.Top == e.Bottom && e.Top == e.Left {
		return e.Top.String()
	}

	if e.Top == e.Bottom && e.Right == e.Left {
		return fmt.Sprintf("%s %s",
			e.Top.String(),
			e.Right.String(),
		)
	}

	if e.Right == e.Left {
		return fmt.Sprintf("%s %s %s",
			e.Top.String(),
			e.Right.String(),
			e.Bottom.String(),
		)
	}

	return fmt.Sprintf("%s %s %s %s",
		e.Top.String(),
		e.Right.String(),
		e.Bottom.String(),
		e.Left.String(),
	)
}
