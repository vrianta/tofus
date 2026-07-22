//go:build js && wasm
// +build js,wasm

package app

import "syscall/js"

type Context struct {
	document js.Value
	// root     js.Value

	Head *Head
	Body *Body
	// Footer *Footer
}

func (a *Context) Run() {

	if a.Head != nil {
		a.Head.render()
	}

	if a.Body != nil {
		a.Body.render()
	}

	// if a.Footer != nil {
	// 	a.Footer.Render()
	// }

	select {}
}
