//go:build js && wasm
// +build js,wasm

package app

import "syscall/js"

func Log(args ...any) {
	console := js.Global().Get("console")
	console.Call("log", args...)
}
