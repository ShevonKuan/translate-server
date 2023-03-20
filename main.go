package main

import (
	"flag"
	"fmt"

	"github.com/ShevonKuan/translate-server/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "1188", "Translate server port. Default: 1188")
	ip := flag.String("ip", "0.0.0.0", "Translate server IP. Default: 0.0.0.0")
	flag.Parse()
	// display information
	fmt.Println("Translate Server has been successfully launched! Listening on " + *ip + ":" + *port)
	fmt.Println("😀 Made by Shevon")

	// set release mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", controller.StatusCode200)
	// A route that will be called when a GET request is made to the `/rss` endpoint.
	// Query: engine=deepl, url=xxx
	r.GET("/rss", controller.TranslateRSS)
	// A route that will be called when a POST request is made to the `/translate` endpoint.
	// Query: engine=deepl
	r.POST("/translate", controller.TranslateAPI)
	r.Run(*ip + ":" + *port) // listen and serve on 0.0.0.0:1188
}
