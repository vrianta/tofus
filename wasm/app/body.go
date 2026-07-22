//go:build js && wasm
// +build js,wasm

package app

import (
	"github.com/vrianta/tofus/ui"
	"github.com/vrianta/tofus/ui/dom"
	"github.com/vrianta/tofus/ui/style"
)

type Body struct {
	Style    style.Context
	Id       string
	Header   ui.Widget
	Children []ui.Widget
	Footer   ui.Widget
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
