package controller

import (
	"errors"
	"net/http"

	"pencairan_user/server/service"

	"github.com/gin-gonic/gin"
)

func Username(c *gin.Context) {
	var urls []string

	if err := c.ShouldBindJSON(&urls); err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("invalid json body"))
	}

	matchedUrls := service.UsernameService.UsernameCheck(urls)
	c.JSON(http.StatusOK, matchedUrls)
}
