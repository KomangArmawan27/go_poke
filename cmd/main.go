package main

import (
	"fmt"
	"go-api/config"
	"go-api/internal/routes"
)

// @title           PokeAPI
// @version         1.0

// @contact.name    Komang Damai
// @contact.email   komangdamai3@gmail.com

// @license.name    MIT License
// @license.url     https://opensource.org/licenses/MIT

// @host            localhost:8080
// @BasePath        /api/v1

// @securityDefinitions.apikey BearerAuth
// @type            apiKey
// @in              header
// @name            Authorization
// @description     Type "Bearer" followed by a space and your token.

func main() {
	config.LoadEnv()
	config.ConnectDatabase()

	// Start the Gin server
	r := routes.SetupRoutes()
	port := config.GetEnv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	fmt.Println("🚀 Server running on port", port)
	r.Run(":" + port)
}
