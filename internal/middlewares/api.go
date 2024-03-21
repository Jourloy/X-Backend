package middlewares

import (
	"github.com/gin-gonic/gin"

	account_rep "github.com/jourloy/X-Backend/internal/modules/account/repository"
	"github.com/jourloy/X-Backend/internal/repositories"
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
