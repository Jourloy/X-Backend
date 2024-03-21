package handlers

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/guards"
	"github.com/jourloy/X-Backend/internal/modules/creature"
)

func InitCreature(r *gin.Engine) {
	g := r.Group(`creature`)

	controller := creature.Init()

	g.Use(guards.CheckAPI())

	g.POST(`one`, controller.GetOne)
	g.POST(`all`, controller.GetAll)

	g.POST(``, controller.Create)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
