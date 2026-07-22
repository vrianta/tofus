//go:build js && wasm
// +build js,wasm

package widgets

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Button struct {
	Id      string
	Text    string
	Style   style.Context
	OnClick func()
}

func (b Button) Render() dom.Element {
	button := dom.CreateElement("button")

	if b.Id != "" {
		button.SetId(b.Id)
	}

	button.SetText(b.Text)
	button.SetStyle(b.Style.String())

	if b.OnClick != nil {
		button.OnClick(b.OnClick)
	}

	return button
}
