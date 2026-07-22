//go:build js && wasm
// +build js,wasm

package app

import "syscall/js"

func Navigate(path string) {
	js.Global().Get("window").Get("location").Set("href", path)
}

func Replace(path string) {
	js.Global().Get("window").Get("location").Call("replace", path)
}

func Back() {
	js.Global().Get("history").Call("back")
}

func Forward() {
	js.Global().Get("history").Call("forward")
}

func Reload() {
	js.Global().Get("location").Call("reload")
}

func Current() string {
	return js.Global().Get("location").Get("pathname").String()
}
