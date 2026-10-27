//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/vrianta/tofus/themes/tofus"
	"github.com/vrianta/tofus/ui/browser"
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/layout"
	"github.com/vrianta/tofus/wasm/app/widgets"
)

var App = app.Context{
	Head: tofus.Head(
		"Example App",
		"An example app built with Tofus UI.",
		[]string{"tofus", "ui", "example"},
		"/favicon.ico",
		browser.UTF8,
		browser.English,
		[]browser.Link{
			{
				Rel:  "stylesheet",
				Href: "https://fonts.googleapis.com/css2?family=Roboto:wght@400;500;700&display=swap",
			},
		},
	),
	Body: &app.Body{
		Id:     "body",
		Header: tofus.Header("Eample App", true),
		Children: []app.Widget{
			AppBar,
			&layout.Column{
				Id: "main-column",
				Style: style.Context{
					Width:  style.Sizes.Percent(100),
					Height: style.Sizes.Percent(100),
					Gap:    style.Sizes.Px(12),
				},
				Children: []app.Widget{
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
			BackgroundColor: style.Colors.Hex("#7b7b7b"),
			Display:         style.Displays.Block,
			Margin:          style.EdgeInsets.SetAll(style.Sizes.Percent(0)),
			Padding:         style.EdgeInsets.SetAll(style.Sizes.Percent(0)),
		},
	},
}

func main() {
	App.Run()
}
