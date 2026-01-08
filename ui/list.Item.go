package ui

type ToolItem struct {
	NameVal string
	IDVal   string
}

func (t ToolItem) Title() string       { return t.NameVal }
func (t ToolItem) Description() string { return "" }
func (t ToolItem) FilterValue() string { return t.IDVal }
