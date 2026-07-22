//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

type Element struct {
	js.Value
	callbacks []js.Func
}

func CreateElement(tag string) Element {
	document := js.Global().Get("document")
	return Element{
		Value: document.Call("createElement", tag),
	}
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

func (e *Element) SetText(text string) {
	e.Set("textContent", text)
}

func (e *Element) SetHTML(html string) {
	e.Set("innerHTML", html)
}

func (e *Element) AppendChild(child Element) {
	e.Value.Call("appendChild", child.Value)
}

func (e *Element) OnClick(fn func()) {
	callback := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn()
		return nil
	})

	e.Call("addEventListener", "click", callback)
}

func (e *Element) SetAttribute(name, value string) {
	e.Call("setAttribute", name, value)
}

func (e *Element) RemoveAttribute(name string) {
	e.Call("removeAttribute", name)
}

func (e *Element) GetAttribute(name string) string {
	v := e.Call("getAttribute", name)
	if v.IsNull() || v.IsUndefined() {
		return ""
	}
	return v.String()
}

func (e *Element) HasAttribute(name string) bool {
	return e.Call("hasAttribute", name).Bool()
}

func (e *Element) OnInput(fn func(string)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(this.Get("value").String())
		return nil
	})

	e.callbacks = append(e.callbacks, cb)
	e.Call("addEventListener", "input", cb)
}

func (e *Element) OnChange(fn func(string)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(this.Get("value").String())
		return nil
	})

	e.callbacks = append(e.callbacks, cb)
	e.Call("addEventListener", "change", cb)
}
