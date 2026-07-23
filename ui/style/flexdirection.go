package style

type FlexDirection string

var FlexDirections = struct {
	Row           FlexDirection
	Column        FlexDirection
	RowReverse    FlexDirection
	ColumnReverse FlexDirection
}{
	Row:           "row",
	Column:        "column",
	RowReverse:    "row-reverse",
	ColumnReverse: "column-reverse",
}

type JustifyContent string

var JustifyContents = struct {
	Start        JustifyContent
	Center       JustifyContent
	End          JustifyContent
	SpaceBetween JustifyContent
	SpaceAround  JustifyContent
	SpaceEvenly  JustifyContent
}{
	Start:        "flex-start",
	Center:       "center",
	End:          "flex-end",
	SpaceBetween: "space-between",
	SpaceAround:  "space-around",
	SpaceEvenly:  "space-evenly",
}

type AlignItems string

var AlignItemsList = struct {
	Start    AlignItems
	Center   AlignItems
	End      AlignItems
	Stretch  AlignItems
	Baseline AlignItems
}{
	Start:    "flex-start",
	Center:   "center",
	End:      "flex-end",
	Stretch:  "stretch",
	Baseline: "baseline",
}

type FontWeight int

var FontWeights = struct {
	Thin       FontWeight
	ExtraLight FontWeight
	Light      FontWeight
	Normal     FontWeight
	Medium     FontWeight
	SemiBold   FontWeight
	Bold       FontWeight
	ExtraBold  FontWeight
	Black      FontWeight
}{
	Thin:       100,
	ExtraLight: 200,
	Light:      300,
	Normal:     400,
	Medium:     500,
	SemiBold:   600,
	Bold:       700,
	ExtraBold:  800,
	Black:      900,
}

type Cursor string

var Cursors = struct {
	Auto       Cursor
	Default    Cursor
	Pointer    Cursor
	Text       Cursor
	Move       Cursor
	Wait       Cursor
	NotAllowed Cursor
}{
	Auto:       "auto",
	Default:    "default",
	Pointer:    "pointer",
	Text:       "text",
	Move:       "move",
	Wait:       "wait",
	NotAllowed: "not-allowed",
}

type Visibility string

var Visibilities = struct {
	Visible Visibility
	Hidden  Visibility
}{
	Visible: "visible",
	Hidden:  "hidden",
}

type TextWrap string

var TextWraps = struct {
	Wrap   TextWrap
	NoWrap TextWrap
}{
	Wrap:   "wrap",
	NoWrap: "nowrap",
}

type ObjectFit string

var ObjectFits = struct {
	Fill      ObjectFit
	Contain   ObjectFit
	Cover     ObjectFit
	None      ObjectFit
	ScaleDown ObjectFit
}{
	Fill:      "fill",
	Contain:   "contain",
	Cover:     "cover",
	None:      "none",
	ScaleDown: "scale-down",
}
