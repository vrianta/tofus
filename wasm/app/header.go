//go:build js && wasm
// +build js,wasm

package app

import (
	"wasm/browser"
	"wasm/internal/js"
)

type Header struct {
	Title       string
	Description string
	Keywords    []string
	Favicon     string

	Charset  browser.Charset
	Language browser.Language
	// Viewport Viewport
}

func (h *Header) render() {
	if !js.IsOk() {
		return
	}

	// <title>
	if h.Title != "" {
		js.SetTitle(h.Title)
	}

	// <html lang="">
	if h.Language != "" {
		js.SetLang(h.Language)
	}

	// meta tags
	if h.Description != "" {
		js.SetDescription(h.Description)
	}

	if len(h.Keywords) > 0 {
		js.SetKeywords(h.Keywords)
	}

	// charset
	if h.Charset != "" {
		js.SetCharset(h.Charset)
	}

	// favicon
	if h.Favicon != "" {
		js.SetFavicon(h.Favicon)
	}
}
