//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
	"github.com/vrianta/tofus/wasm/app/layout"
	"github.com/vrianta/tofus/wasm/app/widgets"
)

var header = layout.Header{
	Children: []app.Widget{
		&layout.Row{
			Style: style.Context{
				Width:           style.Sizes.Percent(100),
				Height:          style.Sizes.Px(64),
				BackgroundColor: style.Colors.Hex("#e2e2e2"),
			},
			Children: []app.Widget{
				&widgets.Text{
					Value: "Heading",
				},
				&layout.Spacer{},
				loginButton,
			},
		},
	},
}

var loginButton = &widgets.Button{
	Text: "Login",
	OnClick: func(e dom.Element) {
		app.Navigate("/login/")
	},
	Style: style.Context{
		Display:        style.DisplaysType.InlineFlex,
		AlignItems:     style.AlignItemsList.Center,
		JustifyContent: style.JustifyContents.Center,

		BackgroundColor: style.Colors.Hex("#2563EB"),
		Color:           style.Colors.White(),

		Padding: style.EdgeInsets.
			SetVertical(style.Sizes.Px(10)).
			SetHorizontal(style.Sizes.Px(18)),

		Border:       style.Borders.None(),
		BorderRadius: style.BorderRadius{}.SetAll(style.Sizes.Px(8)),

		FontSize:   style.Sizes.Px(14),
		FontWeight: style.FontWeights.Medium,

		Cursor:     style.Cursors.Pointer,
		UserSelect: style.UserSelects.None,

		Transition: "all .2s ease",
	},
}
