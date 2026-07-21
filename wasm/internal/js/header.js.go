//go:build js && wasm
// +build js,wasm

package js

import (
	"strings"

	"github.com/vrianta/tofus/wasm/browser"
)

func SetTitle(title string) {
	doc.Set("title", title)
}

func SetLang(lang browser.Language) {
	doc.Get("documentElement").Set("lang", string(lang))
}

func SetDescription(des string) {
	setMeta("description", des)
}

func SetKeywords(k []string) {
	keywords := strings.Join(k, ",")
	setMeta("keywords", keywords)
}

func setMeta(name, content string) {

	meta := doc.Call(
		"querySelector",
		`meta[name="`+name+`"]`,
	)

	if meta.IsNull() || meta.IsUndefined() {
		meta = doc.Call("createElement", "meta")
		meta.Set("name", name)
		head.Call("appendChild", meta)
	}

	meta.Set("content", content)
}

func SetCharset(charset browser.Charset) {

	meta := doc.Call(
		"querySelector",
		`meta[charset]`,
	)

	if meta.IsNull() || meta.IsUndefined() {
		meta = doc.Call("createElement", "meta")
		head.Call("appendChild", meta)
	}

	meta.Set("charset", charset)
}

func SetFavicon(href string) {

	link := doc.Call(
		"querySelector",
		`link[rel="icon"]`,
	)

	if link.IsNull() || link.IsUndefined() {
		link = doc.Call("createElement", "link")
		link.Set("rel", "icon")
		head.Call("appendChild", link)
	}

	link.Set("href", href)
}
