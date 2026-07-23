package render

import (
	"fmt"

	"github.com/vrianta/tofus/ui/style"
)

func Outline(o style.Outline) string {
	if o.Style == "" {
		return ""
	}

	if o.Style == "none" {
		return "none"
	}

	return fmt.Sprintf("%s %s %s",
		Size(o.Width),
		o.Style,
		Color(o.Color),
	)
}
