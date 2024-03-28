package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/omeiirr/quran-cli/data"
)

func SearchText(query string, exactMatch bool, chapterNo int) error {

	if chapterNo > 114 {
		return fmt.Errorf("Chapter not found; enter a valid chapter number between 1 to 114")
	}

	// Concatenate verses into a slice of strings for the specified Surah
	var versesToSearch []string
	if chapterNo > 0 {
		surah := data.QuranPayload[chapterNo-1]
		for _, ayat := range surah.Verses {
			verseString := fmt.Sprintf("%d:%d \t %s", surah.Id, ayat.Id, ayat.Translation)
			versesToSearch = append(versesToSearch, verseString)
		}
	} else {
		// Concatenate all verses into a single slice of strings
		for _, surah := range data.QuranPayload {
			for _, ayat := range surah.Verses {
				verseString := fmt.Sprintf("%d:%d \t %s", surah.Id, ayat.Id, ayat.Translation)
				versesToSearch = append(versesToSearch, verseString)
			}
		}
	}

	// Filter the verses that contain the given query
	var matchingVerses []string
	for _, verse := range versesToSearch {
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
	fzfCmd := exec.Command("fzf", "-m", "--reverse", fmt.Sprintf("--color=hl:%s,hl+:%s,pointer:%s", data.Cfg.ThemeColor, data.Cfg.ThemeColor, data.Cfg.ThemeColor))

	// Append the --exact flag if exactMatch is true
	if data.Cfg.Search.ExactMatch || exactMatch {
		fzfCmd.Args = append(fzfCmd.Args, "--exact")
	}

	fzfCmd.Stdin = strings.NewReader(strings.Join(matchingVerses, "\n"))
	fzfCmd.Stdout = os.Stdout
	fzfCmd.Stderr = os.Stderr

	err := fzfCmd.Run()
	if err != nil {
		return fmt.Errorf("error running FZF: %v", err)
	}

	return nil
}
