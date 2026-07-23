package style

type Display string

type displays struct {
	Block       Display
	Inline      Display
	InlineBlock Display
	Flex        Display
	InlineFlex  Display
	Grid        Display
	None        Display
}

var Displays = displays{
	Block:       "block",
	Inline:      "inline",
	InlineBlock: "inline-block",
	Flex:        "flex",
	InlineFlex:  "inline-flex",
	Grid:        "grid",
	None:        "none",
}

func (d Display) string() string {
	return string(d)
}
