//go:build js && wasm
// +build js,wasm

package widgets

import (
	"github.com/vrianta/tofus/ui/style"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
)

type Button struct {
	Id      string
	Text    string
	Style   style.Context
	Child   app.Widget
	OnClick func(dom.Element)
	OnHover func(dom.Element)
}

func (b Button) Render() dom.Element {
	button := dom.CreateElement("button")

	if b.Id != "" {
		button.SetId(b.Id)
	}

	button.SetText(b.Text)
	button.ApplyStyle(b.Style)

	if b.Child != nil {
		button.AppendChild(b.Child.Render())
	}

	if b.OnClick != nil {
		button.OnClick(b.OnClick)
	}

	if b.OnHover != nil {
		button.OnHover(b.OnHover)
	}

	return button
}

// func IconButton(
// 	icon app.Widget,
// 	onClick func(dom.Element),
// ) app.Widget {
// 	return widgets.Button{
// 		Style: style.Context{
// 			Width:           style.Sizes.Px(40),
// 			Height:          style.Sizes.Px(40),
// 			Padding:         style.EdgeInset{}.SetAll(style.Sizes.Px(8)),
// 			Border:          style.Border{}.None(),
// 			BorderRadius:    style.BorderRadius{}.SetAll(style.Sizes.Percent(50)),
// 			BackgroundColor: style.Colors.Transparent,
// 		},
// 		Child:   icon,
// 		OnClick: onClick,
// 	}
// }
