//go:build js && wasm
// +build js,wasm

package js

import "syscall/js"

type body struct {
	js.Value
}

func (b *body) SetId(id string) {
	b.Set("id", id)
}

func (b *body) SetStyle(style string) {
	b.Get("style").Set("cssText", style)
}

func GetBody() body {

	return body{
		Value: doc.Get("body"),
	}
}
