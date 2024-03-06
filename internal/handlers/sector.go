package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jourloy/X-Backend/internal/modules/sector"
)

func InitSector(g *gin.RouterGroup) {

	controller := sector.Init()

	g.POST(``, controller.Create)
	g.GET(``, controller.GetOne)
	g.GET(`all`, controller.GetAll)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
