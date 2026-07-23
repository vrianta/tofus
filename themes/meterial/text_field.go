//go:build js && wasm
// +build js,wasm

package meterial

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/layout"
	"github.com/vrianta/tofus/wasm/app/widgets"
)

func TextField(label string) app.Widget {
	return &layout.Container{
		Style: style.Context{
			Width:    style.Sizes.Percent(100),
			MaxWidth: style.Sizes.Px(420),
		},
		Child: &layout.Column{
			Style: style.Context{
				Gap: style.Sizes.Px(6),
			},
			Children: []app.Widget{
				&widgets.Text{
					Value: label,
					Style: style.Context{
						Color:      style.Colors.Hex("#6750A4"),
						FontSize:   style.Sizes.Px(12),
						FontWeight: style.FontWeights.Medium,
					},
				},
				&widgets.TextField{
					Placeholder: "",
					Style: style.Context{
						Width:  style.Sizes.Percent(100),
						Height: style.Sizes.Px(56),

						BackgroundColor: style.Colors.White,

						Border: style.Border{
							Width: style.Sizes.Px(1),
							Color: style.Colors.Hex("#79747E"),
						},

						BorderRadius: style.BorderRadius{}.SetAll(style.Sizes.Px(4)),

						Padding: style.EdgeInset{
							Left:  style.Sizes.Px(16),
							Right: style.Sizes.Px(16),
						},

						FontSize: style.Sizes.Px(16),
						Color:    style.Colors.Hex("#1C1B1F"),
					},
				},
			},
		},
	}
}
