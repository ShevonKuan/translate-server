package controller

import (
	"net/http"

	"github.com/ShevonKuan/translate-server/module"
	"github.com/gin-gonic/gin"
)

func TranslateAPI(c *gin.Context) {
	reqj := module.InputObj{}
	c.BindJSON(&reqj)
	translateEngine, err := module.GetEngine(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
	}
	res, statusCode, err := module.Engine[translateEngine](&reqj)

	if err != nil || statusCode != http.StatusOK {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"code":    statusCode,
			"message": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"data":         res.TransText,
		"alternatives": res.Alternatives,
		"engine":       translateEngine,
	})
}
