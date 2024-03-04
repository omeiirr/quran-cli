package functions

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
)

func PrintAyat(surahNo int, ayatNo int, showArabic bool) {

	if surahNo < 1 || surahNo > 114 {
		fmt.Println("Chapter not found; enter a valid chapter number between 1 to 114")
		return
	}
	if ayatNo < 1 || ayatNo > data.QuranPayload[surahNo-1].TotalVerses {

		fmt.Printf("Verse not found; chapter %v has a total of %d verses.", data.QuranPayload[surahNo-1].Transliteration, data.QuranPayload[surahNo-1].TotalVerses)
		return

	}

	fmt.Println(
		lipgloss.NewStyle().
			Foreground(lipgloss.Color(data.Cfg.ThemeColor)).
			Bold(true).
			Render(
				fmt.Sprintf(
					"%v:%v from %v (%v) | %v ", surahNo, ayatNo, data.QuranPayload[surahNo-1].Transliteration, data.QuranPayload[surahNo-1].Translation, data.QuranPayload[surahNo-1].Type,
				),
			),
	)

	if data.Cfg.PrintAyat.ShowArabic || showArabic {

		fmt.Printf(
			"\n%v\n%v\n",
			data.QuranPayload[surahNo-1].Verses[ayatNo-1].Text,
			data.QuranPayload[surahNo-1].Verses[ayatNo-1].Translation,
		)
	} else {

		fmt.Printf(
			"\n%v\n",
			data.QuranPayload[surahNo-1].Verses[ayatNo-1].Translation,
		)
	}

	fmt.Printf("\nTafsir: https://quran.com/%v:%v/tafsirs/%v\n", surahNo, ayatNo, data.Cfg.Tafsir)

}
