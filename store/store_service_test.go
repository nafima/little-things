package store

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

func TestInsertionAndRetrieval(t *testing.T) {
	current := time.Now()
	current.Format("2006-01-02 15:04:05")
	currentDate := current.String()[:19]

	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortURL := "Jsz4k57oAX"

	saved := DataStructure{Url: initialLink, Counter: 0, Created_at: currentDate}

	savedJson, _ := json.Marshal(saved)

	// Persist data mapping
	SaveUrlMapping(shortURL, initialLink, currentDate)

	// Retrieve initial URL
	retrievedUrl := RetrieveInitialUrl(shortURL)
	assert.Equal(t, string(savedJson), retrievedUrl)
}
