package main

import (
	"fmt"
	"github.com/zjzjzjzj1874/zoo/internal/zoo"
	"log"
	"net/http"
)

func main() {
	engine := zoo.New()
	engine.GET("/ping", func(ctx *zoo.Context) {
		ctx.String(http.StatusOK, "%s", "pong")
	})

	engine.GET("/hello", func(ctx *zoo.Context) {
		ctx.JSON(http.StatusOK, zoo.H{
			"hello": "world",
		})
	})

	engine.POST("/login", func(ctx *zoo.Context) {
		ctx.JSON(http.StatusOK, zoo.H{
			"username": ctx.PostForm("username"),
			"password": ctx.PostForm("password"),
		})
	})

	fmt.Println("running on 8082...")
	log.Fatal(engine.Run(":8082"))
}
