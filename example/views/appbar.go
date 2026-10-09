//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/vrianta/tofus/themes/meterial"
	"github.com/vrianta/tofus/wasm/app/dom"
)

var AppBar = meterial.AppBar(
	"AppBar Example",
	meterial.HomeAppBar("Menu", func(e dom.Element) {
		println("Menu Clicked")
	}),
)
