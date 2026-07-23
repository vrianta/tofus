//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Column struct {
	Id       string
	Style    style.Context
	Children []app.Widget
}

func (c *Column) Render() dom.Element {
	div := dom.CreateElement("div")

	if c.Id != "" {
		div.SetId(c.Id)
	}

	// Force column layout
	c.Style.Display = style.DisplaysType.Flex
	c.Style.FlexDirection = style.FlexDirections.Column

	div.ApplyStyle(c.Style)

	for _, child := range c.Children {
		div.AppendChild(child.Render())
	}

	return div
}
