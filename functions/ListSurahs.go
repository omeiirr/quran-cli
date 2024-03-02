package functions

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m,
				tea.Batch(
					tea.Printf("Opening Surah %s. . .", m.table.SelectedRow()[2]),
				)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func ListSurahs() {

	columns := []table.Column{
		{Title: "#", Width: 5},
		{Title: "Surah", Width: 12},
		{Title: "Chapter Arabic", Width: 15},
		{Title: "Chapter English", Width: 30},
		{Title: "Verses", Width: 10},
		{Title: "Type", Width: 10},
	}

	var rows []table.Row
	for _, surah := range data.ChaptersPayload {
		row := []string{
			fmt.Sprintf("%d", surah.Id),
			surah.Name,
			fmt.Sprintf(surah.Transliteration),
			fmt.Sprintf(surah.Translation),
			// fmt.Sprintf("%v (%v)", surah.Transliteration, surah.Translation),
			fmt.Sprintf("%d", surah.TotalVerses),
			surah.Type,
		}
		rows = append(rows, row)
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(true).
		Foreground(lipgloss.Color("#04B575"))

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#04B575")).
		Bold(false)

	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
