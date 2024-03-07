package guards

import (
	"github.com/gin-gonic/gin"
)

// Достает пользователя из контекста. Если нет - отказ
func CheckAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, exist := c.Get(`account`)
		if exist {
			c.Next()
		} else {
			c.JSON(400, `api key is required`)
		}
	}
}
