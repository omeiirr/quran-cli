package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/omeiirr/quran-cli/data"
)

func SearchText(query string) error {

	// Concatenate all verses into a single slice of strings
	var allVerses []string
	for _, surah := range data.QuranPayload {
		for _, ayat := range surah.Verses {
			verseString := fmt.Sprintf("%d:%d \t %s", surah.Id, ayat.Id, ayat.Translation)
			allVerses = append(allVerses, verseString)
		}
	}

	// Filter the verses that contain the given query
	var matchingVerses []string
	for _, verse := range allVerses {
		if strings.Contains(strings.ToLower(verse), strings.ToLower(query)) {
			matchingVerses = append(matchingVerses, verse)
		}
	}

	// Display the matching verses using FZF
	if len(matchingVerses) == 0 {
		fmt.Println("No matching verses found.")
		return nil
	}

	// Use FZF to display the matching verses
	fzfCmd := exec.Command("fzf")
	fzfCmd.Stdin = strings.NewReader(strings.Join(matchingVerses, "\n"))
	fzfCmd.Stdout = os.Stdout
	fzfCmd.Stderr = os.Stderr

	err := fzfCmd.Run()
	if err != nil {
		return fmt.Errorf("error running FZF: %v", err)
	}

	return nil
}
