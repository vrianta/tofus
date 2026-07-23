//go:build js && wasm
// +build js,wasm

package dom

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/vrianta/tofus/ui/style"
)

type Element struct {
	js.Value
	callbacks []js.Func
}

var styleElement js.Value
var styleClassCounter int

func CreateElement(tag string) Element {
	document := js.Global().Get("document")
	return Element{
		Value: document.Call("createElement", tag),
	}
}

func newStyleClassName() string {
	styleClassCounter++
	return fmt.Sprintf("tofus-style-%d", styleClassCounter)
}

func ensureStyleSheet() js.Value {
	if styleElement.IsUndefined() || styleElement.IsNull() {
		doc := js.Global().Get("document")
		styleElement = doc.Call("createElement", "style")
		styleElement.Set("type", "text/css")
		doc.Get("head").Call("appendChild", styleElement)
	}
	return styleElement
}

func appendStyleRules(css string) {
	styleEl := ensureStyleSheet()
	styleEl.Set("textContent", strings.TrimSpace(styleEl.Get("textContent").String()+"\n"+css))
}

func addPseudoRules(className string, ctx style.Context) {
	if ctx.Hover != nil {
		appendStyleRules(fmt.Sprintf(".%s:hover{%s}", className, ctx.Hover.String()))
	}
	if ctx.Active != nil {
		appendStyleRules(fmt.Sprintf(".%s:active{%s}", className, ctx.Active.String()))
	}
	if ctx.Focus != nil {
		appendStyleRules(fmt.Sprintf(".%s:focus{%s}", className, ctx.Focus.String()))
	}
	if ctx.Disabled != nil {
		appendStyleRules(fmt.Sprintf(".%s:disabled{%s}", className, ctx.Disabled.String()))
	}
}

func (e *Element) AddClass(className string) {
	e.Call("classList").Call("add", className)
}

func (e *Element) ApplyStyle(ctx style.Context) {
	if ctx.Hover == nil && ctx.Active == nil && ctx.Focus == nil && ctx.Disabled == nil {
		e.SetStyle(ctx.String())
		return
	}

	className := newStyleClassName()
	e.AddClass(className)
	e.SetStyle(ctx.String())
	addPseudoRules(className, ctx)
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
