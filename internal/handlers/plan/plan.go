package plan_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/guards"
	"github.com/jourloy/X-Backend/internal/modules/plan"
)

func Init(r *gin.Engine) {
	gMain := r.Group(`plan`)
	initPlan(gMain)
}

func initPlan(g *gin.RouterGroup) {
	controller := plan.Init()

	g.Use(guards.CheckAPI())

	g.GET(``, controller.GetOne)
	g.GET(`all`, controller.GetAll)

	g.Use(guards.CheckAdmin())

	g.POST(``, controller.Create)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
