package functions

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
)

func WelcomeScreen() {
	// Welcome screen
	fmt.Println(lipgloss.NewStyle().
		// Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color(data.Cfg.ThemeColor)).
		Padding(1, 2, 1).
		Width(56).
		Margin(1, 2, 1).
		Align(lipgloss.Center).
		Render("Qur'an CLI \nIn the name of Allah,\n the Entirely Merciful, the Especially Merciful"))
}
