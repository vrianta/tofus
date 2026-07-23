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

func Login() app.Widget {
	return &layout.Center{
		Style: style.Context{
			Width:  style.Sizes.Percent(100),
			Height: style.Sizes.Vh(100),

			BackgroundColor: style.Colors.Hex("#111827"),
		},
		Child: &layout.Container{
			Style: style.Context{
				Width:           style.Sizes.Px(420),
				BackgroundColor: style.Colors.Hex("#1F2937"),
				BorderRadius:    style.BorderRadius{}.SetAll(style.Sizes.Px(8)),
				Padding:         style.EdgeInsets.SetAll(style.Sizes.Px(32)),
			},
			Child: &layout.Column{
				Children: []app.Widget{

					&widgets.Text{
						Value: "Welcome Back",
						Style: style.Context{
							FontSize:   style.Sizes.Px(32),
							FontWeight: style.FontWeights.Bold,
							Color:      style.Colors.White,
						},
					},

					&widgets.Text{
						Value: "Sign in to continue",
						Style: style.Context{
							Color:  style.Colors.Hex("#9CA3AF"),
							Margin: style.EdgeInsets.SetTop(style.Sizes.Px(8)),
						},
					},

					&layout.Spacer{
						Style: style.Context{
							Height: style.Sizes.Px(24),
						},
					},

					&widgets.TextField{
						Placeholder: "Email",
					},

					&layout.Spacer{
						Style: style.Context{
							Height: style.Sizes.Px(16),
						},
					},

					&widgets.PasswordField{
						Placeholder: "Password",
					},

					&layout.Spacer{
						Style: style.Context{
							Height: style.Sizes.Px(24),
						},
					},

					&widgets.Button{
						Text: "Login",
						OnClick: func(e dom.Element) {
							app.Log("Login clicked")
						},
					},
				},
			},
		},
	}
}
