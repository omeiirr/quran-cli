package functions

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
	"github.com/spf13/viper"
)

func PrintSurah(surahNo int) {

	// if len(verseRange) == 0 {
	// 	verseRange[0] = 1
	// 	verseRange[1] = data.ChaptersPayload[surahNo].TotalVerses
	// }

	// fmt.Println(data.ChaptersPayload[surahNo].TotalVerses)
	// fmt.Println(verseRange)

	fmt.Println(

		lipgloss.NewStyle().
			Foreground(lipgloss.Color(data.ThemeColor)).
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

	// fmt.Printf("\nSurah %v", data.QuranPayload[surahNo-1].Translation)

	// fmt.Println(data.QuranPayload[surahNo-1].Verses)

	// Print Id and Translation of each Ayat
	showArabic := viper.GetBool("print_surah.show_arabic")
	for _, ayat := range data.QuranPayload[surahNo-1].Verses {
		if showArabic {
			fmt.Printf("%s \n%d. %s\n\n", ayat.Text, ayat.Id, ayat.Translation)
			// fmt.Printf("%d. %s\n%s\n\n", ayat.Id, ayat.Text, ayat.Translation)
		} else {
			fmt.Printf("%d. %s\n", ayat.Id, ayat.Translation)

		}
	}

	// Print Id and Translation of each Ayat
	// for i := verseRange[0]; i <= verseRange[1]; i++ {
	// 	fmt.Printf("%d. %s\n", data.QuranPayload[surahNo].Verses[i].Id, data.QuranPayload[surahNo].Verses[i].Translation)

	// }
}
