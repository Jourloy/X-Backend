package handlers

import (
	"github.com/gin-gonic/gin"
	account_handler "github.com/jourloy/X-Backend/internal/handlers/account"
)

func Init(r *gin.Engine) {
	account_handler.Init(r)
}
