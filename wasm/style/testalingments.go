package style

type TextAlign string

var TextAlignments = struct {
	Left    TextAlign
	Center  TextAlign
	Right   TextAlign
	Justify TextAlign
}{
	Left:    "left",
	Center:  "center",
	Right:   "right",
	Justify: "justify",
}
