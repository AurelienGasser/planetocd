package articles

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
}

// ArticleMetadata ...
type ArticleMetadata struct {
	OriginalURL    string                             `json:"originalUrl"`
	OriginalTitle  string                             `json:"originalTitle"`
	OriginalAuthor string                             `json:"originalAuthor"`
	Languages      map[string]ArticleLanguageMetadata `json:"languages"`
	Image          string                             `json:"image"`
}

// ArticleLanguageMetadata ...
type ArticleLanguageMetadata struct {
	Title string `json:"title"`
}
