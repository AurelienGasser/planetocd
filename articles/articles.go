package articles

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

//go:embed all:articles
var articlesFS embed.FS
var articlesRootPath = "articles/"
var articles = make(map[string]map[int]*Article)
var regexMetadataFile = regexp.MustCompile(`^([0-9]+)_[^/]+.json$`)

// GetArticles ...
func GetArticles() map[string]map[int]*Article {
	ensureLoaded()
	return articles
}

func ensureLoaded() {
	if len(articles) != 0 {
		return
	}

	dirEntries, err := articlesFS.ReadDir("articles")
	if err != nil {
		panic(err)
	}
	for _, dirEntry := range dirEntries {
		matches := regexMetadataFile.FindStringSubmatch(dirEntry.Name())
		if len(matches) != 2 {
			continue
		}
		id := matches[1]
		idN, err := strconv.Atoi(id)
		if err != nil {
			continue
		}
		loadArticle(idN, dirEntry.Name())
	}
}

func loadArticle(id int, metadataFile string) {
	// Load metadata
	metadataJSON, err := articlesFS.ReadFile(articlesRootPath + metadataFile)
	if err != nil {
		log.Fatal(err)
	}
	var metadata ArticleMetadata
	err = json.Unmarshal(metadataJSON, &metadata)
	if err != nil {
		log.Fatal(err)
	}

	// Load each language
	for lang := range metadata.Languages {
		if lang == "zh" {
			return
		}
		if len(lang) != 2 {
			log.Panic("Invalid lang: " + lang)
		}
		article, err := loadArticleInLang(id, lang, metadata)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, ok := articles[lang]
		if !ok {
			articles[lang] = make(map[int]*Article)
		}
		articles[lang][id] = article
	}
}

func loadArticleInLang(id int, lang string, metadata ArticleMetadata) (*Article, error) {
	langMetadata := metadata.Languages[lang]
	mdPages := []string{}

	for _, fileName := range langMetadata.Pages {
		filePath := articlesRootPath + fileName
		mdBytes, err := articlesFS.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("cannot read file %v: %v", filePath, err)
		}
		mdPages = append(mdPages, string(mdBytes))
	}
	return &Article{
		ID:               id,
		Lang:             lang,
		Title:            langMetadata.Title,
		MarkdownPages:    mdPages,
		OriginalURL:      metadata.OriginalURL,
		OriginalTitle:    metadata.OriginalTitle,
		OriginalAuthor:   metadata.OriginalAuthor,
		Image:            metadata.Image,
		PublishedDate:    metadata.PublishedDate,
		Translators:      langMetadata.Translators,
		Tags:             metadata.Tags,
		OriginalMetadata: metadata,
	}, nil
}
