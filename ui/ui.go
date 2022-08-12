package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // list of tasks
	cursor   int              // which task is indicated by cursor
	selected map[int]struct{} // which tasks are selected
}

func InitialModel(s []string) model {
	return model{
		// TODO list of tasks, pull this from db later
		choices: s,
		cursor:  0,
		// this maps the list of choices that are selected. The key is the index of the choice
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it key press?
	case tea.KeyMsg:
		// what key was pressed?
		switch msg.String() {
		//quit app
		case "ctrl+c", "q":
			return m, tea.Quit
		// move up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		// move down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		// select task
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}

		}
	}
	// return updated model
	return m, nil

}

func (m model) View() string {
	// header
	s := "What tasks do you want to focus on?\n\n"

	// list of tasks
	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor
		}

		// is selected?
		checked := " " //no checkmark
		if _, ok := m.selected[i]; ok {
			checked = "x" // checkmark
		}

		// return the row
		s += fmt.Sprintf("%s %s %s\n", checked, cursor, choice)
	}
	// footer
	s += "\nPress q to quit.\n"

	return s
}
