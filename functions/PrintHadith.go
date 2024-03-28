package functions

import (
	"fmt"

	"github.com/omeiirr/quran-cli/data"
	"github.com/omeiirr/quran-cli/models"
)

func PrintHadith(book string, hadithID int, showArabic bool) {

	if book != "bukhari" {
		fmt.Println("Book not found; available books: \n 1. bukhari")
		return
	}

	var hadith models.Hadith

	for _, h := range data.HadithPayload.Hadiths {
		if h.IDInBook == hadithID {
			hadith = h
			break
		}
	}

	if hadith.ID == 0 {
		fmt.Println("Hadith not found.")
		return
	}

	fmt.Println(hadith.English.Narrator)
	fmt.Println(hadith.English.Text)
	fmt.Printf("\n\nMore info: https://sunnah.com/%s:%d\n", book, hadith.IDInBook)
}
