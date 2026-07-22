package style

type userSelect string

type userSelects struct {
	Auto userSelect
	None userSelect
	Text userSelect
	All  userSelect
}

var UserSelects = userSelects{
	Auto: "auto",
	None: "none",
	Text: "text",
	All:  "all",
}
