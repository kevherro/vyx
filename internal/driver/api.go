package driver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/kevherro/vyx/internal/api/completions"
)

const (
	url   = "www.openai.com"
	model = "model"
)

func parseTokens(input []string) ([]string, error) {
	prompt := strings.Join(input, " ")
	cfg := currentConfig()
	payload := &completions.Request{
		Prompt:      prompt,
		Model:       model,
		MaxTokens:   cfg.MaxTokens,
		Temperature: cfg.Temperature,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var completionResponse completions.Response
	err = json.Unmarshal(body, &completionResponse)
	if err != nil {
		return nil, err
	}

	if len(completionResponse.Choices) == 0 {
		if os.Getenv("OPENAI_API_KEY") == "" {
			return strings.Fields("vyx: missing OPENAI_API_KEY"), nil
		}
		return strings.Fields("vyx: unable to generate a response"), nil
	}
	choice := completionResponse.Choices[0]
	text := choice.Text

	return strings.Fields(text), nil
}
