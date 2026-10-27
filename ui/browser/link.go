package browser

type Link struct {
	Rel  string
	Href string
}

const (
	stylesheet = "stylesheet"
	preconnect = "preconnect"
	preload    = "preload"
	icon       = "icon"
)

func (Link) ToStylesheet() string {
	return stylesheet
}

func (Link) ToPreconnect() string {
	return preconnect
}

func (Link) ToPreload() string {
	return preload
}

func (Link) ToIcon() string {
	return icon
}

func SetLink() Language {
	return Ldef
}
