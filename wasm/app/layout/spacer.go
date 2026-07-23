//go:build js && wasm
// +build js,wasm

package layout

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Spacer struct {
	Id    string
	Style style.Context
}

func (s Spacer) Render() dom.Element {
	div := dom.CreateElement("div")

	if s.Id != "" {
		div.SetId(s.Id)
	}

	// Force the spacer to grow
	s.Style.FlexGrow = 1

	div.ApplyStyle(s.Style)

	return div
}
