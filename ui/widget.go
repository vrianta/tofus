//go:build js && wasm
// +build js,wasm

package ui

import "github.com/vrianta/tofus/ui/dom"

type Widget interface {
	Render() dom.Element
}
