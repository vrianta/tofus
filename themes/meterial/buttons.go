//go:build js && wasm
// +build js,wasm

package meterial

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/widgets"
)

func IconButton(
	icon app.Widget,
	onClick func(),
) app.Widget {
	return widgets.Button{
		Style: style.Context{
			Width:           style.Sizes.Px(40),
			Height:          style.Sizes.Px(40),
			Padding:         style.EdgeInset{}.SetAll(style.Sizes.Px(8)),
			Border:          style.Border{}.None(),
			BorderRadius:    style.BorderRadius{}.SetAll(style.Sizes.Percent(50)),
			BackgroundColor: style.Colors.Transparent,
			// HoverColor:      style.Colors.Hex("#E0E0E0"),
			// PressedColor:    style.Colors.Hex("#D5D5D5"),
			// OnClick:         onClick,
		},
		// icon,
	}
}
