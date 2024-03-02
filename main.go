package main

import (
	"github.com/omeiirr/quran-cli/cmd"
	"github.com/omeiirr/quran-cli/data"
)

func init() {
	data.LoadData()
}

func main() {
	cmd.Execute()
}
