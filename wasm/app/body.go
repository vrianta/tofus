//go:build js && wasm
// +build js,wasm

package app

import (
	"wasm/internal/js"
	"wasm/style"
)

type Body struct {
	Style    style.Context
	Id       string
	Children []Widget
}

func (b *Body) render() {

	body := js.GetBody()

	if b.Id != "" {
		body.SetId(b.Id)
	}

	body.SetStyle(b.Style.String())

	for _, child := range b.Children {
		child.render()
	}
}
