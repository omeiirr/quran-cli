package functions

import (
	"fmt"
	"os"
	"os/exec"
)

func EditConfigSettings() {
	configFile := getConfigFilePath()
	editor := getEditor()

	// Open the configuration file in the user's preferred editor
	editCmd := exec.Command(editor, configFile)
	editCmd.Stdin = os.Stdin
	editCmd.Stdout = os.Stdout
	editCmd.Stderr = os.Stderr

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
	return homeDir + "/.quran/config.yaml"
}

func getEditor() string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vi" // Default to vi if EDITOR environment variable is not set
	}
	return editor
}
