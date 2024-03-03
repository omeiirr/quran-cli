package main

import (
	"fmt"

	"github.com/omeiirr/quran-cli/cmd"
	"github.com/omeiirr/quran-cli/data"
	"github.com/spf13/viper"

	_ "embed"
)

//go:embed data/Chapters.json
var chaptersContent []byte

//go:embed data/Quran.json
var quranContent []byte

func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Unable to read config:", err)
	}
	data.LoadData(chaptersContent, quranContent)
}

func main() {
	cmd.Execute()
}
