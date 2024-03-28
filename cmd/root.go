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
	readCmd.Flags().BoolP("arabic", "a", false, "Shows Arabic text along with the English translation")
	readCmd.Flags().BoolP("interactive", "i", false, "Shows the surah in an interactive window, instead of printing out to the console [experimental feature]")

	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().BoolP("exact", "e", false, "Uses exact match for keyword instead of fuzzy match")
	searchCmd.Flags().IntP("chapter", "c", 0, "Search only within the specified chapter (default: search in the whole Quran)")

	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolP("list", "l", false, "List current configuration settings")
	configCmd.Flags().BoolP("edit", "e", false, "Edit the configuration file")

	rootCmd.AddCommand(versionCmd)

	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	// Hadith commands
	rootCmd.AddCommand(hadithCmd)
	hadithCmd.AddCommand(readHadithCmd)
	hadithCmd.AddCommand(searchHadithCmd)
	searchHadithCmd.Flags().BoolP("exact", "e", false, "Uses exact match for keyword instead of fuzzy match")

}

var rootCmd = &cobra.Command{
	Use:   "quran",
	Short: "A command line app to read Quran.",
	Long:  `A command line app to read Quran, get daily verses, search across the Quran, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		functions.WelcomeScreen()
		fmt.Println(`  Assalamu alaikum warahmatullahi wabarakatuhu.

  Use "quran help" for all available commands.
  Use "quran [command] --help" for more information about a command.`)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all chapters/surahs from Quran",
	Long: `List all chapters/surahs from Quran in an interactive table. Select a chapter to read.
Use up/down arrow keys or k/j to move up down.
Press Enter to read the highlighted chapter.
Press q to exit the table.
	`,
	Args:    cobra.NoArgs,
	Aliases: []string{"l"},
	Run: func(cmd *cobra.Command, args []string) {
		functions.ListSurahs()
	},
}

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "Print a random verse from the Quran",
	Long:  "Print a random verse from the Quran",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		functions.SelectRandomVerse()
	},
}

var readCmd = &cobra.Command{
	Use:   "read surah [ayat]",
	Short: "Print entire chapter or a verse, depending on input",
	Long: `Print entire chapter or a verse, depending on input.
First argument is the chapter number.
Second optional argument is the verse number.	
	`,
	Args:    cobra.RangeArgs(1, 2),
	Aliases: []string{"r"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(`Not enough arguments. Use "quran read --help" for more`)
			return
		}

		// convert string to int
		surahNo, err := strconv.Atoi(args[0])
		if err != nil || surahNo > 114 {
			fmt.Println("Chapter not found; enter a valid chapter number between 1 to 114")
			return
		}

		showArabic, _ := cmd.Flags().GetBool("arabic")
		interactive, _ := cmd.Flags().GetBool("interactive")

		switch len(args) {
		case 1:
			if interactive {
				functions.ViewSurah(surahNo, showArabic)
			} else {
				functions.PrintSurah(surahNo, showArabic)
			}

		case 2:
			ayatNo, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Printf("Verse not found; chapter %v has a total of %d verses.\n", data.QuranPayload[surahNo-1].Transliteration, data.QuranPayload[surahNo-1].TotalVerses)
				return
			}
			functions.PrintAyat(surahNo, ayatNo, showArabic)

		default:
			fmt.Println(`Too many arguments. Use "quran read --help" for more`)

		}

	},
}
var searchCmd = &cobra.Command{
	Use:     "search [query]",
	Short:   "Search the Quran for verses containing a given query",
	Long:    `Search the Quran for verses containing a given query using fzf (both fuzzy search and exact match is possible).`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		exactMatch, _ := cmd.Flags().GetBool("exact")
		chapterNo, _ := cmd.Flags().GetInt("chapter")

		var err error
		if len(args) == 0 {
			err = functions.SearchText("", exactMatch, chapterNo)
		} else {
			err = functions.SearchText(args[0], exactMatch, chapterNo)
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration settings",
	Long:  `Manage configuration settings for the quran-cli application.`,
	Run: func(cmd *cobra.Command, args []string) {
		listFlag, _ := cmd.Flags().GetBool("list")
		editFlag, _ := cmd.Flags().GetBool("edit")

		if listFlag {
			functions.ListConfigSettings()
		} else if editFlag {
			functions.EditConfigSettings()
		} else {
			// Display help if no flags are provided
			cmd.Help()
		}
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the current version of quran-cli",
	Long:  `Print the current version of quran-cli`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("quran-cli \t v.0.1")
	},
}

var hadithCmd = &cobra.Command{
	Use:   "hadith",
	Short: "Read or search for hadiths",
	Long:  "Read or search for hadiths from different books and collections",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please specify a subcommand. Use 'quran hadith --help' for more information.")
	},
}

var readHadithCmd = &cobra.Command{
	Use:   "read [book] [id]",
	Short: "Read a specific hadith from a book",
	Long:  "Read a specific hadith from a book by providing the book name and hadith ID",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid number of arguments. Usage: quran hadith read [book] [id]")
			return
		}

		bookName := args[0]
		hadithID, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Invalid Hadith ID")
			return
		}

		functions.PrintHadith(bookName, hadithID, true)
	},
}

var searchHadithCmd = &cobra.Command{
	Use:     "search [query]",
	Short:   "Search the Hadiths containing a given query",
	Long:    `Search the Hadiths containing a given query using fzf (both fuzzy search and exact match is possible).`,
	Aliases: []string{"s"},
	Run: func(cmd *cobra.Command, args []string) {
		exactMatch, _ := cmd.Flags().GetBool("exact")
		chapterNo, _ := cmd.Flags().GetInt("chapter")

		var err error
		if len(args) == 0 {
			err = functions.SearchHadith("", exactMatch, chapterNo)
		} else {
			err = functions.SearchHadith(args[0], exactMatch, chapterNo)
		}
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
