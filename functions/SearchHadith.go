package functions

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/omeiirr/quran-cli/data"
)

func SearchHadith(query string, exactMatch bool, chapterNo int) error {

	var hadithsToSearch []string

	hadiths := data.HadithPayload.Hadiths
	for _, hadith := range hadiths {
		// hadithString := fmt.Sprintf("%s [%d]", hadith.English.Text, hadith.ID)
		hadithString := fmt.Sprintf("%s %s [%d]", hadith.English.Narrator, hadith.English.Text, hadith.ID)
		hadithsToSearch = append(hadithsToSearch, hadithString)
	}

	// Filter the verses that contain the given query
	var matchingHadiths []string
	for _, verse := range hadithsToSearch {
		if strings.Contains(strings.ToLower(verse), strings.ToLower(query)) {
			matchingHadiths = append(matchingHadiths, verse)
		}
	}

	// Display the matching verses using FZF
	if len(matchingHadiths) == 0 {
		fmt.Println("No matching verses found.")
		return nil
	}

	// Use FZF to display the matching verses
	fzfCmd := exec.Command("fzf", "--reverse", "-m", fmt.Sprintf("--color=hl:%s,hl+:%s,pointer:%s", data.Cfg.ThemeColor, data.Cfg.ThemeColor, data.Cfg.ThemeColor))

	// Append the --exact flag if exactMatch is true
	if data.Cfg.Search.ExactMatch || exactMatch {
		fzfCmd.Args = append(fzfCmd.Args, "--exact")
	}

	fzfCmd.Stdin = strings.NewReader(strings.Join(matchingHadiths, "\n"))
	fzfCmd.Stdout = os.Stdout
	fzfCmd.Stderr = os.Stderr

	err := fzfCmd.Run()
	if err != nil {
		return fmt.Errorf("error running FZF: %v", err)
	}

	return nil
}
