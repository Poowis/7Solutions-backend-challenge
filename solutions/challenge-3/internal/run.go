package internal

import (
	"io"
	"net/http"
	"poowis/challenge-3/internal/utils"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()
	router.GET("/beef/summary", getBeefSummary)
	router.Run("localhost:8080")
}

func getBeefSummary(c *gin.Context) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to get data"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "unable to read data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"beef": utils.Summary(string(body)),
	})
}
