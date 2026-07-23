//go:build js && wasm
// +build js,wasm

package builder

import "github.com/vrianta/tofus/wasm/app"

func For[T any](items []T, fn func(T) app.Widget) []app.Widget {
	out := make([]app.Widget, 0, len(items))

	for _, item := range items {
		if w := fn(item); w != nil {
			out = append(out, w)
		}
	}

	return out
}
