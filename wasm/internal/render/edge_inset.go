package render

import (
	"fmt"

	"github.com/vrianta/tofus/ui/style"
)

func EdgeInset(e style.EdgeInset) string {
	if e.Top == e.Right && e.Top == e.Bottom && e.Top == e.Left {
		return Size(e.Top)
	}

	if e.Top == e.Bottom && e.Right == e.Left {
		return fmt.Sprintf("%s %s",
			Size(e.Top),
			Size(e.Right),
		)
	}

	if e.Right == e.Left {
		return fmt.Sprintf("%s %s %s",
			Size(e.Top),
			Size(e.Right),
			Size(e.Bottom),
		)
	}

	return fmt.Sprintf("%s %s %s %s",
		Size(e.Top),
		Size(e.Right),
		Size(e.Bottom),
		Size(e.Left),
	)
}
