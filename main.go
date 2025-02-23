package main

import (
	"scraper/ekstraklasa/scraper"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/table", func(c *gin.Context) {
		standings := scraper.Scrape()

		c.JSON(200, gin.H{
			"standings": standings,
		})
	})
	r.Run()
}
