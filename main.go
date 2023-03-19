package main

import (
	"flag"
	"fmt"

	"github.com/ShevonKuan/translate-server/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", "1188", "DeepL server port. Default: 1188")
	ip := flag.String("ip", "0.0.0.0", "DeepL server IP. Default: 0.0.0.0")
	flag.Parse()
	// display information
	fmt.Println("DeepL Server has been successfully launched! Listening on " + *ip + ":" + *port)
	fmt.Println("😀 Made by Shevon. Designed by sjlleo and missuo.")

	// create a random id
	id := controller.GetRandomNumber()

	// set release mode
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(func(context *gin.Context) {
		context.Set("id", id)
	})
	r.GET("/", controller.StatusCode200)
	r.POST("/translate", controller.TranslateAPI)
	r.Run(*ip + ":" + *port) // listen and serve on 0.0.0.0:1188
}
