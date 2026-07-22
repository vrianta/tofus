package style

import "fmt"

type EdgeInset struct {
	Top    Size
	Right  Size
	Bottom Size
	Left   Size
}

var EdgeInsets = EdgeInset{}

func (e EdgeInset) SetAll(v Size) EdgeInset {
	e.Top = v
	e.Bottom = v
	e.Left = v
	e.Right = v

	return e
}

func (e EdgeInset) SetTop(v Size) EdgeInset {
	e.Top = v
	return e
}

func (e EdgeInset) SetButtom(v Size) EdgeInset {
	e.Bottom = v
	return e
}

func (e EdgeInset) SetLeft(v Size) EdgeInset {
	e.Left = v
	return e
}

func (e EdgeInset) SetRight(v Size) EdgeInset {
	e.Right = v
	return e
}

func (e EdgeInset) String() string {
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

func (e EdgeInset) SetHorizontal(v Size) EdgeInset {
	e.Left = v
	e.Right = v
	return e
}

func (e EdgeInset) SetVertical(v Size) EdgeInset {
	e.Top = v
	e.Bottom = v
	return e
}
