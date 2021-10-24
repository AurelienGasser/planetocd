package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	translate "cloud.google.com/go/translate/apiv3"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/aureliengasser/planetocd/articles"
	"github.com/aureliengasser/planetocd/server"
	"github.com/gomarkdown/markdown"
	"github.com/urfave/cli/v2"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

var DEFAULT_GOOGLE_APPLICATION_CREDENTIALS string = os.Getenv("PLANETOCD_GOOGLE_APPLICATION_CREDENTIALS")
var DEFAULT_INPUT_MD_FILE = "./workdir/in.md"
var DEFAULT_INPUT_HTML_FILE = "./workdir/in.html"
var DEFAULT_OUTPUT_DIR = "./articles/articles/"
var DEFAULT_PAGE_NUMBER = 1

func main() {

	var id int
	var originalTitle string
	var originalURL string
	var originalAuthor string
	var pageNumber int
	var inputFileMD string
	var inputFileHTML string
	var outPath string

	if os.Getenv("GOOGLE_APPLICATION_CREDENTIALS") == "" {
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", DEFAULT_GOOGLE_APPLICATION_CREDENTIALS)
	}

	app := &cli.App{
		Usage: "Translate an article",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "id",
				Usage:       "Output article ID",
				Required:    true,
				Destination: &id,
			},
			&cli.StringFlag{
				Name:        "title",
				Usage:       "Original article title",
				Required:    true,
				Destination: &originalTitle,
			},
			&cli.StringFlag{
				Name:        "url",
				Usage:       "Original article url",
				Required:    true,
				Destination: &originalURL,
			},
			&cli.StringFlag{
				Name:        "author",
				Usage:       "Original article Author",
				Destination: &originalAuthor,
			},
			&cli.IntFlag{
				Name:        "page",
				Usage:       "Page number",
				Value:       DEFAULT_PAGE_NUMBER,
				Destination: &pageNumber,
			},
			&cli.StringFlag{
				Name:        "input-md",
				Usage:       "Input Markdown file path",
				Value:       DEFAULT_INPUT_MD_FILE,
				Destination: &inputFileMD,
			},
			&cli.StringFlag{
				Name:        "input-html",
				Usage:       "Input Markdown HTML file path",
				Value:       DEFAULT_INPUT_HTML_FILE,
				Destination: &inputFileHTML,
			},
			&cli.StringFlag{
				Name:        "output-path",
				Usage:       "Output article directory",
				Value:       DEFAULT_OUTPUT_DIR,
				Destination: &outPath,
			},
		},
		Action: func(c *cli.Context) error {
			CreateTranslatedArticle(
				id,
				originalTitle,
				originalURL,
				originalAuthor,
				pageNumber,
				inputFileMD,
				inputFileHTML,
				outPath)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateTranslatedArticle ....
func CreateTranslatedArticle(
	id int,
	originalTitle string,
	originalURL string,
	originalAuthor string,
	pageNumber int,
	inputFileMD string,
	inputFileHTML string,
	outPath string) {

	idStr := fmt.Sprintf("%04d", id)
	inputMD, err := ioutil.ReadFile(inputFileMD)

	if err != nil {
		log.Fatal(err)
	}

	html := markdown.ToHTML(inputMD, nil, nil)
	slug := server.Slugify(originalTitle)

	metadata := articles.ArticleMetadata{
		OriginalURL:    originalURL,
		OriginalTitle:  originalTitle,
		OriginalAuthor: originalAuthor,
		Languages:      make(map[string]articles.ArticleLanguageMetadata),
		PublishedDate:  time.Now(),
	}

	var existingMetadata articles.ArticleMetadata
	metadataFilePath := path.Join(outPath, idStr+"__"+slug+".json")

	metadataFile, err := ioutil.ReadFile(metadataFilePath)
	if err != nil {
		err := json.Unmarshal(metadataFile, &existingMetadata)
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, lang := range server.SupportedLanguages {
		fileName, err := translateAndWrite(outPath, lang, string(html), idStr, pageNumber)
		if err != nil {
			log.Fatal(err)
		}
		translatedTitle, err := translateText(os.Stdout, "planetocd", "en", lang, originalTitle, "text/plain", "default")
		if err != nil {
			log.Fatal(err)
		}

		var pages []string

		if pageNumber == 1 {
			pages = []string{fileName}
		} else {
			if _, ok := existingMetadata.Languages[lang]; !ok {
				log.Fatal("Couldn't find existing metadata for language: " + lang)
			}
			if len(existingMetadata.Languages[lang].Pages) != 0 {
				log.Fatal("Invalid existing metadata for language: " + lang)
			}
			pages = append(existingMetadata.Languages[lang].Pages, fileName)
		}

		metadata.Languages[lang] = articles.ArticleLanguageMetadata{
			Title: strings.Trim(translatedTitle, "\n"),
			Pages: pages,
		}
	}

	metadataJSON, err := json.MarshalIndent(&metadata, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(metadataFilePath, metadataJSON, 0644)

	copyFile(inputFileMD, path.Join(outPath, idStr+"__original.md"))
	copyFile(inputFileHTML, path.Join(outPath, idStr+"__original.html"))
}

func translateAndWrite(outPath string, lang string, html string, id string, pageNumber int) (string, error) {
	translatedHTML, err := translateText(os.Stdout, "planetocd", "en", lang, html, "text/html", "default")
	if err != nil {
		log.Fatal(err)
	}

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(translatedHTML)
	if err != nil {
		log.Fatal(err)
	}
	fileName := id + "_" + lang + "_0" + strconv.Itoa(pageNumber) + ".md"
	ioutil.WriteFile(path.Join(outPath, fileName), []byte(markdown), 0644)
	return fileName, nil
}

func copyFile(src string, dest string) {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = ioutil.WriteFile(dest, input, 0644)
	if err != nil {
		fmt.Println("Error creating", dest)
		log.Fatal(err)
		return
	}
}

func translateText(w io.Writer, projectID string, sourceLang string, targetLang string, text string, mimeType string, modelType string) (string, error) {
	ctx := context.Background()
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return "", fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer client.Close()

	model := ""
	if modelType != "default" {
		model = "projects/planetocd/locations/global/models/general/" + modelType
	}

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", projectID),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		Model:              model,    // nmt or base
		MimeType:           mimeType, // Mime types: "text/plain", "text/html"
		Contents:           []string{text},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return "", fmt.Errorf("TranslateText: %v", err)
	}

	res := ""
	for _, translation := range resp.GetTranslations() {
		res = res + translation.GetTranslatedText() + "\n"
	}

	return res, nil
}
