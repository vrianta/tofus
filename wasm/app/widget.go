//go:build js && wasm
// +build js,wasm

package app

import "github.com/vrianta/tofus/wasm/app/dom"

type Widget interface {
	Render() dom.Element
}
