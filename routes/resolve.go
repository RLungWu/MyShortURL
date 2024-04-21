package routes

import (
	"net/http"

	"github.com/RLungWu/MyShortURL/db"
	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

func ResolveURL(c *gin.Context) {
	// ResolveURL is the handler for GET /:shortURL
	// It resolves a short URL to the original URL
	shortURL := c.Param("shorturl")
	
	r := db.CreateClient(0)
	defer r.Close()

	value, err := r.Get(db.Context, shortURL).Result()
	if err != redis.Nil{
		c.JSON(http.StatusNotFound, gin.H{"error" : "URL not found in database!"})
		return
	}else if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Cannot connect to database!"})
	}

	rInr := db.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(db.Context, "counter")

	c.Redirect(http.StatusMovedPermanently, value)
}
