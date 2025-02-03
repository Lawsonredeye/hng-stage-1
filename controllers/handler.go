package controllers

import (
	"fmt"
	"io"
	"net/http"
	"stage-two/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ClassifyNumber(c *gin.Context) {
	answer, ok := c.GetQuery("number")
	if !ok {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	number, err := strconv.Atoi(answer)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"number": "alphabet",
			"error": "true",
		})
		return
	}
	URL := fmt.Sprintf("http://numbersapi.com/%d/math", number)
	res, err := http.Get(URL)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	sum, _ := utils.DigitSum(int64(number))

	c.JSON(200, gin.H{
		"number": number,
		"is_prime": utils.IsPrime(int64(number)),
		"is_perfect": utils.IsPerfect(int64(number)),
		"properties": utils.Properties(int64(number)),
		"digit_sum": sum,
		"fun_fact": string(body),
	})
}
