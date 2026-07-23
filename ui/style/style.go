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
	Opacity   float64
	BoxShadow BoxShadow

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
	add("width", c.Width.String())
	add("height", c.Height.String())
	add("min-width", c.MinWidth.String())
	add("min-height", c.MinHeight.String())
	add("max-width", c.MaxWidth.String())
	add("max-height", c.MaxHeight.String())
	add("box-sizing", c.BoxSize.string())

	// Spacing
	add("padding", c.Padding.String())
	add("margin", c.Margin.String())

	// Colors
	add("background-color", c.BackgroundColor.String())
	add("color", c.Color.String())

	// Border
	add("border", c.Border.string())
	add("border-radius", c.BorderRadius.String())
	add("outline", c.Outline.String())

	// Position
	add("position", string(c.Position))
	add("top", c.Top.String())
	add("right", c.Right.String())
	add("bottom", c.Bottom.String())
	add("left", c.Left.String())

	// Flex
	add("display", c.Display.String())
	add("flex-direction", string(c.FlexDirection))
	add("justify-content", string(c.JustifyContent))
	add("align-items", string(c.AlignItems))
	add("gap", c.Gap.String())

	// Typography
	add("font-size", c.FontSize.String())
	add("font-weight", c.FontWeight.String())
	add("font-family", c.FontFamily)
	add("text-align", string(c.TextAlign))

	// Effects
	if c.Opacity > 0 {
		add("opacity", fmt.Sprintf("%g", c.Opacity))
	}
	add("box-shadow", c.BoxShadow.String())

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
		add("on-hover", c.OnHover.String())
	}
	if c.OnActive != "" {
		add("on-active", c.OnActive.String())
	}
	if c.OnFocus != "" {
		add("on-focus", c.OnFocus.String())
	}
	if c.Disabled != nil {
		add("disabled", c.Disabled.String())
	}
	// add("", c.BackgroundColor.String())

	// Custom
	for k, v := range c.Custom {
		add(k, v)
	}

	return strings.Join(css, ";")
}
