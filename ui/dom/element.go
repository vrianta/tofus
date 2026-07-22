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

func (e *Element) AppendHeader(element Element) Element {
	header := CreateElement("header")

	e.AppendChild(element)
	return header
}

func (e *Element) AppendFooter(element Element) Element {
	footer := CreateElement("footer")
	e.AppendChild(element)
	return footer
}

func (e *Element) AppendMain() Element {
	main := CreateElement("main")
	e.AppendChild(main)
	return main
}

func (e *Element) AppendSection() Element {
	section := CreateElement("section")
	e.AppendChild(section)
	return section
}

func (e *Element) AppendDiv() Element {
	div := CreateElement("div")
	e.AppendChild(div)
	return div
}
