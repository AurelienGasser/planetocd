package main

import (
	"io/ioutil"
	"log"
	"os"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/urfave/cli/v2"
)

var DEFAULT_INPUT_MD_FILE = "./workdir/in.md"
var DEFAULT_INPUT_HTML_FILE = "./workdir/in.html"

func main() {

	var in_file string
	var out_file string

	app := &cli.App{
		Usage: "Convert a HTML file to Markdown",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input-file",
				Value:       DEFAULT_INPUT_HTML_FILE,
				Usage:       "path to input HTML file to convert",
				Destination: &in_file,
			},
			&cli.StringFlag{
				Name:        "output-file",
				Value:       DEFAULT_INPUT_MD_FILE,
				Usage:       "path to ouput Markdown file",
				Destination: &out_file,
			},
		},
		Action: func(c *cli.Context) error {
			HTMLToMarkdown(in_file, out_file)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// HTMLToMarkdown ...
func HTMLToMarkdown(in_file string, out_file string) {
	html, err := ioutil.ReadFile(in_file)
	if err != nil {
		log.Fatal(err)
	}
	converter := md.NewConverter("", true, nil)

	markdown, err := converter.ConvertString(string(html))
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(out_file, []byte(markdown), 0644)
}
