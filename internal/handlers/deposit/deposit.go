package deposit_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/guards"
	"github.com/jourloy/X-Backend/internal/modules/deposit"
)

func Init(r *gin.Engine) {
	gMain := r.Group(`deposit`)
	initDeposit(gMain)
}

func initDeposit(g *gin.RouterGroup) {
	controller := deposit.Init()

	g.Use(guards.CheckAPI())

	g.Use(guards.CheckAdmin())

	g.POST(``, controller.Create)
}
