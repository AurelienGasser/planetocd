package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aureliengasser/planetocd/translate/gateway/google"
	"github.com/aureliengasser/planetocd/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	var srcLanguageCode string
	var targetLanguageCode string
	var mimeType string
	var projectID string
	var credentialsFile string
	var modelType string

	app := &cli.App{
		Usage: "Translate a file",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "src-lang",
				Usage:       "Source language code",
				Destination: &srcLanguageCode,
				Required:    false,
				Value:       "en",
			},
			&cli.StringFlag{
				Name:        "lang",
				Usage:       "Target language code",
				Destination: &targetLanguageCode,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "project-id",
				Usage:       "Google Project ID",
				Destination: &projectID,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "credentials-file",
				Usage:       "Google Application Credentials file",
				Destination: &credentialsFile,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "model-type",
				Usage:       "Google Translate Model type",
				Destination: &modelType,
				Required:    false,
				Value:       "default",
			},
			&cli.StringFlag{
				Name:        "mime-type",
				Usage:       "Input Mime type",
				Destination: &mimeType,
				Required:    false,
			},
		},
		Action: func(c *cli.Context) error {
			translated, err := TranslateFile(
				c.Args().Get(0),
				srcLanguageCode,
				targetLanguageCode,
				projectID,
				credentialsFile,
				modelType,
				mimeType)

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

func TranslateFile(
	inputFile string,
	srcLanguageCode string,
	targetLanguageCode string,
	projectID string,
	credentialsFile string,
	modelType string,
	mimeType string) (string, error) {

	inputText, err := utils.GetInputText(inputFile)

	if err != nil {
		return "", err
	}

	return translateGoogle(inputText, srcLanguageCode, targetLanguageCode, projectID, credentialsFile, modelType, mimeType)
}

func translateGoogle(
	inputText string,
	sourceLang string,
	targetLang string,
	projectID string,
	credentialsFile string,
	modelType string,
	mimeType string) (string, error) {

	config := google.NewConfig(projectID, credentialsFile, modelType)
	return google.Translate(config, sourceLang, targetLang, inputText, mimeType)
}
