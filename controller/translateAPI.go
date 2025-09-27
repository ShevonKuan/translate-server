package controller

import (
	"net/http"
	"time"

	"github.com/ShevonKuan/translate-server/module"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func TranslateAPI(c *gin.Context) {
	reqj := module.InputObj{}
	requestMethod := c.Request.Method
	var responseHandler func(int, interface{})

	if err := c.ShouldBindQuery(&reqj); err != nil {
		if err := c.ShouldBindJSON(&reqj); err != nil {
			responseHandler(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": err,
			})
			return
		}
		responseHandler(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err,
		})
		return
	}
	color.Blue("Receive translate request" + time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).String())
	translateEngine, err := module.GetEngine(reqj.Engine)
	if reqj.Engine == "llm" {
		_, _, err := module.LLMTranslate(&reqj, c) // need to handle streaming in the function
		if err != nil {
			responseHandler(http.StatusServiceUnavailable, gin.H{
				"code":    500,
				"message": err,
			})
			return
		}

		return

	} else {

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
}
