//go:build js && wasm
// +build js,wasm

package main

import (
	"wasm/ui"
)

var UI = ui.New()

func main() {}

// func main() {
// 	fmt.Println("first program")

// 	doc := js.Global().Get("document")

// 	div := doc.Call("createElement", "div")
// 	div.Set("innerText", "Hello from Go WASM")

// 	doc.Get("body").Call("appendChild", div)

// 	canvas := doc.Call("getElementById", "app")

// 	ctx := canvas.Call("getContext", "2d")

// 	// background
// 	ctx.Set("fillStyle", "#202020")
// 	ctx.Call("fillRect", 0, 0, 800, 600)

// 	// rectangle
// 	ctx.Set("fillStyle", "#00ff00")
// 	ctx.Call("fillRect", 50, 50, 200, 100)

// 	// text
// 	ctx.Set("fillStyle", "#ffffff")
// 	ctx.Set("font", "24px sans-serif")
// 	ctx.Call("fillText", "Hello from Go WASM", 50, 200)

// 	select {}

// }
