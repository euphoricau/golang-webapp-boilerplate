package main

import (
	"webapp/db"
	"webapp/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Set up Gin router
	router := gin.Default()

	// Set up session store
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Load HTML templates
	router.LoadHTMLGlob("templates/*")

	// Serve static files
	router.Static("/static", "./static")

	// Initialize routes
	routes.InitRoutes(router)

	// Run the server
	router.Run(":8080")
}
