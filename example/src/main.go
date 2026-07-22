//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/vrianta/tofus/ui"
	"github.com/vrianta/tofus/ui/layout"
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/ui/widgets"
	"github.com/vrianta/tofus/wasm/app"
)

var App = app.Context{
	Header: &app.Header{
		Title:       "Test Application",
		Description: "A Test Application to demostrate, how to use the wasm ui builder",
	},
	Body: &app.Body{
		Id:     "body",
		Header: &header,
		Children: []ui.Widget{
			&layout.Column{
				Id: "main-column",
				Style: style.Context{
					Width:  style.Sizes.Percent(100),
					Height: style.Sizes.Percent(100),
					Gap:    style.Sizes.Px(12),
				},
				Children: []ui.Widget{
					&widgets.Text{
						Value: "Tofus UI",
					},
					&widgets.Text{
						Value: "Flutter inspired UI library for Go.",
					},
				},
			},
		},
		Style: style.Context{
			Width:           style.Sizes.Percent(100),
			Height:          style.Sizes.Vh(100),
			BackgroundColor: style.Colors.Hex("#fcfcfc"),
			Display:         style.DisplaysType.Block,
			Margin:          style.EdgeInsets.SetAll(style.Sizes.Percent(0)),
			Padding:         style.EdgeInsets.SetAll(style.Sizes.Percent(0)),
		},
	},
}

func main() {
	App.Run()
}
