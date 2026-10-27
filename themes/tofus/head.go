//go:build js && wasm
// +build js,wasm

package tofus

import (
	"github.com/vrianta/tofus/ui/browser"
	"github.com/vrianta/tofus/wasm/app"
)

func Head(
	title string,
	description string,
	keywords []string,
	favicon string,
	charset browser.Charset,
	language browser.Language,
	link []browser.Link,
) *app.Head {
	link = append(link, browser.Link{
		Rel:  "stylesheet",
		Href: "/pico-css/pico.min.css",
	})
	return &app.Head{
		Title:       title,
		Description: description,
		Keywords:    keywords,
		Favicon:     favicon,
		Charset:     charset,
		Language:    language,
		Link:        link,
	}
}
