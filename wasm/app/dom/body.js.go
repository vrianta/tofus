//go:build js && wasm
// +build js,wasm

package dom

func GetBody() Element {

	return Element{
		Value: doc.Get("body"),
	}
}
