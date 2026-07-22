//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

type Element struct {
	js.Value
}

func CreateElement(tag string) Element {
	document := js.Global().Get("document")
	return Element{
		Value: document.Call("createElement", tag),
	}
}

func (e *Element) AppendChild(child Element) {
	e.Value.Call("appendChild", child.Value)
}

func (e Element) JSValue() js.Value {
	return e.Value
}

func (e *Element) SetId(id string) {
	e.Set("id", id)
}

func (e *Element) SetStyle(style string) {
	e.Get("style").Set("cssText", style)
}
