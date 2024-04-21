package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *gin.Engine) {
	app.POST("/api/v1", routes.CreateShortURL)
	app.GET("/:shortURL", routes.ResolveURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	app := gin.Default()
	app.Use(gin.Logger())

	setUpRoutes(app)
	log.Fatal(app.RunListener(os.Getenv("PORT")))
}
