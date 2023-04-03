package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/tomyfalgui/game-of-life-go/board"
)

type tickMsg time.Time

type model struct {
	grid [][]int
}

func initialModel() model {
	// newBoard, err := board.GenerateRandom(150, 80)
	// if err != nil {
	// 	log.Fatal("could not generate board")
	// }
	return model{
		grid: GosperGlider(),
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tickCmd(), tea.EnterAltScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}

	case tickMsg:
		m.grid = board.NextState(m.grid)
		return m, tickCmd()

	}
	return m, nil
}

func (m model) View() string {
	// The header
	s := "Conway's Game of Life\n\n"
	// Iterate over our choices
	for _, y := range m.grid {
		for _, cell := range y {
			if cell == 0 {
				s += " "
			} else {
				c := lipgloss.NewStyle().SetString(" ").Background(lipgloss.Color("#FF7B9C"))
				s += c.String()
			}
		}
		s += "\n"
	}

	// The footer
	s += "\nPress q to quit.\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*200, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
