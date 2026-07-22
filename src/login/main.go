//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
)

var App = app.Context{
	Head: &app.Head{
		Title:       "Login",
		Description: "Login to the application",
	},
	Body: &app.Body{
		Id: "body",
		Children: []app.Widget{
			Login(),
		},
		Style: style.Context{
			Width:           style.Sizes.Percent(100),
			Height:          style.Sizes.Vh(100),
			BackgroundColor: style.Colors.Hex("#7b7b7b"),
			Display:         style.DisplaysType.Block,
			Margin:          style.EdgeInsets.SetAll(style.Sizes.Percent(0)),
			Padding:         style.EdgeInsets.SetAll(style.Sizes.Percent(0)),
		},
	},
}

func main() {
	App.Run()
}
