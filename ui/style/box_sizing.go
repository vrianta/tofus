package style

type BoxSizing string

type boxSizings struct {
	ContentBox BoxSizing
	BorderBox  BoxSizing
}

var BoxSizings = boxSizings{
	ContentBox: "content-box",
	BorderBox:  "border-box",
}
