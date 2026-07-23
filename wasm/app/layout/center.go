//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Center struct {
	Id    string
	Style style.Context
	Child app.Widget
}

func (c *Center) Render() dom.Element {
	div := dom.CreateElement("div")

	if c.Id != "" {
		div.SetId(c.Id)
	}

	s := c.Style

	if s.Display == "" {
		s.Display = style.DisplaysType.Flex
	}

	if s.Width == style.Sizes.None {
		s.Width = style.Sizes.Percent(100)
	}

	if s.Height == style.Sizes.None {
		s.Height = style.Sizes.Percent(100)
	}

	if s.JustifyContent == "" {
		s.JustifyContent = style.JustifyContents.Center
	}

	if s.AlignItems == "" {
		s.AlignItems = style.AlignItemsList.Center
	}

	div.ApplyStyle(s)

	if c.Child != nil {
		div.AppendChild(c.Child.Render())
	}

	return div
}
