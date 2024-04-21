package main

import (
	"fmt"
	"log"
	"os"

	"github.com/RLungWu/MyShortURL/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setUpRoutes(app *gin.Engine) {
	app.POST("/api/v1", routes.CreateShortURL())
	app.GET("/:shortURL", routes.ResolveURL())
}

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	app := gin.Default()
	app.Use(gin.Logger())

	setUpRoutes(app)

	port := os.Getenv("PORT")
	if port == ""{
		log.Fatal("PORT must be set in .env file")
	}

	log.Fatal(app.Run(":" + port))

}
