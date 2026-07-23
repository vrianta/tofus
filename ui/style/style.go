package style

import (
	"fmt"
	"strings"
)

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
	Disabled *Context
}

func (c Context) String() string {
	var css []string

	add := func(key, value string) {
		if value != "" {
			css = append(css, fmt.Sprintf("%s:%s", key, value))
		}
	}

	// Layout
	add("width", c.Width.string())
	add("height", c.Height.string())
	add("min-width", c.MinWidth.string())
	add("min-height", c.MinHeight.string())
	add("max-width", c.MaxWidth.string())
	add("max-height", c.MaxHeight.string())
	add("box-sizing", c.BoxSize.string())

	// Spacing
	add("padding", c.Padding.string())
	add("margin", c.Margin.string())

	// Colors
	add("background-color", c.BackgroundColor.string())
	add("color", c.Color.string())

	// Border
	add("border", c.Border.string())
	add("border-radius", c.BorderRadius.string())
	add("outline", c.Outline.string())

	// Position
	add("position", string(c.Position))
	add("top", c.Top.string())
	add("right", c.Right.string())
	add("bottom", c.Bottom.string())
	add("left", c.Left.string())

	// Flex
	add("display", c.Display.string())
	add("flex-direction", string(c.FlexDirection))
	add("justify-content", string(c.JustifyContent))
	add("align-items", string(c.AlignItems))
	add("gap", c.Gap.string())

	// Typography
	add("font-size", c.FontSize.string())
	add("font-weight", c.FontWeight.string())
	add("font-family", c.FontFamily)
	add("text-align", string(c.TextAlign))

	// Effects
	if c.Opacity > 0 {
		add("opacity", fmt.Sprintf("%g", c.Opacity))
	}
	add("box-shadow", c.Shadow.string())

	// Overflow
	add("overflow", string(c.Overflow))

	// Visibility TODO
	// if !c.Visible && c.Display != "" {
	// 	add("display", "none")
	// }

	add("cursor", string(c.Cursor))
	add("visibility", string(c.Visibility))
	add("text-wrap", string(c.TextWrap))
	add("object-fit", string(c.ObjectFit))

	if c.ZIndex != 0 {
		add("z-index", fmt.Sprintf("%d", c.ZIndex))
	}

	if c.FlexGrow != 0 {
		add("flex-grow", fmt.Sprintf("%d", c.FlexGrow))
	}

	if c.FlexShrink != 0 {
		add("flex-shrink", fmt.Sprintf("%d", c.FlexShrink))
	}

	add("transition", c.Transition)
	add("transform", c.Transform)

	// onHover, onActive, onFocus, disabled
	if c.OnHover != "" {
		add("on-hover", c.OnHover.string())
	}
	if c.OnActive != "" {
		add("on-active", c.OnActive.string())
	}
	if c.OnFocus != "" {
		add("on-focus", c.OnFocus.string())
	}
	if c.Disabled != nil {
		add("disabled", c.Disabled.String())
	}
	// add("", c.BackgroundColor.string())

	// Custom
	for k, v := range c.Custom {
		add(k, v)
	}

	return strings.Join(css, ";")
}
