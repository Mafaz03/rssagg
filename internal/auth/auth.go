package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetApiKey(header http.Header) (string, error) {
	val := header.Get("Auth")
	if val == "" {
		return "", fmt.Errorf("no API header Found, contained: %v", header)
	}
	vals := strings.Split(val, " ")

	if len(vals[0]) != 64 {
		return "", fmt.Errorf("somethings wrong witht the api key: %v", vals[0])
	}
	return vals[0], nil
}
