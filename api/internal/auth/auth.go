package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts and API key from the headers of an HTTP request
// Ex.
// Authorization: ApiKey {insert api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication information found")
	}

	// otherwise, a valid value exists

	// split on spaces
	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization header")
	}

	// malformed header
	if vals[0] != "ApiKey" {
		return "", errors.New("first part of authorization header is invalid")
	}

	return vals[1], nil

}
