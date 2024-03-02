package data

import (
	"encoding/json"
	"log"
	"os"

	"github.com/omeiirr/quran-cli/models"
)

var ChaptersPayload []models.Surah
var QuranPayload []models.Surah

func LoadData() {
	// Load list of all Surahs
	chaptersContent, err := os.ReadFile("./data/Chapters.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(chaptersContent, &ChaptersPayload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	// Load Quran
	quranContent, err := os.ReadFile("./data/Quran.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(quranContent, &QuranPayload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
