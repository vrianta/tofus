package style

type BoxSize string

type boxSizes struct{}

var BoxSizes boxSizes

const (
	boxSizingContent BoxSize = "content-box"
	boxSizingBorder  BoxSize = "border-box"
)

func (boxSizes) Content() BoxSize {
	return boxSizingContent
}

func (boxSizes) Border() BoxSize {
	return boxSizingBorder
}
