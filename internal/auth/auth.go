package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract api key from header in http request
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed part of auth header")
	}

	return vals[1], nil
}
