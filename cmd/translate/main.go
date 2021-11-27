package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aureliengasser/planetocd/translate/gateway"
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
				Name:        "token",
				Usage:       "Access token",
				Destination: &token,
				Required:    false,
			},
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
			},
		},
		Action: func(c *cli.Context) error {
			TranslateFile(
				c.Args().Get(0),
				inputExtension,
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

func TranslateFile(inputFile string, inputExtension string, token string, targetLanguage string) {

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

	text, err := gateway.Translate(
		inputText,
		inputExtension,
		strings.ToUpper(targetLanguage),
		token,
		gateway.FORMALITY_MORE,
	)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(text)
}
