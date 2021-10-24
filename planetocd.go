package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/aureliengasser/planetocd/server"
)

func main() {
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
