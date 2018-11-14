package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Router method defines the routing behaviors
// And generate a router for others to use
func Router() *gin.Engine {
	// Generate a default router to configure
	router := gin.Default()

	// Set the handler function for paths
	router.GET("/ginTest", func(c *gin.Context) {
		// Use JSON as the response
		c.JSON(http.StatusOK, gin.H{
			"message": "You've successfully received a message from a gin server.",
		})
	})

	// Configuration done, router returned
	return router
}
