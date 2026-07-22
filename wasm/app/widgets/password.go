//go:build js && wasm
// +build js,wasm

package widgets

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type PasswordField struct {
	Id          string
	Value       string
	Placeholder string
	Name        string

	Style style.Context

	OnInput  func(string)
	OnChange func(string)
}

func (p *PasswordField) Render() dom.Element {
	tf := &TextField{
		Id:          p.Id,
		Value:       p.Value,
		Placeholder: p.Placeholder,
		Name:        p.Name,
		Type:        "password",

		Style: p.Style,

		OnInput:  p.OnInput,
		OnChange: p.OnChange,
	}

	return tf.Render()
}
