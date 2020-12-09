package controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

// ControllerGolang
type ControllerGolang interface {
	AddGolang(ctx *gin.Context) (int, error)
	GetAllGolang() []entity.Golang
}

type controller struct {
	service service.ServiceGolang
}

func New(services service.ServiceGolang) ControllerGolang {
	return &controller{
		service: services,
	}
}

// AddGolang

func (g *controller) AddGolang(ctx *gin.Context) (int, error) {
	var golangval entity.Golang
	err := ctx.ShouldBindJSON(&golangval)
	if err != nil {
		return 0, err
	}
	id := g.service.AddGolang(golangval)
	return id, nil
}

func (g *controller) GetAllGolang() []entity.Golang {
	return g.service.GetAllGolang()
}
