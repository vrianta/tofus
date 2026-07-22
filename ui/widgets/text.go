//go:build js && wasm
// +build js,wasm

package widgets

import (
	"github.com/vrianta/tofus/ui"
	"github.com/vrianta/tofus/ui/dom"
	"github.com/vrianta/tofus/ui/style"
)

type Text struct {
	Id    string
	Value string
	Style style.Context
}

var _ ui.Widget = (*Text)(nil)

func (t *Text) Render() dom.Element {
	element := dom.CreateElement("span")

	if t.Id != "" {
		element.SetId(t.Id)
	}

	element.SetText(t.Value)
	element.SetStyle(t.Style.String())

	return element
}
