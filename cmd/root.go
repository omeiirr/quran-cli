package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/omeiirr/quran-cli/data"
	"github.com/omeiirr/quran-cli/functions"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(randomCmd)
	rootCmd.AddCommand(readCmd)
	// rootCmd.AddCommand(chapterCmd)
	// rootCmd.AddCommand(verseCmd)
	rootCmd.AddCommand(versionCmd)
}

var rootCmd = &cobra.Command{
	Use:   "quran",
	Short: "A command line app to read Quran.",
	Long:  `A command line app to read Quran, get daily verses, read chapters, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		functions.WelcomeScreen()
		fmt.Println(`
  Assalamu alaikum warahmatullahi wabarakatuhu.
  Use "quran help" for all available commands.
  Use "quran [command] --help" for more information about a command.`)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all chapters/surahs from Quran.",
	Long: `
	Lists all chapters/surahs from Quran in an interactive table. Select a chapter to read.
	Use up/down arrow keys or k/j to move up down.
	Press Enter to read the highlighted chapter.
	Press q to exit the table.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		functions.ListSurahs()
	},
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Print a random verse from the Quran.",
	Long:  `Print a random verse from the Quran.`,
	Run: func(cmd *cobra.Command, args []string) {
		functions.SelectRandomVerse()
	},
}

var readCmd = &cobra.Command{
	Use:   "read surah [ayat]",
	Short: "Prints entire chapter or a verse, depending on input.",
	Long: `
	Prints entire chapter or a verse, depending on input.
	First argument is the chapter number.
	Second optional argument is the verse number.	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// string to int
		surahNo, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Chapter not found; enter a valid chapter number between 1 to 114")
			return
		}

		switch len(args) {
		case 0:
			fmt.Println("Too few arguments")

		case 1:
			functions.PrintSurah(surahNo)

		case 2:
			ayatNo, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Printf("Verse not found; chapter %v has a total of %d verses.", data.QuranPayload[surahNo-1].Transliteration, data.QuranPayload[surahNo-1].TotalVerses)
				return
			}
			functions.PrintAyat(surahNo, ayatNo)
		// case 3:
		// 	start, err := strconv.Atoi(args[1])
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	end, err := strconv.Atoi(args[2])
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	x := [2]int{start, end}
		// 	fmt.Println(x)
		// 	functions.PrintSurah(surahNo)
		// for i := start; i <= end; i++ {
		// 	functions.PrintAyat(surahNo, i)
		// }

		default:
			fmt.Println("Too many arguments")

		}

	},
}

// var chapterCmd = &cobra.Command{
// 	Use:   "chapter -n",
// 	Short: "Print the chapter acc to the number provided.",
// 	Long:  `Print the chapter acc to the number provided.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// string to int
// 		i, err := strconv.Atoi(args[0])
// 		if err != nil {
//
// 			panic(err)
// 		}
// 		functions.PrintSurah(i)
// 	},
// }
// var verseCmd = &cobra.Command{
// 	Use:   "verse x y",
// 	Short: "Prints verse y from chapter x.",
// 	Long:  `Prints verse y from chapter x.`,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		// string to int
// 		surahNo, err := strconv.Atoi(args[0])
// 		if err != nil {
//
// 			panic(err)
// 		}
// 		ayatNo, err := strconv.Atoi(args[1])
// 		if err != nil {
//
// 			panic(err)
// 		}
// 		functions.PrintAyat(surahNo, ayatNo)
// 	},
// }

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of quran-cli",
	Long:  `All software has versions. This is for quran-cli.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quran-cli \t v.0.1 -- HEAD")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
