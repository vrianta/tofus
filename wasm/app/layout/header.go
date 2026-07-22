//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Header struct {
	Id       string
	Style    style.Context
	Children []app.Widget
}

func (h *Header) Render() dom.Element {
	header := dom.CreateElement("header")

	if h.Id != "" {
		header.SetId(h.Id)
	}

	header.SetStyle(h.Style.String())

	for _, child := range h.Children {
		header.AppendChild(child.Render())
	}

	return header
}
