//go:build js && wasm
// +build js,wasm

package tofus

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/layout"
	"github.com/vrianta/tofus/wasm/app/widgets"
	"github.com/vrianta/tofus/wasm/builder"
)

func Header(title string, search bool, actions ...app.Widget) app.Widget {
	children := []app.Widget{
		&widgets.Text{
			Value: title,
			Style: style.Context{
				Width:      style.Sizes.Auto,
				Color:      style.Colors.White(),
				FontSize:   style.Sizes.Px(22),
				FontWeight: style.FontWeights.Bold,
			},
		},
		&layout.Spacer{},
		builder.If(search, TextField("Search...")).Widget(),
		// actions...,
	}

	// Right-side actions.
	children = append(children, actions...)

	return &layout.Row{
		// Style: style.Context{
		// 	Width:           style.Sizes.Auto,
		// 	Height:          style.Sizes.Px(64),
		// 	BackgroundColor: style.Colors.Hex("#6750A4"),
		// 	AlignItems:      style.AlignItemsList.Center,
		// 	Gap:             style.Sizes.Px(12),
		// 	Padding: style.EdgeInset{
		// 		Left:  style.Sizes.Px(20),
		// 		Right: style.Sizes.Px(20),
		// 	},
		// },
		Children: children,
	}
}
