package controller

import "github.com/gin-gonic/gin"

func StatusCode200(c *gin.Context) { // 200 OK
	c.JSON(200, gin.H{
		"code":    200,
		"message": "DeepL Free API. Go to /translate with POST. Inspired by http://github.com/OwO-Network/DeepLX",
	})

}
