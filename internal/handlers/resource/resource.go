package resource_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/guards"
	"github.com/jourloy/X-Backend/internal/modules/resource"
)

func Init(r *gin.Engine) {
	gMain := r.Group(`resource`)
	initResource(gMain)
}

func initResource(g *gin.RouterGroup) {
	controller := resource.Init()

	g.Use(guards.CheckAPI())

	g.Use(guards.CheckAdmin())

	g.POST(``, controller.Create)
}
