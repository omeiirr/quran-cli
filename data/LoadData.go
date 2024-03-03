package data

import (
	"encoding/json"
	"log"

	"github.com/omeiirr/quran-cli/models"
)

var ChaptersPayload []models.Surah
var QuranPayload []models.Surah

func LoadData(chaptersContent, quranContent []byte) {

	err := json.Unmarshal(chaptersContent, &ChaptersPayload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	err = json.Unmarshal(quranContent, &QuranPayload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
