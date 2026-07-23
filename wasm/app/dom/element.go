//go:build js && wasm
// +build js,wasm

package dom

import (
	"fmt"
	"strings"
	"syscall/js"

	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/internal/render"
)

type Element struct {
	js.Value
	callbacks []js.Func
}

var styleElement js.Value
var styleClassCounter int

// CreateElement creates a new DOM element with the given tag name and wraps it
// in the local Element type.
func CreateElement(tag string) Element {
	document := js.Global().Get("document")
	return Element{
		Value: document.Call("createElement", tag),
	}
}

// CreateElementNS creates a new DOM element in the provided namespace.
func CreateElementNS(ns, tag string) Element {
	document := js.Global().Get("document")
	return Element{
		Value: document.Call("createElementNS", ns, tag),
	}
}

// newStyleClassName returns a unique CSS class name for generated style rules.
func newStyleClassName() string {
	styleClassCounter++
	return fmt.Sprintf("tofus-style-%d", styleClassCounter)
}

// ensureStyleSheet lazily creates and returns a shared <style> element in the
// document head used to inject generated CSS rules.
func ensureStyleSheet() js.Value {
	if styleElement.IsUndefined() || styleElement.IsNull() {
		doc := js.Global().Get("document")
		styleElement = doc.Call("createElement", "style")
		styleElement.Set("type", "text/css")
		doc.Get("head").Call("appendChild", styleElement)
	}
	return styleElement
}

// appendStyleRules appends the provided CSS rules text to the shared stylesheet.
func appendStyleRules(css string) {
	styleEl := ensureStyleSheet()
	styleEl.Set("textContent", strings.TrimSpace(styleEl.Get("textContent").String()+"\n"+css))
}

func addPseudoRules(className string, ctx style.Context) {
	if ctx.OnHover != "" {
		appendStyleRules(fmt.Sprintf(".%s:hover{%s}", className, string(ctx.OnHover)))
	}
	if ctx.OnActive != "" {
		appendStyleRules(fmt.Sprintf(".%s:active{%s}", className, string(ctx.OnActive)))
	}
	if ctx.OnFocus != "" {
		appendStyleRules(fmt.Sprintf(".%s:focus{%s}", className, string(ctx.OnFocus)))
	}
	// if ctx.Disabled != nil {
	// 	appendStyleRules(fmt.Sprintf(".%s:disabled{%s}", className, ctx.Disabled.String()))
	// }
}

// AddClass adds a CSS class name to the element's class list.
func (e *Element) AddClass(className string) {
	e.Call("classList").Call("add", className)
}

// ApplyStyle applies a style context to the element. If the context contains
// pseudo-state styles such as Hover, Active, Focus, or Disabled, a generated
// CSS class is created and the pseudo rules are appended to the shared stylesheet.
func (e *Element) ApplyStyle(ctx style.Context) {
	if ctx.OnHover == "" && ctx.OnActive == "" && ctx.OnFocus == "" { //&& ctx.Disabled == nil {
		e.SetStyle(ctx)
		return
	}

	className := newStyleClassName()
	e.AddClass(className)
	e.SetStyle(ctx)
	addPseudoRules(className, ctx)
}

func (e Element) JSValue() js.Value {
	return e.Value
}

// SetId sets the element's id attribute.
func (e *Element) SetId(id string) {
	e.Set("id", id)
}

// SetStyle sets the element's inline CSS text from the provided style context.
func (e *Element) SetStyle(style style.Context) {
	e.Get("style").Set("cssText", render.Style(style))
}

// SetText sets the element's text content.
func (e *Element) SetText(text string) {
	e.Set("textContent", text)
}

// SetHTML sets the element's inner HTML.
func (e *Element) SetHTML(html string) {
	e.Set("innerHTML", html)
}

// AppendChild appends a child element to this element.
func (e *Element) AppendChild(child Element) {
	e.Value.Call("appendChild", child.Value)
}

// OnClick registers a click event callback for the element.
func (e *Element) OnClick(fn func(Element)) {
	callback := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("click", callback)
}

// addEventListener registers a JS event listener on the element and keeps the
// callback alive in the element's callback slice.
func (e *Element) addEventListener(event string, callback js.Func) {
	e.callbacks = append(e.callbacks, callback)
	e.Call("addEventListener", event, callback)
}

// SetAttribute sets an attribute on the element.
func (e *Element) SetAttribute(name, value string) {
	e.Call("setAttribute", name, value)
}

// RemoveAttribute removes an attribute from the element.
func (e *Element) RemoveAttribute(name string) {
	e.Call("removeAttribute", name)
}

// GetAttribute returns the value of the named attribute, or an empty string if
// the attribute is not present.
func (e *Element) GetAttribute(name string) string {
	v := e.Call("getAttribute", name)
	if v.IsNull() || v.IsUndefined() {
		return ""
	}
	return v.String()
}

// HasAttribute reports whether the element has the named attribute.
func (e *Element) HasAttribute(name string) bool {
	return e.Call("hasAttribute", name).Bool()
}
