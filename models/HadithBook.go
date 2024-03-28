package models

// Define structs to represent the JSON structure

type HadithBook struct {
	ID       int       `json:"id"`
	Metadata Metadata  `json:"metadata"`
	Chapters []Chapter `json:"chapters"`
	Hadiths  []Hadith  `json:"hadiths"`
}

type Metadata struct {
	ID      int   `json:"id"`
	Length  int   `json:"length"`
	Arabic  Title `json:"arabic"`
	English Title `json:"english"`
}

type Title struct {
	Title        string `json:"title"`
	Author       string `json:"author"`
	Introduction string `json:"introduction"`
}

type Chapter struct {
	ID      int    `json:"id"`
	BookID  int    `json:"bookId"`
	Arabic  string `json:"arabic"`
	English string `json:"english"`
}

type Hadith struct {
	ID        int    `json:"id"`
	Arabic    string `json:"arabic"`
	English   Text   `json:"english"`
	ChapterID int    `json:"chapterId"`
	BookID    int    `json:"bookId"`
	IDInBook  int    `json:"idInBook"`
}

type Text struct {
	Narrator string `json:"narrator"`
	Text     string `json:"text"`
}
