package building_templates

type BuildingTemplate struct {
	MaxDurability int  `json:"maxDurability"`
	Durability    int  `json:"durability"`
	MaxStorage    int  `json:"maxStorage"`
	UsedStorage   int  `json:"usedStorage"`
	Level         int  `json:"level"`
	AttackRange   int  `json:"attackRange"`
	CanTrade      bool `json:"canTrade"`
}

var Townhall = BuildingTemplate{
	MaxDurability: 1000,
	Durability:    1000,
	MaxStorage:    200,
	UsedStorage:   0,
	Level:         0,
	AttackRange:   10,
	CanTrade:      true,
}

var Tower = BuildingTemplate{
	MaxDurability: 500,
	Durability:    500,
	MaxStorage:    100,
	UsedStorage:   0,
	Level:         0,
	AttackRange:   20,
	CanTrade:      false,
}

var Wall = BuildingTemplate{
	MaxDurability: 500,
	Durability:    500,
	MaxStorage:    0,
	UsedStorage:   0,
	Level:         0,
	AttackRange:   0,
	CanTrade:      false,
}

var Storage = BuildingTemplate{
	MaxDurability: 500,
	Durability:    500,
	MaxStorage:    1000,
	UsedStorage:   0,
	Level:         0,
	AttackRange:   0,
	CanTrade:      false,
}

var Market = BuildingTemplate{
	MaxDurability: 500,
	Durability:    500,
	MaxStorage:    1000,
	UsedStorage:   0,
	Level:         0,
	AttackRange:   0,
	CanTrade:      true,
}
