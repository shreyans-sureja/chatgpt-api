package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/shreyans-sureja/chatgpt-api/constants"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type ChatgptPayload struct {
	Model            string   `json:"model"`
	Prompt           string   `json:"prompt"`
	Temperature      int      `json:"temperature"`
	MaxTokens        int      `json:"max_tokens"`
	TopP             float64  `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
}

type ChatgptResponsePayload struct {
	Error map[string]interface{} `json:"error"`
}

func ChatgptAPICall(cp ChatgptPayload) (ChatgptResponsePayload, error) {
	var chatgptResponsePayload ChatgptResponsePayload
	ctx := context.Background()
	url := fmt.Sprintf("%s/v1/completions", constants.CHATGPT_BASE_URL)

	reqBody, err := json.Marshal(cp)
	if err != nil {
		return chatgptResponsePayload, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return chatgptResponsePayload, err
	}

	req.Header.Add("Authorization", "Bearer "+os.Getenv("CHATGPT_API_KEY"))
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	serviceResponse, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return chatgptResponsePayload, err
	}
	err = json.Unmarshal(serviceResponse, &chatgptResponsePayload)
	if err != nil {
		return chatgptResponsePayload, err
	}

	return chatgptResponsePayload, nil
}
