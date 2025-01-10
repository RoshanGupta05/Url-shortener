package routes

import (
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to the URL Shortener API",
			"endpoints": []string{
				"POST /shorten - To shorten a URL",
				"GET /:shortURL - To redirect to the original URL",
				"GET /metrics - To get top domains shortened",
			},
		})
	})

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})

	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:shortURL", handlers.RedirectURL)
	r.GET("/metrics", handlers.GetTopDomains)
}
