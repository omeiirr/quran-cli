package functions

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
)

const useHighPerformanceRenderer = true

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"

		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"

		return titleStyle.Copy().BorderStyle(b)
	}()

	lineStyle = func() lipgloss.Style {
		return lipgloss.NewStyle()
	}()
)

type SurahModel struct {
	surahNo  int
	content  []string
	ready    bool
	viewport viewport.Model
}

func ViewSurah(surahNo int, showArabic bool) {

	// Concatenate verses into a slice of strings for the specified Surah
	var content []string
	if surahNo > 0 {
		surah := data.QuranPayload[surahNo-1]
		// fmt.Print(len(surah.Verses))
		content = append(content, "\n")
		for _, ayat := range surah.Verses {
			verseString := fmt.Sprintf("%d:%d  %s", surah.Id, ayat.Id, ayat.Translation)
			content = append(content, verseString)
		}
	}

	p := tea.NewProgram(
		SurahModel{content: content, surahNo: surahNo},
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}

func (m SurahModel) Init() tea.Cmd {
	return nil
}

func (m SurahModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if k := msg.String(); k == "ctrl+c" || k == "q" || k == "esc" {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(strings.Join(m.content, "\n\n"))

			m.ready = true

			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m SurahModel) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m SurahModel) headerView() string {
	title := titleStyle.Render(
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(data.Cfg.ThemeColor)).
			Render(
				fmt.Sprintf(
					"\nSurah %v | %v (%v) | %v\n",
					data.QuranPayload[m.surahNo-1].Id,
					data.QuranPayload[m.surahNo-1].Transliteration,
					data.QuranPayload[m.surahNo-1].Translation,
					data.QuranPayload[m.surahNo-1].Type,
				),
			),
	)

	line := lineStyle.
		Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title))))

	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m SurahModel) footerView() string {
	info := infoStyle.
		Foreground(lipgloss.Color(data.Cfg.ThemeColor)).
		Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))

	line := lineStyle.
		Render(strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info))))

	return lipgloss.JoinVertical(lipgloss.Center, lipgloss.JoinHorizontal(lipgloss.Center, line, info),
		`Use arrows, PgUp/PgDn, or scroll to navigate. Press q or esc to quit.`)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
