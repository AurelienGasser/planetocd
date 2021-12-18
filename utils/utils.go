package utils

import (
	"bufio"
	"io"
	"os"
)

// GetInputText reads the input from stdin or a file.
func GetInputText(inputFile string) (string, error) {
	if inputFile == "-" {
		// stdin
		reader := bufio.NewReader(os.Stdin)
		inputB, _ := io.ReadAll(reader)
		return string(inputB), nil
	}

	file, err := os.Open(inputFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	inputB, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(inputB), nil
}
