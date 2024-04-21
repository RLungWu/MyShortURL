package helpers

import "github.com/asaskevich/govalidator"

func CheckURL(url string) (string, error) {
	if !govalidator.IsURL(url) {
		return url, ErrInvalidURL
	}
	if !IsDomainValid(url) {
		return url, ErrURLNotAllowed
	}
	url = EnforceHTTP(url)
	return url, nil
}
