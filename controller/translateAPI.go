package controller

import (
	"net/http"

	"github.com/ShevonKuan/translate-server/module"
	"github.com/gin-gonic/gin"
)

func TranslateAPI(c *gin.Context) {
	reqj := module.InputObj{}
	requestMethod := c.Request.Method
	var responseHandler func(int, interface{})
	if requestMethod == "GET" {
		responseHandler = c.JSONP // JSONP use param "callback"
	} else if requestMethod == "POST" {
		responseHandler = c.JSON
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"code":    http.StatusMethodNotAllowed,
			"message": "Method Not Allowed",
		})
		return
	}
	c.BindJSON(&reqj)
	translateEngine, err := module.GetEngine(reqj.Engine)
	if err != nil {
		responseHandler(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
		return
	}
	res, statusCode, err := module.Engine[translateEngine](&reqj)

	if err != nil || statusCode != http.StatusOK {
		responseHandler(http.StatusServiceUnavailable, gin.H{
			"code":    statusCode,
			"message": err,
		})
		return

	}
	responseHandler(http.StatusOK, gin.H{
		"code":         http.StatusOK,
		"data":         res.TransText,
		"alternatives": res.Alternatives,
		"engine":       translateEngine,
	})
	return
}
