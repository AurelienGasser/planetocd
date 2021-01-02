package articles

import "time"

// Article ...
type Article struct {
	ID             int
	Lang           string
	Title          string
	Markdown       string
	OriginalURL    string
	OriginalAuthor string
	OriginalTitle  string
	Image          string
	PublishedDate  time.Time
}

// ArticleMetadata ...
type ArticleMetadata struct {
	OriginalURL    string                             `json:"originalUrl"`
	OriginalTitle  string                             `json:"originalTitle"`
	OriginalAuthor string                             `json:"originalAuthor"`
	Languages      map[string]ArticleLanguageMetadata `json:"languages"`
	Image          string                             `json:"image"`
	PublishedDate  time.Time                          `json:"published_date"`
}

// ArticleLanguageMetadata ...
type ArticleLanguageMetadata struct {
	Title string `json:"title"`
}
