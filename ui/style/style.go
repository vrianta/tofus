package style

type Context struct {
	// Layout
	Width     Size
	Height    Size
	MinWidth  Size
	MinHeight Size
	MaxWidth  Size
	MaxHeight Size
	BoxSize   BoxSize

	// Spacing
	Padding EdgeInset
	Margin  EdgeInset

	// Colors
	BackgroundColor Color
	Color           Color

	// Border
	Border       Border
	BorderRadius BorderRadius
	Outline      Outline

	// Position
	Position Position
	Top      Size
	Right    Size
	Bottom   Size
	Left     Size

	// Flex
	Display        Display
	FlexDirection  FlexDirection
	JustifyContent JustifyContent
	AlignItems     AlignItems
	Gap            Size

	// Typography
	FontSize   Size
	FontWeight FontWeight
	FontFamily string
	TextAlign  TextAlign

	// Effects
	Opacity float64
	Shadow  Shadow

	// Overflow
	Overflow Overflow

	// Visibility
	Visible bool

	// Escape hatch
	Custom map[string]string

	ZIndex     int
	Cursor     Cursor
	Visibility Visibility

	FlexGrow   int
	FlexShrink int

	TextWrap TextWrap

	ObjectFit ObjectFit

	Transition string
	Transform  string

	UserSelect userSelect

	OnHover  Color // onHover Color
	OnActive Color // onActive Color
	OnFocus  Color // onFocus Color
	// Disabled *Context
}
