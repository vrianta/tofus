//go:build js && wasm
// +build js,wasm

package defaults

import "github.com/vrianta/tofus/ui/style"

func TextField() style.Context {
	return style.Context{
		Display: style.DisplaysType.Block,
		Width:   style.Sizes.Percent(100),
		Height:  style.Sizes.Px(42),

		Padding: style.EdgeInset{
			Left:  style.Sizes.Px(12),
			Right: style.Sizes.Px(12),
		},

		BackgroundColor: style.Colors.White,
		Color:           style.Colors.Hex("#111827"),

		FontSize:   style.Sizes.Px(14),
		FontWeight: style.FontWeights.Normal,

		Border: style.Borders.Solid(
			style.Sizes.Px(1),
			style.Colors.Hex("#D1D5DB"),
		),

		BorderRadius: style.BorderRadius{}.SetAll(style.Sizes.Px(8)),

		Outline: style.Outlines.None(),

		BoxSize: style.BoxSizes.Border(),

		Transition: "border-color 0.2s ease, box-shadow 0.2s ease",
	}
}
