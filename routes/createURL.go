package routes

import (
	"errors"
	"net/http"
	"time"

	"github.com/RLungWu/MyShortURL/db"
	"github.com/RLungWu/MyShortURL/internal/helpers"
	"github.com/asaskevich/govalidator"
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

func checkURL(url string) (string, error) {
	//Check if the URL is valid
	if !govalidator.IsURL(url) {
		return url, errors.New("Invalid URL")
	}

	//check for domain error
	if !helpers.IsDomainValid(url) {
		return url, errors.New("URL is not allowed")
	}

	//enforce https
	url = helpers.EnforceHTTP(url)
	return url, nil
}

func CreateShortURL(c *gin.Context) {
	// CreateShortURL is the handler for POST /api/v1
	// It creates a new short URL
	body := new(request)
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	checkedURL, err := checkURL(body.URL)
	if err != nil {
		status := http.StatusBadRequest
		if err.Error() == "URL is not allowed" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	body.URL = checkedURL

	var id string
	if body.Custom != "" {
		id = uuid.New().String()[:6]
	} else {
		id = body.Custom
	}

	r := db.CreateClient(0)
	defer r.Close()

	val, _ = r.Get(db.Context, id).Result()
	if val != ""{
		c.JSON(http.StatusForbidden, gin.H{"error": "Custom URL already exists!"})
	}

}
