# Quran CLI

Quran CLI is a command-line application designed to help users read the Quran, fuzzy search the entire Quran for keywords, access daily verses, and more.

## Contents

1. Introduction
2. Installation
3. Pre-requisites
4. Usage
5. License
6. Acknowledgements

## Installation

1. Run the install script

2. Download a pre-compiled binary for your operating system from the releases page

3. Install it using Go

```bash
go install github.com/omeiirr/quran-cli
```

Usage

Once installed, you can use Quran CLI with the following commands:

`quran help` Lists all available commands.

`quran list` Lists all chapters/surahs from the Quran. Allows you to select a chapter to read.

`quran random` Prints a random verse from the Quran.

`quran read [surah] [ayat]` Prints an entire chapter or a verse, depending on the input.

Specify the chapter number (surah) and an optional verse number (ayat).

`quran search`: Search the Quran for verses containing a given query.

If you want to search within a specific chapter, for instance the 3rd chapter, , you can use `quran search --chapter 3 <query>` or `quran search -c3 <query>`

`quran version`: Prints the version number of quran-cli.

Use `quran [command] --help` for more information about a command.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contributing

Contributions are welcome! Please feel free to submit issues, bugs, feature requests, and pull requests.

---

### Authors

[Omeir](github.com/omeiirr)

### Acknowledgments

Special thanks to the developers of [bubbletea](https://github.com/charmbracelet/bubbletea/), [cobra](https://github.com/spf13/cobra), [fzf](https://github.com/junegunn/fzf), and [viper](https://github.com/spf13/viper) for providing the command-line framework used in this project.

Also extending my gratitutde to the developers of [quran.json](https://github.com/risan/quran-json) for providing the content used in this app.
