package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetApiKeySuccess(t *testing.T) {
	testKey := "1a2b3c4d5"
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("ApiKey %v", testKey))

	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("Error getting key: %v", err)
	}
	if key != testKey {
		t.Fatalf("Keys don't match; expected %v, got %v", testKey, key)
	}
}

func TestGetApiKeyEmpty(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("GetApiKey didn't throw error on empty headers list")
	}
}

func TestGetApiKeyMalformed(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "12345")

	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatalf("GetApiKey didn't throw error on malformed authorization headers")
	}

	headers.Set("Authorization", "ApiKey")
	_, err = GetAPIKey(headers)
	if err == nil {
		t.Fatalf("GetApiKey didn't throw error when missing api key")
	}
}
