//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui"
	"github.com/vrianta/tofus/ui/dom"
	"github.com/vrianta/tofus/ui/style"
)

type Column struct {
	Id       string
	Style    style.Context
	Children []ui.Widget
}

func (c *Column) Render() dom.Element {
	div := dom.CreateElement("div")

	if c.Id != "" {
		div.SetId(c.Id)
	}

	// Force column layout
	c.Style.Display = style.DisplaysType.Flex
	c.Style.FlexDirection = style.FlexDirections.Column

	div.SetStyle(c.Style.String())

	for _, child := range c.Children {
		div.AppendChild(child.Render())
	}

	return div
}
