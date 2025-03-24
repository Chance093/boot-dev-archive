package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("Authorization header missing")
	}

	if !strings.Contains(authHeader, "ApiKey ") {
		return "", errors.New("Authorization header missing ApiKey")
	}

	apiKey := strings.TrimPrefix(authHeader, "ApiKey ")

	return apiKey, nil
}
