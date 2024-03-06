package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jourloy/X-Backend/internal/modules/warrior"
)

func InitWarrior(g *gin.RouterGroup) {

	controller := warrior.Init()

	g.POST(``, controller.Create)
	g.GET(``, controller.GetOne)
	g.GET(`all`, controller.GetAll)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
