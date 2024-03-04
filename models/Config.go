package models

type Config struct {
	ThemeColor string           `mapstructure:"theme_color"`
	PrintSurah PrintSurahConfig `mapstructure:"print_surah"`
	PrintAyat  PrintAyatConfig  `mapstructure:"print_ayat"`
	Tafsir     string           `mapstructure:"tafsir"`
}

type PrintSurahConfig struct {
	ShowArabic bool `mapstructure:"show_arabic"`
}

type PrintAyatConfig struct {
	ShowArabic bool `mapstructure:"show_arabic"`
}
