package articles

import "time"

// Article describes an article in a specific language
type Article struct {
	ID               int
	Lang             string
	Title            string
	MarkdownPages    []string
	OriginalURL      string
	OriginalAuthor   string
	OriginalTitle    string
	Image            string
	PublishedDate    time.Time
	Translators      []ArticleTranslator
	Tags             []string
	OriginalMetadata ArticleMetadata
}

// ArticleMetadata ...
type ArticleMetadata struct {
	OriginalURL    string                             `json:"originalUrl"`
	OriginalTitle  string                             `json:"originalTitle"`
	OriginalAuthor string                             `json:"originalAuthor"`
	Languages      map[string]ArticleLanguageMetadata `json:"languages"`
	Image          string                             `json:"image"`
	PublishedDate  time.Time                          `json:"published_date"`
	Tags           []string                           `json:"tags"`
}

// ArticleLanguageMetadata ...
type ArticleLanguageMetadata struct {
	Title       string              `json:"title"`
	Pages       []string            `json:"pages"`
	Translators []ArticleTranslator `json:"translators"`
}

type ArticleTranslator struct {
	Name  string `json:"name"`
	IsBot bool   `json:"isBot"`
}
