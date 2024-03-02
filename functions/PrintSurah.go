package functions

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
)

func PrintSurah(surahNo int) {

	// if len(verseRange) == 0 {
	// 	verseRange[0] = 1
	// 	verseRange[1] = data.ChaptersPayload[surahNo].TotalVerses
	// }

	// fmt.Println(data.ChaptersPayload[surahNo].TotalVerses)
	// fmt.Println(verseRange)

	// lipgloss.NewStyle().
	// 	Foreground(lipgloss.Color("#04B575")).
	// 	Render(
	// )
	fmt.Println(

		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).
			Bold(true).
			Render(
				"\nSurah",
				data.QuranPayload[surahNo-1].Transliteration,
				data.QuranPayload[surahNo-1].Translation,
				// data.QuranPayload[surahNo-1].Type,
			),
	)

	// fmt.Printf("\nSurah %v", data.QuranPayload[surahNo-1].Translation)

	// fmt.Println(data.QuranPayload[surahNo-1].Verses)

	// Print Id and Translation of each Ayat
	for _, ayat := range data.QuranPayload[surahNo-1].Verses {
		fmt.Printf("%d. %s\n", ayat.Id, ayat.Translation)
	}

	// Print Id and Translation of each Ayat
	// for i := verseRange[0]; i <= verseRange[1]; i++ {
	// 	fmt.Printf("%d. %s\n", data.QuranPayload[surahNo].Verses[i].Id, data.QuranPayload[surahNo].Verses[i].Translation)

	// }
}
