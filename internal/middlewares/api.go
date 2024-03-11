package middlewares

import (
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
)

// Достает API ключ из заголовков и если есть - получает пользователя
func API() gin.HandlerFunc {
	return func(c *gin.Context) {
		api := c.Request.Header.Get(`api-key`)

		if api != `` {
			accRep := account_rep.Repository
			account, _ := accRep.GetOne(&repositories.AccountGet{ApiKey: &api})
			if account != nil {
				c.Set(`account`, *account)
				c.Set(`accountID`, account.Username)
				c.Set(`accountUsername`, account.Username)
			}
		}

		c.Next()
	}
}
