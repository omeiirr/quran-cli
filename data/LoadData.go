package data

import (
	"encoding/json"
	"log"

	"github.com/omeiirr/quran-cli/models"
	"github.com/spf13/viper"
)

var ChaptersPayload []models.Surah
var QuranPayload []models.Surah
var HadithPayload models.HadithBook

var Cfg models.Config

func LoadData(chaptersContent, quranContent []byte) {

	err := json.Unmarshal(chaptersContent, &ChaptersPayload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	err = json.Unmarshal(quranContent, &QuranPayload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}

}

func LoadHadith(hadithContent []byte) {
	err := json.Unmarshal(hadithContent, &HadithPayload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
