package render

import (
	"fmt"
	"strings"

	"github.com/vrianta/tofus/ui/style"
)

// Render Style Context to HTML attributes

func Style(c style.Context) string {
	var css []string

	add := func(key, value string) {
		if value != "" {
			css = append(css, fmt.Sprintf("%s:%s", key, value))
		}
	}

	// Layout
	add("width", Size(c.Width))
	add("height", Size(c.Height))
	add("min-width", Size(c.MinWidth))
	add("min-height", Size(c.MinHeight))
	add("max-width", Size(c.MaxWidth))
	add("max-height", Size(c.MaxHeight))
	add("box-sizing", RenderBoxSize(c.BoxSize))

	// Spacing
	add("padding", EdgeInset(c.Padding))
	add("margin", EdgeInset(c.Margin))

	// Colors
	add("background-color", Color(c.BackgroundColor))
	add("color", Color(c.Color))

	// Border
	add("border", Border(c.Border))
	add("border-radius", BorderRadius(c.BorderRadius))
	add("outline", Outline(c.Outline))

	// Position
	add("position", string(c.Position))
	add("top", Size(c.Top))
	add("right", Size(c.Right))
	add("bottom", Size(c.Bottom))
	add("left", Size(c.Left))

	// Flex
	add("display", Display(c.Display))
	add("flex-direction", string(c.FlexDirection))
	add("justify-content", string(c.JustifyContent))
	add("align-items", string(c.AlignItems))
	add("gap", Size(c.Gap))

	// Typography
	add("font-size", Size(c.FontSize))
	add("font-weight", FontWeight(c.FontWeight))
	add("font-family", c.FontFamily)
	add("text-align", string(c.TextAlign))

	// Effects
	if c.Opacity > 0 {
		add("opacity", fmt.Sprintf("%g", c.Opacity))
	}
	add("box-shadow", BoxShadow(c.Shadow))

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
		add("on-hover", Color(c.OnHover))
	}
	if c.OnActive != "" {
		add("on-active", Color(c.OnActive))
	}
	if c.OnFocus != "" {
		add("on-focus", Color(c.OnFocus))
	}
	// if c.Disabled != nil {
	// 	add("disabled", c.Disabled)
	// }
	// add("", c.BackgroundColor.string())

	// Custom
	for k, v := range c.Custom {
		add(k, v)
	}

	return strings.Join(css, ";")
}
