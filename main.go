package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new router instance
	router := gin.Default()

	router.Static("/static", "./static")
	// Load the HTML template
	router.LoadHTMLGlob("static/*")

	// Define a route for the home page
	router.GET("/", func(c *gin.Context) {
		// Render the "index.html" template
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home Page",
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/css")
	})
	// Define a route for the about page
	router.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "This is a simple web page created with Go and Gin!")
	})

	// Start the server
	router.Run(":8080")
}
