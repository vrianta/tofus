//go:build js && wasm
// +build js,wasm

package meterial

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/layout"
	"github.com/vrianta/tofus/wasm/app/widgets"
	"github.com/vrianta/tofus/wasm/builder"
)

func AppBar(
	title string,
	leading app.Widget,
	actions ...app.Widget,
) app.Widget {

	children := []app.Widget{
		leading,

		builder.If(title != "" && leading == nil, &widgets.Text{
			Value: title,
			Style: style.Context{
				FontSize:   style.Sizes.Px(22),
				FontWeight: style.FontWeights.SemiBold,
			},
		}).Widget(),

		layout.Spacer{},
	}
	children = append(children, actions...)

	return &layout.Container{
		Style: style.Context{
			Width:           style.Sizes.Full,
			Height:          style.Sizes.Px(64),
			Display:         style.Displays.Flex,
			AlignItems:      style.AlignItemsList.Center,
			BackgroundColor: style.Colors.White(),
			Padding: style.EdgeInset{}.
				SetHorizontal(style.Sizes.Px(16)),
			Border: style.Borders.Solid(
				style.Sizes.Px(1),
				style.Colors.RGBA(0, 0, 0, 0.08),
			),
			Shadow: style.Shadow{
				Inset:  false,
				X:      style.Sizes.Px(0),
				Y:      style.Sizes.Px(1),
				Blur:   style.Sizes.Px(3),
				Spread: style.Sizes.Px(0),
				Color:  style.Colors.RGBA(0, 0, 0, 0.08),
			},
		},
		Child: layout.Row{
			Gap:      style.Sizes.Px(12),
			Children: children,
		},
	}
}
