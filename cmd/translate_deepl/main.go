package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aureliengasser/planetocd/translate/gateway/deepl"
	"github.com/aureliengasser/planetocd/utils"
	"github.com/urfave/cli/v2"
)

var DEFAULT_DEEPL_TOKEN_PATH = os.Getenv("PLANETOCD_DEEPL_TOKEN_PATH")

func main() {
	var token string
	var targetLanguage string
	var inputExtension string

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
				Value:       ".txt",
			},
		},
		Action: func(c *cli.Context) error {

			translated, err := TranslateFile(
				c.Args().Get(0),
				inputExtension,
				token,
				targetLanguage)

			if err != nil {
				return err
			}

			fmt.Print(translated)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func TranslateFile(inputFile string, inputExtension string, token string, targetLanguage string) (string, error) {

	inputText, err := utils.GetInputText(inputFile)

	if err != nil {
		return "", err
	}

	return translateDeepl(inputText, inputExtension, token, targetLanguage)
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
