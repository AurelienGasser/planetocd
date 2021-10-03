package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"

	"github.com/aureliengasser/planetocd/server"
	"github.com/aureliengasser/planetocd/translate"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "action",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			action := c.String("action")
			switch action {
			case "server":
				startServer()
			case "convert":
				convert()
			case "translate":
				// GOOGLE_APPLICATION_CREDENTIALS=~/.local/planetocd-86fb09efe9c9.json
				doTranslate()
			default:
				log.Fatalf("Unknown action %v", action)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func startServer() {
	isLocal := isLocalEnvironment()
	var port int

	if isLocal {
		port = server.DefaultPort
	} else {
		port = getPort()
	}

	fmt.Printf("Starting server on port %v\n", port)
	server.Listen(port, isLocal)
}

func convert() {
	translate.HTMLToMarkdown()
}

func doTranslate() {
	translate.CreateTranslatedArticle(
		"0001",
		"http://ocdla.com/cognitivebehavioraltherapy",
		"",
		"OCD Treatment: Cognitive-Behavioral Therapy",
		3,
	)
}

func isLocalEnvironment() bool {
	environment, ok := os.LookupEnv("ENVIRONMENT")
	if !ok {
		return true
	}

	return environment != "production"
}

func getPort() int {
	portStr, ok := os.LookupEnv("PORT")

	if !ok {
		return server.DefaultPort
	}

	port, err := strconv.Atoi(portStr)

	if err != nil {
		fmt.Printf("Invalid port %s\n", portStr)
		os.Exit(1)
	}

	return port
}
