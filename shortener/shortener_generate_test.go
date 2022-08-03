package shortener

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShortLinkGenerator(t *testing.T) {
	current := time.Now()
	current.Format("2006-01-02 15:04:05")
	currentDate := current.String()[:19]

	initialLink_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink_1 := GenerateShortLink(initialLink_1, currentDate)

	initialLink_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink_2 := GenerateShortLink(initialLink_2, currentDate)

	initialLink_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3 := GenerateShortLink(initialLink_3, currentDate)

	assert.Equal(t, shortLink_1, shortLink_1)
	assert.Equal(t, shortLink_2, shortLink_2)
	assert.Equal(t, shortLink_3, shortLink_3)
}
