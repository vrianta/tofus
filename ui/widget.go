package ui

type Widget interface {
	render(ctx Context) Element
}
