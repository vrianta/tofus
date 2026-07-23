package render

import (
	"strconv"

	"github.com/vrianta/tofus/ui/style"
)

func FontWeight(f style.FontWeight) string {
	return strconv.Itoa(int(f))
}
