package ui

type Widget interface {
	Build(ctx Context) Element
}
