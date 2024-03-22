package building_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/guards"
	"github.com/jourloy/X-Backend/internal/modules/building"
)

func Init(r *gin.Engine) {
	gMain := r.Group(`building`)
	initBuilding(gMain)
}

func initBuilding(g *gin.RouterGroup) {
	controller := building.Init()

	g.Use(guards.CheckAPI())

	g.POST(`one`, controller.GetOne)
	g.POST(`all`, controller.GetAll)

	g.Use(guards.CheckAdmin())

	g.POST(``, controller.Create)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
