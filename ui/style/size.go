package style

import (
	"fmt"
)

type Size string

type sizes struct {
	Auto Size
	Full Size
	None Size
}

var Sizes = sizes{
	Auto: "auto",
	Full: "100%",
	None: "",
}

func (s *sizes) Px(v int) Size {
	return Size(fmt.Sprintf("%dpx", v))
}

func (s *sizes) Percent(v int) Size {
	return Size(fmt.Sprintf("%d%%", v))
}

func (s *sizes) Rem(v float64) Size {
	return Size(fmt.Sprintf("%.2frem", v))
}

func (s *sizes) Em(v float64) Size {
	return Size(fmt.Sprintf("%.2fem", v))
}

func (s *sizes) Vw(v float64) Size {
	return Size(fmt.Sprintf("%.2fvw", v))
}

func (s *sizes) Vh(v float64) Size {
	return Size(fmt.Sprintf("%.2fvh", v))
}
