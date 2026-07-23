package render

import (
	"fmt"

	"github.com/vrianta/tofus/ui/style"
)

func BoxShadow(s style.Shadow) string {
	if s.X == "" &&
		s.Y == "" &&
		s.Blur == "" &&
		s.Spread == "" &&
		s.Color == "" {
		return ""
	}

	prefix := ""
	if s.Inset {
		prefix = "inset "
	}

	return fmt.Sprintf(
		"%s%s %s %s %s %s",
		prefix,
		Size(s.X),
		Size(s.Y),
		Size(s.Blur),
		Size(s.Spread),
		Color(s.Color),
	)
}
