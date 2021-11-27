package google

import (
	"context"
	"fmt"
	"os"

	translate "cloud.google.com/go/translate/apiv3"
	"google.golang.org/api/option"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

type GoogleConfig struct {
	Model           string
	Parent          string
	CredentialsFile string
}

func NewConfig(projectID string, credentialsFile string, modelType string) *GoogleConfig {
	if projectID == "" {
		projectID = os.Getenv("PLANETOCD_GOOGLE_PROJECT_ID")
	}

	if credentialsFile == "" {
		credentialsFile = os.Getenv("PLANETOCD_GOOGLE_APPLICATION_CREDENTIALS")
	}

	model := ""
	if modelType != "default" {
		model = "projects/planetocd/locations/global/models/general/" + modelType
	}

	return &GoogleConfig{
		Model:           model,
		Parent:          fmt.Sprintf("projects/%s/locations/global", projectID),
		CredentialsFile: credentialsFile,
	}
}

func Translate(config *GoogleConfig, sourceLang string, targetLang string, text string, mimeType string) (string, error) {
	ctx := context.Background()

	client, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile(config.CredentialsFile))
	if err != nil {
		return "", fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		Parent:             config.Parent,
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		Model:              config.Model,
		MimeType:           mimeType, // Mime types: "text/plain", "text/html"
		Contents:           []string{text},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return "", fmt.Errorf("TranslateText: %v", err)
	}

	res := ""
	for _, translation := range resp.GetTranslations() {
		res = res + translation.GetTranslatedText() + "\n"
	}

	return res, nil
}
