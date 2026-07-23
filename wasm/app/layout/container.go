//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Container struct {
	Id    string
	Style style.Context
	Child app.Widget
}

// // Render implements [app.Widget].
// func (c *Container) Render() dom.Element {
// 	panic("unimplemented")
// }

func (c *Container) Render() dom.Element {
	div := dom.CreateElement("div")

	if c.Id != "" {
		div.SetId(c.Id)
	}

	div.ApplyStyle(c.Style)

	if c.Child != nil {
		div.AppendChild(c.Child.Render())
	}

	return div
}
