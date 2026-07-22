package layout

import (
	"github.com/vrianta/tofus/ui"
	"github.com/vrianta/tofus/ui/style"
)

type Column struct {
	Id       string
	Style    style.Context
	Children []ui.Widget
}
