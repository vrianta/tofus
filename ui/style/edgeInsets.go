package style

import "fmt"

type EdgeInsets struct {
	Top    Size
	Right  Size
	Bottom Size
	Left   Size
}

func (e EdgeInsets) SetAll(v Size) EdgeInsets {
	e.Top = v
	e.Bottom = v
	e.Left = v
	e.Right = v

	return e
}

func (e EdgeInsets) String() string {
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
