//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/vrianta/tofus/themes/tofus"
	"github.com/vrianta/tofus/wasm/app/dom"
)

var AppBar = tofus.AppBar(
	"AppBar Example",
	tofus.HomeAppBar("Menu", func(e dom.Element) {
		println("Menu Clicked")
	}),
)
