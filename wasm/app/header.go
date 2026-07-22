//go:build js && wasm
// +build js,wasm

package app

import (
	"github.com/vrianta/tofus/ui/browser"
	"github.com/vrianta/tofus/ui/dom"
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
	if !dom.IsOk() {
		return
	}

	// <title>
	if h.Title != "" {
		dom.SetTitle(h.Title)
	}

	// <html lang="">
	if h.Language != "" {
		dom.SetLang(h.Language)
	}

	// meta tags
	if h.Description != "" {
		dom.SetDescription(h.Description)
	}

	if len(h.Keywords) > 0 {
		dom.SetKeywords(h.Keywords)
	}

	// charset
	if h.Charset != "" {
		dom.SetCharset(h.Charset)
	}

	// favicon
	if h.Favicon != "" {
		dom.SetFavicon(h.Favicon)
	}
}
