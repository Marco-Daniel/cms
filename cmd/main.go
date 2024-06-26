package main

import (
	"cms/core"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := core.Config{
		DBURI: "",
	}

	server, err := core.CreateServer(cfg)
	if err != nil {
		fmt.Println("it went wrong:", err)
		return
	}

	server.RegiserRoutes([]func(*gin.Engine){
		func(r *gin.Engine) {
			r.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"message": "Welcome the the CMS!"})
			})
		},
	})

	server.AddMiddleware([]gin.HandlerFunc{
		func(ctx *gin.Context) {
			fmt.Println("Request recieved")
			ctx.Next()
		},
	})

	server.Start()
}
