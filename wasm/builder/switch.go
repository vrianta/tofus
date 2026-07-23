//go:build js && wasm
// +build js,wasm

package builder

import "github.com/vrianta/tofus/wasm/app"

type switchBuilder[T comparable] struct {
	value   T
	matched bool
	widget  app.Widget
}

func Switch[T comparable](value T) *switchBuilder[T] {
	return &switchBuilder[T]{
		value: value,
	}
}

func (s *switchBuilder[T]) Case(value T, widget app.Widget) *switchBuilder[T] {
	if s.matched {
		return s
	}

	if s.value == value {
		s.matched = true
		s.widget = widget
	}

	return s
}

func (s *switchBuilder[T]) Default(widget app.Widget) app.Widget {
	if s.matched {
		return s.widget
	}

	return widget
}

func (s *switchBuilder[T]) Widget() app.Widget {
	if s.matched {
		return s.widget
	}

	return nil
}

func (s *switchBuilder[T]) When(fn func(T) bool, widget app.Widget) *switchBuilder[T] {
	if s.matched {
		return s
	}

	if fn(s.value) {
		s.matched = true
		s.widget = widget
	}

	return s
}
