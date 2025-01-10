package handlers

import (
	"net/http"
	"url-shortener/storage"
	"url-shortener/utils"

	"github.com/gin-gonic/gin"
)

func ShortenURL(c *gin.Context) {
	var request struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := utils.GenerateShortURL(request.URL)
	originalURL := storage.GetOriginalURL(shortURL)

	if originalURL == "" {
		storage.SaveURLMapping(shortURL, request.URL)
	}

	c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
}

func RedirectURL(c *gin.Context) {
	shortURL := c.Param("shortURL")
	originalURL := storage.GetOriginalURL(shortURL)

	if originalURL == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

func GetTopDomains(c *gin.Context) {
	topDomains := storage.GetTopDomains()
	c.JSON(http.StatusOK, topDomains)
}
