//go:build js && wasm
// +build js,wasm

package widgets

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type TextField struct {
	Id          string
	Value       string
	Placeholder string
	Name        string
	Type        string

	Style style.Context

	OnChange func(string)
	OnInput  func(string)
}

func (t *TextField) Render() dom.Element {
	input := dom.CreateElement("input")

	if t.Id != "" {
		input.SetId(t.Id)
	}

	input.SetAttribute("type", "text")

	if t.Type != "" {
		input.SetAttribute("type", t.Type)
	}

	if t.Name != "" {
		input.SetAttribute("name", t.Name)
	}

	if t.Placeholder != "" {
		input.SetAttribute("placeholder", t.Placeholder)
	}

	if t.Value != "" {
		input.SetAttribute("value", t.Value)
	}

	input.ApplyStyle(t.Style)

	if t.OnInput != nil {
		input.OnInput(func(e dom.Element, value string) {
			t.Value = value
			t.OnInput(value)
		})
	}

	if t.OnChange != nil {
		input.OnChange(func(e dom.Element, value string) {
			t.Value = value
			t.OnChange(value)
		})
	}

	return input
}
