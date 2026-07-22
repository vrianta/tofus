//go:build js && wasm
// +build js,wasm

package dom

import "syscall/js"

var doc = js.Global().Get("document")
var head = doc.Get("head")
var header = doc.Get("header")

var footer = doc.Get("footer")

func IsOk() bool {
	if doc.IsUndefined() || doc.IsNull() {
		return false
	}

	return true
}
