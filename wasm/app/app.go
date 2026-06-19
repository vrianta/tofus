//go:build js && wasm
// +build js,wasm

package app

import "syscall/js"

type Context struct {
	document js.Value
	// root     js.Value

	Header *Header
	// Body   *Body
	// Footer *Footer
}

func (a *Context) Run() {
	a.document = js.Global().Get("document")

	if a.Header != nil {
		a.Header.init(js.Global().Get("document"))
		a.Header.Render()
	}

	// if a.Body != nil {
	// 	a.Body.Render()
	// }

	// if a.Footer != nil {
	// 	a.Footer.Render()
	// }
}
