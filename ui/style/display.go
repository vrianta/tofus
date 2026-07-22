package style

type Display string

type displays struct {
	Block       Display
	Inline      Display
	InlineBlock Display
	Flex        Display
	Grid        Display
	None        Display
}

var DisplaysType = displays{
	Block:       "block",
	Inline:      "inline",
	InlineBlock: "inline-block",
	Flex:        "flex",
	Grid:        "grid",
	None:        "none",
}

func (d Display) String() string {
	return string(d)
}
