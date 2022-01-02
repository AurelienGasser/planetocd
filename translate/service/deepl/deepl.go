package deepl

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/aureliengasser/planetocd/translate/gateway/deepl"
)

var DEFAULT_DEEPL_TOKEN_PATH = os.Getenv("PLANETOCD_DEEPL_TOKEN_PATH")

func GetDefaultToken() (string, error) {
	file, err := os.Open(DEFAULT_DEEPL_TOKEN_PATH)
	if err != nil {
		return "", fmt.Errorf("error while opening the token file: %w", err)
	}
	defer file.Close()
	tokenB, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error while reading the token file: %w", err)
	}
	return string(tokenB), nil
}

func GetFormalityForLanguage(lang string) string {
	switch strings.ToLower(lang) {
	case "fr", "es":
		return deepl.FORMALITY_MORE
	}
	return deepl.FORMALITY_DEFAULT
}

func Translate(inputText string, inputExtension string, targetLang string) (string, error) {
	authKey, err := GetDefaultToken()

	if err != nil {
		return "", err
	}

	return deepl.Translate(
		inputText,
		inputExtension,
		strings.ToUpper(targetLang),
		authKey,
		GetFormalityForLanguage(targetLang),
	)
}
