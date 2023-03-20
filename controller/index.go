package controller

import "github.com/gin-gonic/gin"

func StatusCode200(c *gin.Context) { // 200 OK
	c.JSON(200, gin.H{
		"code":    200,
		"message": "Tranlsate API",
	})

}
