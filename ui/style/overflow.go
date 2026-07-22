package style

type Overflow string

var Overflows = struct {
	Visible Overflow
	Hidden  Overflow
	Scroll  Overflow
	Auto    Overflow
}{
	Visible: "visible",
	Hidden:  "hidden",
	Scroll:  "scroll",
	Auto:    "auto",
}
