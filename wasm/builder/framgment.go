//go:build js && wasm
// +build js,wasm

package builder

import "github.com/vrianta/tofus/wasm/app"

func Fragment(widgets ...app.Widget) []app.Widget {
	out := make([]app.Widget, 0, len(widgets))

	for _, w := range widgets {
		if w != nil {
			out = append(out, w)
		}
	}

	return out
}
