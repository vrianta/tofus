//go:build js && wasm
// +build js,wasm

package app

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Body struct {
	Style    style.Context
	Id       string
	Header   Widget
	Children []Widget
	Footer   Widget
}

func (b *Body) render() {

	body := dom.GetBody()

	if b.Id != "" {
		body.SetId(b.Id)
	}

	body.SetStyle(b.Style.String())
	// render header
	if b.Header != nil {
		body.AppendChild(b.Header.Render())
	}

	for _, child := range b.Children {
		body.AppendChild(child.Render())
	}

	if b.Footer != nil {
		body.AppendChild(b.Footer.Render())
	}
}
