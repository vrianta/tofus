//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/vrianta/tofus/ui"
	"github.com/vrianta/tofus/ui/layout"
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/ui/widgets"
)

var header = layout.Header{
	Style: style.Context{
		Width:           style.Sizes.Percent(100),
		Height:          style.Sizes.Px(64),
		BackgroundColor: style.Colors.Hex("#202020"),
	},
	Children: []ui.Widget{
		&widgets.Text{
			Value: "Tofus",
		},
	},
}
