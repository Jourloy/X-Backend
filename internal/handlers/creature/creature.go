package creature_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/guards"
	"github.com/jourloy/X-Backend/internal/modules/creature"
)

func Init(r *gin.Engine) {
	gMain := r.Group(`creature`)
	initCreature(gMain)

}

func initCreature(g *gin.RouterGroup) {
	controller := creature.Init()

	g.Use(guards.CheckAPI())

	g.POST(`one`, controller.GetOne)
	g.POST(`all`, controller.GetAll)

	g.Use(guards.CheckAdmin())

	g.POST(``, controller.Create)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
