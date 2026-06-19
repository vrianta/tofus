package style

type Position string

var Positions = struct {
	Static   Position
	Relative Position
	Absolute Position
	Fixed    Position
	Sticky   Position
}{
	Static:   "static",
	Relative: "relative",
	Absolute: "absolute",
	Fixed:    "fixed",
	Sticky:   "sticky",
}
