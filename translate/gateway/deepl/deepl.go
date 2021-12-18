package deepl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	DOCUMENT_UPLOAD_URL = "https://api-free.deepl.com/v2/document"
	DOCUMENT_STATUS_URL = "https://api-free.deepl.com/v2/document/%v"
	DOCUMENT_RESULT_URL = "https://api-free.deepl.com/v2/document/%v/result"
	FORMALITY_DEFAULT   = "default"
	FORMALITY_MORE      = "more"
	FORMALITY_LESS      = "less"
)

func Translate(inputText string, inputExtension string, targetLang string, authKey string, formality string) (string, error) {
	client := &http.Client{}

	doc, err := uploadDocument(client, inputText, inputExtension, targetLang, authKey, formality)
	if err != nil {
		return "", fmt.Errorf("error while uploading to Deepl %v in %v: %w", inputExtension, targetLang, err)
	}

	err = waitForTranslation(client, authKey, doc)

	if err != nil {
		return "", fmt.Errorf("error while translating %v in %v (%v): %w", inputExtension, targetLang, doc.DocumentID, err)
	}

	translatedText, err := getDocumentResult(client, authKey, doc)

	if err != nil {
		return "", fmt.Errorf("error while fetching the translation for %v in %v (%v): %w", inputExtension, targetLang, doc.DocumentID, err)
	}

	return translatedText, nil
}

func getDocumentResult(client *http.Client, authKey string, doc *DeeplDocument) (string, error) {
	targetURL := fmt.Sprintf(DOCUMENT_RESULT_URL, doc.DocumentID)

	form := url.Values{}
	form.Add("auth_key", authKey)
	form.Add("document_key", doc.DocumentKey)

	req, err := http.NewRequest("POST", targetURL, strings.NewReader(form.Encode()))

	if err != nil {
		return "", err
	}

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error while uploading the document. Status code: %v, Response: %v", resp.Status, string(content))
	}

	return string(content), nil
}

func uploadDocument(
	client *http.Client,
	inputText string,
	inputExtension string,
	targetLang string,
	authKey string,
	formality string,
) (*DeeplDocument, error) {
	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	fileWriter, err := writer.CreateFormFile("file", fmt.Sprintf("input.%v", inputExtension))
	if err != nil {
		return nil, err
	}
	fileWriter.Write([]byte(inputText))
	writer.WriteField("source_lang", "EN")
	writer.WriteField("target_lang", targetLang)
	writer.WriteField("auth_key", authKey)
	writer.WriteField("formality", formality)
	writer.Close()

	r, _ := http.NewRequest("POST", DOCUMENT_UPLOAD_URL, &body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error while uploading the document. Status code: %v, Response: %v", resp.Status, string(content))
	}

	doc := DeeplDocument{}

	err = json.Unmarshal(content, &doc)
	if err != nil {
		return nil, err
	}

	return &doc, nil
}

func waitForTranslation(client *http.Client, authKey string, doc *DeeplDocument) error {
	for {
		status, err := getDocumentStatus(client, authKey, doc)
		if err != nil {
			return err
		}
		if status.Status == "queued" || status.Status == "translating" {
			delay := status.SecondsRemaining
			if delay <= 0 {
				delay = 1
			}
			fmt.Fprintf(os.Stderr, "Deepl document %v, Status: %v. Waiting for %v seconds\n", doc.DocumentID, status.Status, delay)
			time.Sleep(time.Duration(delay) * time.Second)
		} else if status.Status == "done" {
			return nil
		} else if status.Status == "error" {
			return fmt.Errorf("a Deepl error occurred while translating the document")
		} else {
			return fmt.Errorf("error while translating the document. Message = %v", status.Message)
		}
	}
}

func getDocumentStatus(client *http.Client, authKey string, doc *DeeplDocument) (*DeeplDocumentStatus, error) {
	targetURL := fmt.Sprintf(DOCUMENT_STATUS_URL, doc.DocumentID)

	form := url.Values{}
	form.Add("auth_key", authKey)
	form.Add("document_key", doc.DocumentKey)

	req, err := http.NewRequest("POST", targetURL, strings.NewReader(form.Encode()))

	if err != nil {
		return nil, err
	}

	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error while getting document status. Status code: %v, Response: %v", resp.Status, string(content))
	}

	docStatus := DeeplDocumentStatus{}

	err = json.Unmarshal(content, &docStatus)
	if err != nil {
		return nil, err
	}

	return &docStatus, nil
}

type DeeplDocument struct {
	DocumentID  string `json:"document_id"`
	DocumentKey string `json:"document_key"`
	Message     string `json:"message"`
}

type DeeplDocumentStatus struct {
	DocumentID       string `json:"document_id"`
	Status           string `json:"status"`
	SecondsRemaining int    `json:"seconds_remaining"`
	BilledCharacters int    `json:"billed_characters"`
	Message          string `json:"message"`
}
