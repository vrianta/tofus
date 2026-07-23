//go:build js && wasm
// +build js,wasm

package meterial

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
	"github.com/vrianta/tofus/wasm/app/widgets"
)

func IconButton(
	icon app.Widget,
	onClick func(dom.Element),
) app.Widget {
	return widgets.Button{
		Style: style.Context{
			Width:           style.Sizes.Px(40),
			Height:          style.Sizes.Px(40),
			Padding:         style.EdgeInset{}.SetAll(style.Sizes.Px(8)),
			Border:          style.Borders.None(),
			BorderRadius:    style.BorderRadius{}.SetAll(style.Sizes.Percent(50)),
			BackgroundColor: style.Colors.Transparent,
		},
		Child:   icon,
		OnClick: onClick,
	}
}
