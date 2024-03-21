package account_handler

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/guards"
	"github.com/jourloy/X-Backend/internal/modules/account"
)

func Init(r *gin.Engine) {
	gMain := r.Group(`account`)
	initAccount(gMain)
}

func initAccount(g *gin.RouterGroup) {
	controller := account.Init()

	g.POST(``, controller.Create)

	g.Use(guards.CheckAPI())

	g.GET(``, controller.GetMe)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
