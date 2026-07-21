//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/style"
)

var App = app.Context{
	Header: &app.Header{
		Title:       "Test Application",
		Description: "A Test APlication to demostrate, how to use the wasm ui builder",
	},
	Body: &app.Body{
		Id: "body",

		Style: style.Context{
			Width:           style.Sizes.Percent(100),
			Height:          style.Sizes.Vh(100),
			BackgroundColor: style.Colors.Hex("#150d0d"),
			Display:         style.DisplaysType.Block,
			Padding:         style.EdgeInsets{}.SetAll(style.Sizes.Px(20)),
		},
	},
}

func main() {
	App.Run()
}
