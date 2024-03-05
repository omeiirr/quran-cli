package models

type Config struct {
	ThemeColor string           `mapstructure:"theme_color"`
	Tafsir     string           `mapstructure:"tafsir"`
	PrintSurah PrintSurahConfig `mapstructure:"print_surah"`
	PrintAyat  PrintAyatConfig  `mapstructure:"print_ayat"`
	Search     SearchConfig     `mapstructure:"search"`
}

type PrintSurahConfig struct {
	ShowArabic bool `mapstructure:"show_arabic"`
}

type PrintAyatConfig struct {
	ShowArabic bool `mapstructure:"show_arabic"`
}

type SearchConfig struct {
	ExactMatch bool `mapstructure:"exact_match"`
}
