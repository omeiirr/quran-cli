# Quran CLI

Quran CLI is a command-line application designed to help users read the Quran, fuzzy search the entire Quran for keywords, access daily verses, and more.

[quran-cli-demo.webm](https://github.com/omeiirr/quran-cli/assets/54888682/0af0e82e-9f71-44cd-a027-9b58bc3bd695)


## Contents

- [Installation](https://github.com/omeiirr/quran-cli#installation)
- [Usage](https://github.com/omeiirr/quran-cli#usage)
- [License](https://github.com/omeiirr/quran-cli#license)
- [Contributing](https://github.com/omeiirr/quran-cli#contributing)

## Installation

The following options are available for installation:

1. Run the install script

2. Download a pre-compiled binary for your operating system from the releases page

3. Install it using Go

```bash
go install github.com/omeiirr/quran-cli
```

## Usage

```
❯ quran --help
A command line app to read Quran, get daily verses, search across the Quran, and more.

Usage:
  quran [flags]
  quran [command]

Available Commands:
  config      Manage configuration settings
  hadith      Interact with hadiths
  help        Help about any command
  list        List all chapters/surahs from Quran
  random      Print a random verse from the Quran
  read        Print entire chapter or a verse, depending on input
  search      Search the Quran for verses containing a given query
  version     Print the current version of quran-cli

Flags:
  -h, --help   help for quran

Use "quran [command] --help" for more information about a command.
```

```
❯ quran hadith --help
Interact with hadiths from different books and collections

Usage:
  quran hadith [flags]
  quran hadith [command]

Available Commands:
  read        Read a specific hadith from a book. Available: 'bukhari'
  search      Search the Hadiths containing a given query

Flags:
  -h, --help   help for hadith

Use "quran hadith [command] --help" for more information about a command.

```

Once installed, you can use Quran CLI with the following commands:

- `quran help` Lists all available commands.
- `quran list` Lists all chapters/surahs from the Quran. Allows you to select a chapter to read.
- `quran random` Prints a random verse from the Quran.
- `quran read [surah] [ayat]` Prints an entire chapter or a verse, depending on the input. 

   Specify the chapter number (surah) and an optional verse number (ayat).
- `quran search`: Search the Quran for verses containing a given query.

   If you want to search within a specific chapter, for instance the 3rd chapter,

   you can use `quran search --chapter 3 <query>` or `quran search -c3 <query>`
- `quran version`: Prints the version number of quran-cli.

Use `quran [command] --help` for more information about a command.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/omeiirr/quran-cli/blob/main/LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to submit issues, bugs, feature requests, and pull requests.

---

### Authors

[Omeir](github.com/omeiirr)

### Acknowledgments

Special thanks to the developers of [bubbletea](https://github.com/charmbracelet/bubbletea/), [cobra](https://github.com/spf13/cobra), [fzf](https://github.com/junegunn/fzf), and [viper](https://github.com/spf13/viper) for providing the command-line framework used in this project.

Also extending my gratitutde to the developers of [quran.json](https://github.com/risan/quran-json) for providing the content used in this app.
