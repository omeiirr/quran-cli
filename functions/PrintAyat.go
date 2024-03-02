package functions

import (
	"fmt"

	"github.com/omeiirr/quran-cli/data"
)

func PrintAyat(surahNo int, ayatNo int) {

	fmt.Printf("%v:%v from %v (%v) ", surahNo, ayatNo, data.QuranPayload[surahNo-1].Transliteration, data.QuranPayload[surahNo-1].Translation)

	fmt.Printf(
		"\n%v\n%v\n",
		data.QuranPayload[surahNo-1].Verses[ayatNo-1].Text,
		data.QuranPayload[surahNo-1].Verses[ayatNo-1].Translation,
	)

	fmt.Printf("\nTafsir: https://quran.com/%v:%v/tafsirs/en-tafisr-ibn-kathir\n", surahNo, ayatNo)

}
