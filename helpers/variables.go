package helpers

import "errors"

var (
	ErrInvalidURL        = errors.New("invalid URL")
	ErrURLNotAllowed     = errors.New("URL is not allowed")
	ErrInvalidAPILimit   = errors.New("invalid API Quota")
	ErrCannotConnect     = errors.New("cannot connect to server")
	ErrInvalidRateLimit  = errors.New("invalid rate limit")
	ErrRateLimitExceeded = "rate limit exceeded, try again in %v"
	ErrFailedToDecrement = "failed to decrement quota: %s"
)
