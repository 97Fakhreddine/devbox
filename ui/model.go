package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// ToolItem wraps a tool for Bubble Tea list
//type ToolItem struct {
//	NameVal string
//	IDVal   string
//}
//
//func (t ToolItem) Title() string       { return t.NameVal }
//func (t ToolItem) Description() string { return "" }
//func (t ToolItem) FilterValue() string { return t.IDVal }

// Model is the Bubble Tea model
type Model struct {
	Tools       []list.Item
	List        list.Model
	SelectedIdx int
	Done        bool
}

// NewModel creates the model
func NewModel(tools []list.Item) Model {
	const defaultWidth = 20
	l := list.New(tools, list.NewDefaultDelegate(), defaultWidth, 10)
	l.Title = "Select a tool to install"
	return Model{
		Tools: tools,
		List:  l,
		Done:  false,
	}
}

// ------------------ Bubble Tea Interface ------------------

// Init implements tea.Model
func (m Model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.Done = true
			return m, tea.Quit
		case "enter":
			m.SelectedIdx = m.List.Index()
			m.Done = true
			return m, tea.Quit
		}
	}

	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

// View implements tea.Model
func (m Model) View() string {
	if m.Done {
		if m.SelectedIdx >= 0 && m.SelectedIdx < len(m.Tools) {
			return fmt.Sprintf("You selected: %s\n", m.Tools[m.SelectedIdx].FilterValue())
		}
		return "No selection\n"
	}
	return m.List.View()
}

// Run starts Bubble Tea
func Run(model Model) (int, error) {
	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		return -1, err
	}
	// safe type assertion
	m, ok := finalModel.(Model)
	if !ok {
		return -1, fmt.Errorf("unexpected model type")
	}
	return m.SelectedIdx, nil
}
