//go:build js && wasm
// +build js,wasm

package app

import (
	"syscall/js"
	"wasm/browser"
)

type Header struct {
	doc js.Value

	Title       string
	Description string
	Keywords    string
	Favicon     string

	Charset  browser.Charset
	Language browser.Language
	// Viewport Viewport
}

func (h *Header) init(doc js.Value) {
	h.doc = doc
}

func (h *Header) setMeta(name, content string) {
	head := h.doc.Get("head")

	meta := h.doc.Call(
		"querySelector",
		`meta[name="`+name+`"]`,
	)

	if meta.IsNull() || meta.IsUndefined() {
		meta = h.doc.Call("createElement", "meta")
		meta.Set("name", name)
		head.Call("appendChild", meta)
	}

	meta.Set("content", content)
}

func (h *Header) setCharset(charset browser.Charset) {
	head := h.doc.Get("head")

	meta := h.doc.Call(
		"querySelector",
		`meta[charset]`,
	)

	if meta.IsNull() || meta.IsUndefined() {
		meta = h.doc.Call("createElement", "meta")
		head.Call("appendChild", meta)
	}

	meta.Set("charset", charset)
}

func (h *Header) setFavicon(href string) {
	head := h.doc.Get("head")

	link := h.doc.Call(
		"querySelector",
		`link[rel="icon"]`,
	)

	if link.IsNull() || link.IsUndefined() {
		link = h.doc.Call("createElement", "link")
		link.Set("rel", "icon")
		head.Call("appendChild", link)
	}

	link.Set("href", href)
}

func (h *Header) Render() {
	if h.doc.IsUndefined() || h.doc.IsNull() {
		return
	}

	// <title>
	if h.Title != "" {
		h.doc.Set("title", h.Title)
	}

	// <html lang="">
	if h.Language != "" {
		h.doc.Get("documentElement").Set("lang", h.Language)
	}

	// meta tags
	if h.Description != "" {
		h.setMeta("description", h.Description)
	}

	if h.Keywords != "" {
		h.setMeta("keywords", h.Keywords)
	}

	// charset
	if h.Charset != "" {
		h.setCharset(h.Charset)
	}

	// favicon
	if h.Favicon != "" {
		h.setFavicon(h.Favicon)
	}
}
