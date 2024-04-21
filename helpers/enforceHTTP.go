package helpers

import "strings"

func EnforceHTTP(url string) string {
	// EnforceHTTP enforces the use of HTTPS
	// It returns the URL with HTTPS enforced
	return strings.Replace(url, "http://", "https://", 1)
}
