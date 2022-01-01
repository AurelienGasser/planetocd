package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aureliengasser/planetocd/translate/service/deepl"
	"github.com/aureliengasser/planetocd/utils"
	"github.com/urfave/cli/v2"
)

func main() {
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

func TranslateFile(inputFile string, inputExtension string, targetLanguage string) (string, error) {

	inputText, err := utils.GetInputText(inputFile)

	if err != nil {
		return "", err
	}

	return deepl.Translate(
		inputText,
		inputExtension,
		targetLanguage,
	)
}
