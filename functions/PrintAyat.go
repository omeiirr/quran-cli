package functions

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
)

func PrintAyat(surahNo int, ayatNo int) {

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

	if data.Cfg.PrintAyat.ShowArabic {

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
