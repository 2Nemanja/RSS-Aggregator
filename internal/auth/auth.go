package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts API Key from the Headers of an HTTP request
// Example:
// Authorization: APIKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication ifo found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the auth header")
	}
	return vals[1], nil

}
