//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

type Element struct {
	value js.Value
}

func CreateElement(tag string) Element {
	document := js.Global().Get("document")
	return Element{
		value: document.Call("createElement", tag),
	}
}

func (e Element) JSValue() js.Value {
	return e.value
}
