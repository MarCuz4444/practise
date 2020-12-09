package main

import (
	"io"
	"net/http"

	"os"

	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

var (
	servicegolang    service.ServiceGolang       = service.New()
	controllergolang controller.ControllerGolang = controller.New(servicegolang)
)

func setupOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	server := gin.New()
	setupOutput()

	server.GET("/golang", func(g *gin.Context) {
		g.JSON(200, controllergolang.GetAllGolang())
	})

	server.POST("/golang", func(g *gin.Context) {
		id, err := controllergolang.AddGolang(g)
		if err != nil {
			g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			g.JSON(200, gin.H{
				"id": id,
			})
		}
	})
	server.Run()
}
