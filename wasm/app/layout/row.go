//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Row struct {
	Id       string
	Style    style.Context
	Gap      style.Size
	Children []app.Widget
}

func (r Row) Render() dom.Element {
	row := dom.CreateElement("div")

	if r.Id != "" {
		row.SetId(r.Id)
	}

	// Force row layout
	r.Style.Display = style.Displays.Flex
	r.Style.FlexDirection = style.FlexDirections.Row

	// Apply gap if provided
	if r.Gap != "" {
		r.Style.Gap = r.Gap
	}

	row.ApplyStyle(r.Style)

	for _, child := range r.Children {
		if child != nil {
			row.AppendChild(child.Render())
		}
	}

	return row
}
