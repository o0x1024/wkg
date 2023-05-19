package toolService

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
	"wkg/model/request"
)

type ChatGPTResp struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTStruct struct {
	Model     string        `json:"model"`
	Messages  []ChatMessage `json:"messages"`
	MaxTokens int           `json:"max_tokens"`
}

func SendQuestion(gpt *request.ChatGPT) (string, error) {
	if gpt.SecretKey != "gelen" {
		return "", errors.New("secret key error")
	}
	uri := "https://openai.proxy.mailjob.net/v1/chat/completions"

	chatMags := []ChatMessage{}
	chatMsg := ChatMessage{Role: "user", Content: gpt.Question}
	chatMags = append(chatMags, chatMsg)

	chatData := ChatGPTStruct{
		Model:     "gpt-3.5-turbo",
		MaxTokens: 2000,
		Messages:  chatMags,
	}

	chatByte, _ := json.Marshal(&chatData)

	req, err := http.NewRequest("POST", uri, bytes.NewReader(chatByte))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer "+gpt.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	// proxy, _ := url.Parse("http://127.0.0.1:8088")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		// Proxy:               http.ProxyURL(proxy),
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 50,
	}

	client := &http.Client{Transport: tr, Timeout: time.Duration(120 * time.Second)}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	returnMessage := ""
	chatResp := &ChatGPTResp{}
	if resp.StatusCode == 200 {
		restr, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}

		err = json.Unmarshal(restr, chatResp)
		if err != nil {
			return "", err
		}
		returnMessage = chatResp.Choices[0].Message.Content
	}
	defer resp.Body.Close()
	return returnMessage, nil
}
