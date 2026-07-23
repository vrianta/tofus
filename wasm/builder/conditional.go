//go:build js && wasm
// +build js,wasm

package builder

import "github.com/vrianta/tofus/wasm/app"

type conditional struct {
	matched bool
	widget  app.Widget
}

func If(condition bool, widget app.Widget) *conditional {
	if condition {
		return &conditional{
			matched: true,
			widget:  widget,
		}
	}

	return &conditional{}
}

func (c *conditional) ElseIf(condition bool, widget app.Widget) *conditional {
	if c.matched {
		return c
	}

	if condition {
		c.matched = true
		c.widget = widget
	}

	return c
}

func (c *conditional) Else(widget app.Widget) app.Widget {
	if c.matched {
		return c.widget
	}

	return widget
}

func (c *conditional) Widget() app.Widget {
	if c.matched {
		return c.widget
	}

	return nil
}
