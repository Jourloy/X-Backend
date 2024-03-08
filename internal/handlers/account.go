package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jourloy/X-Backend/internal/modules/account"
)

// @Tag.name Аккаунт
// @Tag.description Эндпоинты для работы с аккаунтом

func InitAccount(g *gin.RouterGroup) {

	controller := account.Init()

	g.POST(``, controller.Create)
	g.GET(``, controller.GetMe)
	g.PATCH(``, controller.UpdateOne)
	g.DELETE(``, controller.DeleteOne)
}
