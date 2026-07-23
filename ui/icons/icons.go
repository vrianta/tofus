//go:build js && wasm
// +build js,wasm

package icons

import "github.com/vrianta/tofus/wasm/app/dom"

type Icon struct {
	PathData string
	ViewBox  string
}

func (i Icon) Render() dom.Element {
	svg := dom.CreateElementNS("http://www.w3.org/2000/svg", "svg")
	svg.SetAttribute("viewBox", i.ViewBox)
	svg.SetAttribute("fill", "currentColor")
	svg.SetAttribute("width", "100%")
	svg.SetAttribute("height", "100%")
	svg.SetAttribute("style", "display:block; width:100%; height:100%")

	path := dom.CreateElementNS("http://www.w3.org/2000/svg", "path")
	path.SetAttribute("d", i.PathData)
	svg.AppendChild(path)

	return svg
}

var (
	Menu = Icon{
		ViewBox:  "0 0 24 24",
		PathData: "M3 6h18v2H3V6zm0 5h18v2H3v-2zm0 5h18v2H3v-2z",
	}
	ArrowBack = Icon{
		ViewBox:  "0 0 24 24",
		PathData: "M20 11H7.83l5.59-5.59L12 4l-8 8 8 8 1.41-1.41L7.83 13H20v-2z",
	}
	Close = Icon{
		ViewBox:  "0 0 24 24",
		PathData: "M18.3 5.71L12 12.01 5.7 5.71 4.29 7.12 10.59 13.41 4.29 19.71 5.7 21.12 12 14.83 18.3 21.12 19.71 19.71 13.41 13.41 19.71 7.12 18.3 5.71z",
	}
	Home = Icon{
		ViewBox:  "0 0 24 24",
		PathData: "M10 20v-6h4v6h5v-8h3L12 3 2 12h3v8z",
	}
)
