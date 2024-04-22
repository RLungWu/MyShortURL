package main

import (
	"fmt"
	"log"

	"github.com/RLungWu/MyShortURL/routes"

	"github.com/gin-gonic/gin"
)

func setUpRoutes(app *gin.Engine) {
	app.POST("/api/v1/", routes.CreateShortURL)
	app.GET("/:shortURL", routes.ResolveURL)
}

func main() {
	app := gin.Default()
	app.Use(gin.Logger())

	setUpRoutes(app)

	fmt.Println("Server is running on port 3000")

	log.Fatal(app.Run(":3000"))

}
