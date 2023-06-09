package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// HTTPClient é uma struct que encapsula um cliente HTTP
type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient cria e retorna uma nova instância de HTTPClient
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
	}
}

// Post faz uma requisição HTTP POST para a URL fornecida com o payload fornecido
// e retorna a resposta como uma string
func (c *HTTPClient) Post(url string, payload interface{}) (string, error) {
	// Converte o payload para JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Cria uma nova requisição HTTP
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	// Define o tipo de conteúdo para JSON
	req.Header.Set("Content-Type", "application/json")

	// Envia a requisição HTTP
	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Lê a resposta
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
