package main

import (
	"github.com/omeiirr/quran-cli/cmd"
	"github.com/omeiirr/quran-cli/data"

	_ "embed"
)

//go:embed data/Chapters.json
var chaptersContent []byte

//go:embed data/Quran.json
var quranContent []byte

func init() {
	data.LoadData(chaptersContent, quranContent)
}

func main() {
	cmd.Execute()
}
