package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jourloy/X-Backend/internal/modules/trader"
)

func InitTrader(g *gin.RouterGroup) {

	controller := trader.Init()

	g.POST(``, controller.Create)
	g.GET(``, controller.GetOne)
	g.GET(`all`, controller.GetAll)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
