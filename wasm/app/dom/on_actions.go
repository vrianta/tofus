//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

// OnHover registers a callback invoked when the pointer enters or leaves the
// element. The callback receives the element instance.
func (e *Element) OnHover(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("mouseenter", cb)
	e.addEventListener("mouseleave", cb)
}

// OnMouseHover registers a callback invoked on pointer move events while the
// pointer is over the element. The callback receives the element instance.
func (e *Element) OnMouseHover(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("mousemove", cb)
}

// OnMouseMove registers a callback invoked for mouse move events on the element.
// The callback receives the element instance.
func (e *Element) OnMouseMove(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("mousemove", cb)
}

// OnPointerMove registers a callback invoked for pointer move events on the element.
// The callback receives the element instance.
func (e *Element) OnPointerMove(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("pointermove", cb)
}

// OnMouseEnter registers a callback invoked when the mouse enters the element.
// The callback receives the element instance.
func (e *Element) OnMouseEnter(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("mouseenter", cb)
}

// OnMouseLeave registers a callback invoked when the mouse leaves the element.
// The callback receives the element instance.
func (e *Element) OnMouseLeave(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("mouseleave", cb)
}

// OnPointerEnter registers a callback invoked when a pointer enters the element.
// The callback receives the element instance.
func (e *Element) OnPointerEnter(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("pointerenter", cb)
}

// OnPointerLeave registers a callback invoked when a pointer leaves the element.
// The callback receives the element instance.
func (e *Element) OnPointerLeave(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("pointerleave", cb)
}

// OnMouseDown registers a callback invoked when a mouse button is pressed on the element.
func (e *Element) OnMouseDown(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("mousedown", cb)
}

// OnMouseUp registers a callback invoked when a mouse button is released on the element.
func (e *Element) OnMouseUp(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("mouseup", cb)
}

// OnPointerDown registers a callback invoked when a pointer button is pressed on the element.
func (e *Element) OnPointerDown(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("pointerdown", cb)
}

// OnPointerUp registers a callback invoked when a pointer button is released on the element.
func (e *Element) OnPointerUp(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("pointerup", cb)
}

// OnWheel registers a callback invoked when the element receives a wheel event.
// The callback receives the element instance and wheel deltas.
func (e *Element) OnWheel(fn func(_e Element, deltaX, deltaY float64)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		event := args[0]
		fn(*e, event.Get("deltaX").Float(), event.Get("deltaY").Float())
		return nil
	})

	e.addEventListener("wheel", cb)
}

// OnContextMenu registers a callback invoked when the element receives a context menu event.
// The callback receives the element instance.
func (e *Element) OnContextMenu(fn func(Element)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e)
		return nil
	})

	e.addEventListener("contextmenu", cb)
}

// OnInput registers a callback that receives the current input value whenever
// an input event occurs on the element.
func (e *Element) OnInput(fn func(Element, string)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e, this.Get("value").String())
		return nil
	})

	e.addEventListener("input", cb)
}

// OnChange registers a callback that receives the current input value whenever
// a change event occurs on the element.
func (e *Element) OnChange(fn func(Element, string)) {
	cb := js.FuncOf(func(this js.Value, args []js.Value) any {
		fn(*e, this.Get("value").String())
		return nil
	})

	e.addEventListener("change", cb)
}
