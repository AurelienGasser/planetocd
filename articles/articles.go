package articles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
)

var articles = make(map[string]map[int]*Article)
var articlesRootPath = "articles/articles/"
var regexTranslationFile = regexp.MustCompile(`_([a-z]{2}).md$`)
var regexMetadataFile = regexp.MustCompile(`.*/([0-9]+)_[^/]+.json$`)

// GetArticles ...
func GetArticles() map[string]map[int]*Article {
	ensureLoaded()
	return articles
}

func ensureLoaded() {
	if len(articles) != 0 {
		return
	}

	articles, _ := filepath.Glob(articlesRootPath + "*__*.json")
	for _, metadataFile := range articles {
		loadArticle(metadataFile)
	}
}

func loadArticle(metadataFile string) {
	// Get article ID
	matches := regexMetadataFile.FindStringSubmatch(metadataFile)
	if len(matches) != 2 {
		log.Panic("Error parsing file name " + metadataFile)
	}
	id := matches[1]
	idN, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	// Load metadata
	metadataJSON, err := ioutil.ReadFile(metadataFile)
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
		if len(lang) != 2 {
			log.Panic("Invalid lang: " + lang)
		}
		article, err := loadArticleInLang(id, idN, lang, metadata)
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, ok := articles[lang]
		if !ok {
			articles[lang] = make(map[int]*Article)
		}
		articles[lang][idN] = article
	}
}

func loadArticleInLang(id string, idN int, lang string, metadata ArticleMetadata) (*Article, error) {
	langMetadata := metadata.Languages[lang]
	mdPages := []string{}

	for _, fileName := range langMetadata.Pages {
		filePath := articlesRootPath + fileName
		mdBytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("cannot read file %v: %v", filePath, err)
		}
		mdPages = append(mdPages, string(mdBytes))
	}
	return &Article{
		ID:             idN,
		Lang:           lang,
		Title:          langMetadata.Title,
		MarkdownPages:  mdPages,
		OriginalURL:    metadata.OriginalURL,
		OriginalTitle:  metadata.OriginalTitle,
		OriginalAuthor: metadata.OriginalAuthor,
		Image:          metadata.Image,
		PublishedDate:  metadata.PublishedDate,
		Translators:    langMetadata.Translators,
		Tags:           metadata.Tags,
	}, nil
}
