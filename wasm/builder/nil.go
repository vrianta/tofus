//go:build js && wasm
// +build js,wasm

package builder

import "github.com/vrianta/tofus/wasm/app"

func Show(condition bool, widget app.Widget) app.Widget {
	if condition {
		return widget
	}

	return nil
}

func Hide(condition bool, widget app.Widget) app.Widget {
	if condition {
		return nil
	}

	return widget
}
