package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aureliengasser/planetocd/translate/gateway/deepl"
	"github.com/aureliengasser/planetocd/translate/gateway/google"
	"github.com/urfave/cli/v2"
)

var DEFAULT_DEEPL_TOKEN_PATH = os.Getenv("PLANETOCD_DEEPL_TOKEN_PATH")

const (
	ENGINE_DEEPL  = "deepl"
	ENGINE_GOOGLE = "google"
)

func main() {
	var token string
	var targetLanguage string
	var inputExtension string
	var engine string

	app := &cli.App{
		Usage: "Translate a file",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Usage:       "Target language",
				Destination: &targetLanguage,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "engine",
				Usage:       "Translation engine",
				Destination: &engine,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "Access token",
				Destination: &token,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "ext",
				Usage:       "Input string extension (corresponding to MIME type)",
				Destination: &inputExtension,
				Required:    false,
			},
		},
		Action: func(c *cli.Context) error {
			TranslateFile(
				c.Args().Get(0),
				inputExtension,
				engine,
				token,
				targetLanguage)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func TranslateFile(inputFile string, inputExtension string, engine string, token string, targetLanguage string) {

	inputText := ""

	if inputFile == "-" {
		reader := bufio.NewReader(os.Stdin)
		inputB, _ := io.ReadAll(reader)
		inputText = string(inputB)
	} else {
		file, err := os.Open(inputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		inputB, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		inputText = string(inputB)
		if inputExtension == "" {
			inputExtension = filepath.Ext(inputFile)
		}
	}

	var text string
	var err error

	if engine == ENGINE_GOOGLE {
		text, err = translateGoogle(inputText, inputExtension, token, targetLanguage)
	} else if engine == ENGINE_DEEPL {
		text, err = translateDeepl(inputText, inputExtension, token, targetLanguage)
	} else {
		log.Fatalf("Engine must be %v or %v", ENGINE_DEEPL, ENGINE_GOOGLE)
	}

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(text)
}

func translateDeepl(inputText string, inputExtension string, token string, targetLanguage string) (string, error) {

	if token == "" {
		file, err := os.Open(DEFAULT_DEEPL_TOKEN_PATH)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		tokenB, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}
		token = string(tokenB)
	}

	return deepl.Translate(
		inputText,
		inputExtension,
		strings.ToUpper(targetLanguage),
		token,
		deepl.FORMALITY_MORE,
	)
}

func translateGoogle(inputText string, inputExtension string, targetLanguage string) (string, error) {

	return google.Translate(
		inputText,
		inputExtension,
		strings.ToUpper(targetLanguage),
		token,
		deepl.FORMALITY_MORE,
	)
}
