package functions

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
)

func PrintSurah(surahNo int, showArabic bool) {

	if surahNo < 1 || surahNo > 114 {
		fmt.Println("Chapter not found; enter a valid chapter number between 1 to 114")
		return
	}

	fmt.Println(

		lipgloss.NewStyle().
			Foreground(lipgloss.Color(data.Cfg.ThemeColor)).
			Bold(true).
			Render(
				fmt.Sprintf(
					"\nSurah %v | %v (%v) | %v\n",
					data.QuranPayload[surahNo-1].Id,
					data.QuranPayload[surahNo-1].Transliteration,
					data.QuranPayload[surahNo-1].Translation,
					data.QuranPayload[surahNo-1].Type,
				),
			),
	)

	// Print Id and Translation of each Ayat
	for _, ayat := range data.QuranPayload[surahNo-1].Verses {
		if data.Cfg.PrintSurah.ShowArabic || showArabic {
			fmt.Printf("%s \n%d. %s\n\n", ayat.Text, ayat.Id, ayat.Translation)
		} else {
			fmt.Printf("%d. %s\n", ayat.Id, ayat.Translation)

		}
	}

}
