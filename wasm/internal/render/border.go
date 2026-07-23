package render

import (
	"fmt"

	"github.com/vrianta/tofus/ui/style"
)

func Border(b style.Border) string {
	if b.Width == "" {
		b.Width = "0px"
	}
	if b.Style == "" {
		b.Style = "solid"
	}
	if b.Color == "" {
		b.Color = style.Colors.Black()
	}
	return fmt.Sprintf(
		"%s %s %s",
		Size(b.Width),
		b.Style,
		Color(b.Color),
	)
}
