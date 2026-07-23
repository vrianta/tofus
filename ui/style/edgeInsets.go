package style

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
