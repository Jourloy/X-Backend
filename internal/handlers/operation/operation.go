package operation_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/guards"
	"github.com/jourloy/X-Backend/internal/modules/operation"
)

func Init(r *gin.Engine) {
	gMain := r.Group(`operation`)
	initOperation(gMain)
}

func initOperation(g *gin.RouterGroup) {
	controller := operation.Init()

	g.Use(guards.CheckAPI())

	g.POST(`one`, controller.GetOne)
	g.POST(`all`, controller.GetAll)

	g.POST(``, controller.Create)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
