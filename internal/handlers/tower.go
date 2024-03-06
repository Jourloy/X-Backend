package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jourloy/X-Backend/internal/modules/tower"
)

func InitTower(g *gin.RouterGroup) {

	controller := tower.Init()

	g.POST(``, controller.Create)
	g.GET(``, controller.GetOne)
	g.GET(`all`, controller.GetAll)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
