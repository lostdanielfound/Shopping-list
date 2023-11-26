package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices  []string         // itesm on the list
	cursor   int              // which item our curor is pointing at
	selected map[int]struct{} // which items are selected
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "crtl+c", "q": // Quit the application
			return m, tea.Quit
		case "up", "k": // Move cursor up
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j": // Move cursor down
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "What would you like to buy at the market?\n\n"

	// Iterating through the choices
	for i, choice := range m.choices {

		// Is the cursor pointing at this choice?
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		// Is this choice selected?
		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		// render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"
	return s
}

func initialModel() model {
	return model{
		choices:  []string{"Buy Pizza", "Buy Bread sticks", "Buy Corn dogs"},
		selected: make(map[int]struct{}),
	}
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("error at start up! Exiting...\n")
		os.Exit(1)
	}
}
