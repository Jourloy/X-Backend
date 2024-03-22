package repositories_init

import (
	account_rep "github.com/jourloy/X-Backend/internal/repositories/account"
	building_rep "github.com/jourloy/X-Backend/internal/repositories/building"
	creature_rep "github.com/jourloy/X-Backend/internal/repositories/creature"
	deposit_rep "github.com/jourloy/X-Backend/internal/repositories/deposit"
	item_rep "github.com/jourloy/X-Backend/internal/repositories/item"
	node_rep "github.com/jourloy/X-Backend/internal/repositories/node"
	operation_rep "github.com/jourloy/X-Backend/internal/repositories/operation"
	plan_rep "github.com/jourloy/X-Backend/internal/repositories/plan"
	resource_rep "github.com/jourloy/X-Backend/internal/repositories/resource"
	sector_rep "github.com/jourloy/X-Backend/internal/repositories/sector"
)

func Init() {
	account_rep.Init()
	building_rep.Init()
	creature_rep.Init()
	deposit_rep.Init()
	item_rep.Init()
	node_rep.Init()
	operation_rep.Init()
	plan_rep.Init()
	resource_rep.Init()
	sector_rep.Init()
}
