//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui"
	"github.com/vrianta/tofus/ui/dom"
	"github.com/vrianta/tofus/ui/style"
)

type Header struct {
	Id       string
	Style    style.Context
	Children []ui.Widget
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
