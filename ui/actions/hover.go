package actions

import "github.com/vrianta/tofus/ui/style"

type Hover struct {
	Style   style.Context
	OnEnter func()
	OnLeave func()
}
