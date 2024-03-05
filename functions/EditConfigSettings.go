package functions

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func EditConfigSettings() {
	configFile := getConfigFilePath()
	editor := getEditor()

	var editCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		editCmd = exec.Command("cmd", "/c", "start", editor, configFile)
	default:
		editCmd = exec.Command(editor, configFile)
		editCmd.Stdin = os.Stdin
		editCmd.Stdout = os.Stdout
		editCmd.Stderr = os.Stderr
	}

	err := editCmd.Run()
	if err != nil {
		fmt.Println("Error opening editor:", err)
		os.Exit(1)
	}
}

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user's home directory:", err)
		os.Exit(1)
	}
	configFilePath := filepath.Join(homeDir, ".quran", "config.yaml")
	return configFilePath
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		switch runtime.GOOS {
		case "windows":
			editor = "notepad"
		default:
			editor = "vi" // Default to vi for Unix-like systems
		}
	}
	return editor
}
