package faq

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func askFlaskServer(question string) (string, error) {
	client := &http.Client{}

	reqBody := []byte(fmt.Sprintf(`{"question": "%s"}`, question))
	req, err := http.NewRequest("POST", "http://127.0.0.1:5000/api", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("Failed to create request: %s", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Failed to send request: %s", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var jsonResponse map[string]string
	err = json.Unmarshal(body, &jsonResponse)
	if err != nil {
		return "", fmt.Errorf("Failed to parse JSON response: %s", err)
	}

	return jsonResponse["answer"], nil
}
