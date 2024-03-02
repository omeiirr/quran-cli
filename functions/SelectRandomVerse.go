package functions

import (
	"fmt"
	"math/rand"

	"github.com/charmbracelet/lipgloss"
	"github.com/omeiirr/quran-cli/data"
	"github.com/omeiirr/quran-cli/models"
)

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func SelectRandomVerse() models.Ayat {
	surahNo := randInt(0, 114)
	ayatNo := randInt(0, data.QuranPayload[surahNo].TotalVerses)

	selectedAyat := data.QuranPayload[surahNo].Verses[ayatNo]

	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Render("Quranic verse of the day:"))

	fmt.Printf(
		"%v:%v from %v (%v)\n", surahNo+1, ayatNo+1, data.ChaptersPayload[surahNo].Transliteration, data.ChaptersPayload[surahNo].Translation)

	fmt.Println(
		lipgloss.NewStyle().
			Foreground(lipgloss.Color("#04B575")).Render(selectedAyat.Translation))

	fmt.Printf("\nTafsir: https://quran.com/%v:%v/tafsirs/en-tafisr-ibn-kathir\n", surahNo+1, ayatNo+1)
	return selectedAyat
}
