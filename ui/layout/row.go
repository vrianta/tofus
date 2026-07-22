//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui"
	"github.com/vrianta/tofus/ui/dom"
	"github.com/vrianta/tofus/ui/style"
)

type Row struct {
	Id       string
	Style    style.Context
	Children []ui.Widget
}

func (r Row) Render() dom.Element {
	row := dom.CreateElement("div")

	if r.Id != "" {
		row.SetId(r.Id)
	}

	// Force row layout
	r.Style.Display = style.DisplaysType.Flex
	r.Style.FlexDirection = style.FlexDirections.Row

	row.SetStyle(r.Style.String())

	for _, child := range r.Children {
		row.AppendChild(child.Render())
	}

	return row
}
