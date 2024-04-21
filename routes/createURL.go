package routes

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/RLungWu/MyShortURL/db"
	"github.com/RLungWu/MyShortURL/helpers"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type request struct {
	URL    string        `json:"url" binding:"required"`
	Custom string        `json:"custom"`
	Expiry time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	Custom         string        `json:"custom"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"x-rate-remaining"`
	XRateLimitRest time.Duration `json:"x-rate-limit-rest"`
}

var (
	ErrInvalidURL        = errors.New("invalid URL")
	ErrURLNotAllowed     = errors.New("URL is not allowed")
	ErrInvalidAPILimit   = errors.New("invalid API Quota")
	ErrCannotConnect     = errors.New("cannot connect to server")
	ErrInvalidRateLimit  = errors.New("invalid rate limit")
	ErrRateLimitExceeded = "rate limit exceeded, try again in %v"
	ErrFailedToDecrement = "failed to decrement quota: %s"
)

func CreateShortURL(c *gin.Context) {
	// CreateShortURL is the handler for POST /api/v1
	// It creates a new short URL

	//Load the request body into a struct
	body := new(request)
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//Check if the URL is valid
	checkedURL, err := helpers.CheckURL(body.URL)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, ErrURLNotAllowed) {
			status = http.StatusForbidden
		} else if errors.Is(err, ErrInvalidURL) {
			status = http.StatusBadRequest
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	body.URL = checkedURL

	//Check the rate limit
	err = helpers.CheckRateLimit(c.ClientIP())
	if err != nil {
		status := http.StatusInternalServerError
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	var id string
	if body.Custom == "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.Custom
	}

	r := db.CreateClient(0)
	defer r.Close()

	val, _ := r.Get(db.Context, id).Result()
	if val != "" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Custom URL already exists!"})
		return
	}

	if body.Expiry == 0 {
		body.Expiry = 24
	}

	err = r.Set(db.Context, id, body.URL, body.Expiry*time.Hour).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot connect to database!"})
		return
	}

	resp := response{
		URL:            body.URL,
		Custom:         "",
		Expiry:         body.Expiry,
		XRateRemaining: 10,
		XRateLimitRest: 30 * time.Second,
	}

	r.Decr(db.Context, c.ClientIP())

	val, _ = r.Get(db.Context, c.ClientIP()).Result()
	resp.XRateRemaining, _ = strconv.Atoi(val)

	ttl, _ := r.TTL(db.Context, c.ClientIP()).Result()
	resp.XRateLimitRest = ttl / time.Nanosecond / time.Minute

	resp.Custom = os.Getenv("BASE_URL") + "/" + id


	
	c.JSON(http.StatusOK, resp)
}
