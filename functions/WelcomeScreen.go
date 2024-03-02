package functions

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)
func WelcomeScreen(
){

	
	// Welcome screen
	fmt.Println(lipgloss.NewStyle().
	// Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#04B575")).
	Padding(1, 2, 1).
	Width(56).
	Margin(1, 2, 1).
	Align(lipgloss.Center).
	Render("Qur'an CLI \nIn the name of Allah,\n the Entirely Merciful, the Especially Merciful"))
}