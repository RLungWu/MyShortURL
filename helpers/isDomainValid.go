package helpers

import (
	"strings"
)

func IsDomainValid(url string) bool {
	// IsDomainValid checks if the domain is valid
	// It returns true if the domain is allowed
	// and false if the domain is not allowed
	domain := "localhost:3000"
	if url == domain {
		return false
	}

	trimmedURL := strings.TrimPrefix(strings.TrimPrefix(strings.TrimPrefix(url, "http://"), "https://"), "www.")
	hostName := strings.Split(trimmedURL, "/")[0]

	return hostName != domain
}
