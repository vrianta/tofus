package render

import (
	"fmt"

	"github.com/vrianta/tofus/ui/style"
)

func BorderRadius(b style.BorderRadius) string {
	// all sides equal
	if b.TopLeft == b.TopRight &&
		b.TopLeft == b.BottomRight &&
		b.TopLeft == b.BottomLeft {
		return Size(b.TopLeft)
	}

	// top-left/bottom-right and top-right/bottom-left
	if b.TopLeft == b.BottomRight &&
		b.TopRight == b.BottomLeft {
		return fmt.Sprintf(
			"%s %s",
			Size(b.TopLeft),
			Size(b.TopRight),
		)
	}

	return fmt.Sprintf(
		"%s %s %s %s",
		Size(b.TopLeft),
		Size(b.TopRight),
		Size(b.BottomRight),
		Size(b.BottomLeft),
	)
}
