package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nafima/little-things/shortener"
	"github.com/nafima/little-things/store"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
}

type DataStructure struct {
	Url        string `json:"url"`
	Counter    int32  `json:"counter"`
	Created_at string `json:"created_at"`
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	current := time.Now()
	current.Format("2006-01-02 15:04:05")
	currentDate := current.String()[:19]

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, currentDate)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, currentDate)

	host := "http://localhost/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	var decodedData DataStructure
	json.Unmarshal([]byte(initialUrl), &decodedData)
	decodedData.Counter = decodedData.Counter + 1
	store.UpdateUrlMapping(shortUrl, decodedData.Url, decodedData.Created_at, decodedData.Counter)
	c.Redirect(302, decodedData.Url)
}
