package handlers

import (
	"github.com/gin-gonic/gin"

	account_handler "github.com/jourloy/X-Backend/internal/handlers/account"
	app_handler "github.com/jourloy/X-Backend/internal/handlers/app"
	building_handler "github.com/jourloy/X-Backend/internal/handlers/building"
	creature_handler "github.com/jourloy/X-Backend/internal/handlers/creature"
	deposit_handler "github.com/jourloy/X-Backend/internal/handlers/deposit"
	item_handler "github.com/jourloy/X-Backend/internal/handlers/item"
	operation_handler "github.com/jourloy/X-Backend/internal/handlers/operation"
	plan_handler "github.com/jourloy/X-Backend/internal/handlers/plan"
	resource_handler "github.com/jourloy/X-Backend/internal/handlers/resource"
	sector_handler "github.com/jourloy/X-Backend/internal/handlers/sector"
)

func Init(r *gin.Engine) {
	app_handler.Init(r)

	account_handler.Init(r)
	sector_handler.Init(r)

	deposit_handler.Init(r)
	resource_handler.Init(r)

	building_handler.Init(r)
	creature_handler.Init(r)

	item_handler.Init(r)

	operation_handler.Init(r)

	plan_handler.Init(r)
}
