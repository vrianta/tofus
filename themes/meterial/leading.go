//go:build js && wasm
// +build js,wasm

package meterial

import (
	"github.com/vrianta/tofus/ui/icons"
	"github.com/vrianta/tofus/wasm/app"
	"github.com/vrianta/tofus/wasm/app/dom"
)

func MenuAppBar(
	title string,
	onMenu func(dom.Element),
	actions ...app.Widget,
) app.Widget {
	return AppBar(
		title,
		IconButton(icons.Menu, onMenu),
		actions...,
	)
}

func BackAppBar(
	title string,
	onBack func(dom.Element),
	actions ...app.Widget,
) app.Widget {
	return AppBar(
		title,
		IconButton(icons.ArrowBack, onBack),
		actions...,
	)
}

func CloseAppBar(
	title string,
	onClose func(dom.Element),
	actions ...app.Widget,
) app.Widget {
	return AppBar(
		title,
		IconButton(icons.Close, onClose),
		actions...,
	)
}

func HomeAppBar(
	title string,
	onHome func(dom.Element),
	actions ...app.Widget,
) app.Widget {
	return AppBar(
		title,
		IconButton(icons.Home, onHome),
		actions...,
	)
}

func DrawerAppBar(
	title string,
	onDrawer func(dom.Element),
	actions ...app.Widget,
) app.Widget {
	return AppBar(
		title,
		IconButton(icons.Menu, onDrawer),
		actions...,
	)
}

func LogoAppBar(
	title string,
	logo app.Widget,
	actions ...app.Widget,
) app.Widget {
	return AppBar(
		title,
		logo,
		actions...,
	)
}

// func AvatarAppBar(
// 	title string,
// 	image string,
// 	onProfile func(dom.Element),
// 	actions ...app.Widget,
// ) app.Widget {
// 	return AppBar(
// 		title,
// 		Avatar(image, onProfile),
// 		actions...,
// 	)
// }
