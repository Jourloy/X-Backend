package repositories

import (
	"time"

	"gorm.io/gorm"
)

// Модель аккаунта
type Account struct {
	// Задается при создании

	Race     string `json:"race"`
	Username string `json:"username"`

	// Задается по умолчанию

	ApiKey  string `json:"apiKey"`
	Balance int    `json:"balance"`
	IsAdmin bool   `json:"-"`

	// Мета

	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type AccountCreate struct {
	Username string `form:"username" json:"username"`
	Race     string `form:"race" json:"race"`
}

type AccountGet struct {
	ID       *string `form:"id" json:"id"`
	ApiKey   *string `form:"apiKey" json:"apiKey"`
	Username *string `form:"username" json:"username"`
	Balance  *int    `form:"balance" json:"balance"`
	Race     *string `form:"race" json:"race"`
	Limit    *int    `form:"limit" json:"limit"`
}

type AccountRepository interface {
	Create(create *AccountCreate) (*Account, error)
	GetOne(query *AccountGet) (*Account, error)
	UpdateOne(account *Account) error
	DeleteOne(account *Account) error
}

// Модель сектора
type Sector struct {
	// Задается при создании

	X int `json:"x"`
	Y int `json:"y"`

	// Дети

	Nodes     []Node     `json:"nodes"`
	Buildings []Building `json:"buildings"`
	Plans     []Plan     `json:"plans"`
	Creatures []Creature `json:"creatures"`
	Deposits  []Deposit  `json:"deposits"`
	Resources []Resource `json:"resources" gorm:"foreignKey:ParentID"`
	Items     []Item     `json:"items" gorm:"foreignKey:ParentID"`

	// Мета
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type SectorCreate struct {
	X int `form:"x" json:"x"`
	Y int `form:"y" json:"y"`
}

// Структура поиска сектора
type SectorGet struct {
	ID    *string `form:"id" json:"id"`
	X     *int    `form:"x" json:"x"`
	Y     *int    `form:"y" json:"y"`
	Limit *int    `form:"limit" json:"limit"`
}

// Репозиторий сектора
type SectorRepository interface {
	Create(sector *SectorCreate) (*Sector, error)
	GetOne(query *SectorGet) (*Sector, error)
	GetAll(query *SectorGet) (*[]Sector, error)
	UpdateOne(sector *Sector) error
	DeleteOne(sector *Sector) error
}

// Модель узла
type Node struct {
	// Задается при создании

	X         int  `json:"x"`
	Y         int  `json:"y"`
	Walkable  bool `json:"walkable"`
	Difficult int  `json:"difficult"`

	// Родители

	SectorID string `json:"sectorId"`

	// Мета
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type NodeCreate struct {
	X         int    `form:"x" json:"x"`
	Y         int    `form:"y" json:"y"`
	Walkable  bool   `form:"walkable" json:"walkable"`
	Difficult int    `form:"difficult" json:"difficult"`
	SectorID  string `form:"sectorId" json:"sectorId"`
}

// Модель поиска узла
type NodeGet struct {
	X         *int    `form:"x" json:"x"`
	Y         *int    `form:"y" json:"y"`
	Walkable  *bool   `form:"walkable" json:"walkable"`
	Difficult *int    `form:"difficult" json:"difficult"`
	SectorID  *string `form:"sectorId" json:"sectorId"`
	Limit     *int    `form:"limit" json:"limit"`
}

// Репозиторий сектора
type NodeRepository interface {
	Create(create *NodeCreate) (*Node, error)
	GetOne(query *NodeGet) (*Node, error)
	GetAll(query *NodeGet) (*[]Node, error)
	UpdateOne(node *Node) error
	DeleteOne(node *Node) error
}

// Модель залежи
type Deposit struct {
	// Задается при создании

	X      int    `json:"x"`
	Y      int    `json:"y"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`

	// Родители

	SectorID string `json:"sectorId"`

	// Мета

	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура создания залежи
type DepositCreate struct {
	X      int    `form:"x" json:"x"`
	Y      int    `form:"y" json:"y"`
	Type   string `form:"type" json:"type"`
	Amount int    `form:"amount" json:"amount"`
}

// Структура поиска залежей
type DepositGet struct {
	Type     *string `form:"type" json:"type"`
	Amount   *int    `form:"amount" json:"amount"`
	X        *int    `form:"x" json:"x"`
	Y        *int    `form:"y" json:"y"`
	SectorID *string `form:"sectorId" json:"sectorId"`
	Limit    *int    `form:"limit" json:"limit"`
}

// Репозиторий залежей
type DepositRepository interface {
	Create(create DepositCreate) (*Deposit, error)
	GetOne(query DepositGet) (*Deposit, error)
	GetAll(query DepositGet) (*[]Deposit, error)
	UpdateOne(deposit *Deposit) error
	DeleteOne(deposit *Deposit) error
}

// Модель ресурсов
type Resource struct {
	// Задается при создании

	X    int    `json:"x"`
	Y    int    `json:"y"`
	Type string `json:"type"`

	// Задается по умолчанию

	Amount int `json:"amount"`

	// Задается по шаблону

	Weight int `json:"weight"`

	// Родители

	ParentID   string `json:"parentId"`
	ParentType string `json:"parentType"`
	SectorID   string `json:"sectorId"`
	CreatorID  string `json:"creatorId"`

	// Мета

	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура создания ресурсов
type ResourceCreate struct {
	X    int    `form:"x" json:"x"`
	Y    int    `form:"y" json:"y"`
	Type string `form:"type" json:"type"`
}

// Структура поиска ресурсов
type ResourceGet struct {
	Type       *string `form:"type" json:"type`
	Amount     *int    `form:"amount" json:"amount"`
	Weight     *int    `form:"weight" json:"weight"`
	X          *int    `form:"x" json:"x"`
	Y          *int    `form:"y" json:"y"`
	ParentID   *string `form:"parentId" json:"parentId"`
	ParentType *string `form:"parentType" json:"parentType"`
	SectorID   *string `form:"sectorId" json:"sectorId"`
	CreatorID  *string `form:"creatorId" json:"creatorId"`
	Limit      *int    `form:"limir" json:"limir"`
}

// Репозиторий ресурсов
type ResourceRepository interface {
	Create(create ResourceCreate) (*Resource, error)
	GetOne(query ResourceGet) (*Resource, error)
	GetAll(query ResourceGet) (*[]Resource, error)
	UpdateOne(resource *Resource)
	DeleteOne(resource *Resource)
}

// Модель предмета
type Item struct {
	ID string `json:"id" gorm:"primarykey"`

	X int `json:"x"`
	Y int `json:"y"`

	Type   string `json:"type"`
	Weight int    `json:"weight"`

	ParentID   string `json:"parentId"`
	ParentType string `json:"parentType"`
	SectorID   string `json:"sectorId"`
	CreatorID  string `json:"creatorId"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура поиска предмета
type ItemGetAll struct {
	Type       *string `form:"" json:""`
	X          *int    `form:"" json:""`
	Y          *int    `form:"" json:""`
	ParentID   *string `form:"" json:""`
	ParentType *string `form:"" json:""`
	SectorID   *string `form:"" json:""`
	CreatorID  *string `form:"" json:""`
	Limit      *int    `form:"" json:""`
}

// Репозиторий предмета
type ItemRepository interface {
	Create(item *Item)
	GetOne(item *Item)
	GetAll(query ItemGetAll) []Item
	UpdateOne(item *Item)
	DeleteOne(item *Item)
}

// Модель операции
type Operation struct {
	// Динамические поля, задаются пользователем
	Price      int    `json:"price"`
	Amount     int    `json:"amount"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	IsResource bool   `json:"isResource"`
	IsItem     bool   `json:"isItem"`

	// Родители
	BuildingID string `json:"buildingID"`
	SectorID   string `json:"sectorId"`
	AccountID  string `json:"accountId"`

	// Мета
	ID        string         `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура создания операции
type OperationCreate struct {
	Price      int    `json:"price"`
	Amount     int    `json:"amount"`
	Type       string `json:"type"`
	Name       string `json:"name"`
	IsResource bool   `json:"isResource"`
	IsItem     bool   `json:"isItem"`
	BuildingID string `json:"buildingID"`
	SectorID   string `json:"sectorId"`
	AccountID  string `json:"accountId"`
}

// Структура поиска операции
type OperationGet struct {
	ID         *string  `json:"id,omitempty"`
	Price      *int     `json:"price,omitempty"`
	Amount     *int     `json:"amount,omitempty"`
	Type       *string  `json:"type,omitempty"`
	Name       *string  `json:"name,omitempty"`
	IsResource *bool    `json:"isResource,omitempty"`
	IsItem     *bool    `json:"isItem,omitempty"`
	BuildingID *float64 `json:"buildingID,omitempty"`
	SectorID   *string  `json:"sectorId,omitempty"`
	AccountID  *string  `json:"accountId,omitempty"`
	Limit      *int     `json:"limit,omitempty"`
}

// Репозиторий операции
type OperationRepository interface {
	Create(create *OperationCreate) (*Operation, error)
	GetOne(query *OperationGet) (*Operation, error)
	GetAll(query *OperationGet) (*[]Operation, error)
	UpdateOne(operation *Operation) error
	DeleteOne(operation *Operation) error
}

//////// Постройки ////////

// Модель постройки
type Building struct {
	ID string `json:"id" gorm:"primarykey"`

	X int `json:"x"`
	Y int `json:"y"`

	// Динамические поля, задаются пользователем
	Type string `json:"type"`

	// Динамические поля, задаются шаблоном
	MaxDurability int  `json:"maxDurability"`
	Durability    int  `json:"durability"`
	MaxStorage    int  `json:"maxStorage"`
	UsedStorage   int  `json:"usedStorage"`
	Level         int  `json:"level"`
	AttackRange   int  `json:"attackRange"`
	CanTrade      bool `json:"catTrade"`

	// Дети
	Items      []Item      `json:"items" gorm:"foreignKey:ParentID"`
	Resources  []Resource  `json:"resources" gorm:"foreignKey:ParentID"`
	Operations []Operation `json:"operations"`

	// Родители
	SectorID  string `json:"sectorId"`
	AccountID string `json:"accountId"`

	// Мета
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура создания постройки
type BuildingCreate struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Type      string `json:"type"`
	SectorID  string `json:"sectorId"`
	AccountID string `json:"accountId"`
}

// Структура поиска постройки
type BuildingGet struct {
	ID            *string  `json:"id,omitempty"`
	Type          *string  `json:"type,omitempty"`
	MaxDurability *int     `json:"maxDurability,omitempty"`
	Durability    *int     `json:"durability,omitempty"`
	MaxStorage    *float64 `json:"maxStorage,omitempty"`
	UsedStorage   *float64 `json:"usedStorage,omitempty"`
	Level         *float64 `json:"level,omitempty"`
	AttackRange   *float64 `json:"attackRange,omitempty"`
	CanTrade      *bool    `json:"canTrade,omitempty"`
	SectorID      *string  `json:"sectorId,omitempty"`
	AccountID     *string  `json:"accountId,omitempty"`
	Limit         *int     `json:"limit,omitempty"`
}

// Репозиторий постройки
type BuildingRepository interface {
	Create(create *BuildingCreate) (*Building, error)
	GetOne(query *BuildingGet) (*Building, error)
	GetAll(query *BuildingGet) (*[]Building, error)
	UpdateOne(building *Building) error
	DeleteOne(building *Building) error
}

// Модель планируемой постройки
type Plan struct {
	ID string `json:"id" gorm:"primarykey"`

	X int `json:"x"`
	Y int `json:"y"`

	// Динамические поля, задаются пользователем
	Type string `json:"type"`

	MaxProgress int `json:"maxProgress"`
	Progress    int `json:"progress"`

	// Дети
	Items     []Item     `json:"items" gorm:"foreignKey:ParentID"`
	Resources []Resource `json:"resources" gorm:"foreignKey:ParentID"`

	// Родители
	SectorID  string `json:"sectorId"`
	AccountID string `json:"accountId"`

	// Мета
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура создания планируемой постройки
type PlanCreate struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Type      string `json:"type"`
	SectorID  string `json:"sectorId"`
	AccountID string `json:"accountId"`
}

// Структура поиска планируемой постройки
type PlanGetAll struct {
	MaxProgress *int
	Progress    *int
	Type        *string
	Y           *int
	X           *int
	Limit       *int
}

// Репозиторий планируемой постройки
type IPlanRepository interface {
	Create(plan *PlanCreate)
	GetOne(plan *Plan)
	GetAll(query PlanGetAll, accountID string) []Plan
	UpdateOne(plan *Plan)
	DeleteOne(plan *Plan)
}

//////// Существа ////////

// Модель существа
type Creature struct {
	ID string `json:"id" gorm:"primarykey"`

	X int `json:"x"`
	Y int `json:"y"`

	// Динамические поля, задаются пользователем
	Race      string `json:"race"`
	IsWorker  bool   `json:"isWorker"`
	IsTrader  bool   `json:"isTrader"`
	IsWarrior bool   `json:"isWarrior"`

	// Динамические поля, задаются шаблоном
	MaxStorage         int     `json:"maxStorage"`
	UsedStorage        int     `json:"usedStorage"`
	RequireFood        float64 `json:"requireFood"`
	FatiguePerStep     float64 `json:"fatiguePerStep"`
	FatigueModificator float64 `json:"fatigueModificator"`
	Fatigue            float64 `json:"fatigue"`
	MaxHealth          int     `json:"maxHealth"`
	Health             int     `json:"health"`

	// Дети
	Items     []Item     `json:"items" gorm:"foreignKey:ParentID"`
	Resources []Resource `json:"resources" gorm:"foreignKey:ParentID"`

	// Родители
	SectorID  string `json:"sectorId"`
	AccountID string `json:"accountId"`

	// Мета
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура создания существа
type CreatureCreate struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Race      string `json:"race"`
	IsWorker  bool   `json:"isWorker"`
	IsTrader  bool   `json:"isTrader"`
	IsWarrior bool   `json:"isWarrior"`
	SectorID  string `json:"sectorId"`
	AccountID string `json:"accountId"`
}

// Структура поиска существа
type CreatureGet struct {
	ID                 *string  `json:"id,omitempty"`
	Race               *string  `json:"race,omitempty"`
	MaxStorage         *int     `json:"maxStorage,omitempty"`
	UsedStorage        *int     `json:"usedStorage,omitempty"`
	RequireCoins       *float64 `json:"requireCoins,omitempty"`
	RequireFood        *float64 `json:"requireFood,omitempty"`
	Fatigue            *float64 `json:"fatigue,omitempty"`
	FatiguePerStep     *float64 `json:"fatiguePerStep,omitempty"`
	FatigueModificator *float64 `json:"fatigueModificator,omitempty"`
	MaxHealth          *int     `json:"maxHealth,omitempty"`
	Health             *int     `json:"health,omitempty"`
	IsWorker           *bool    `json:"isWorker,omitempty"`
	IsTrader           *bool    `json:"isTrader,omitempty"`
	IsWarrior          *bool    `json:"isWarrior,omitempty"`
	SectorID           *string  `json:"sectorId,omitempty"`
	AccountID          *string  `json:"accountId,omitempty"`
	Limit              *int     `json:"limit,omitempty"`
}

// Репозиторий существа
type CreatureRepository interface {
	Create(creatrue *CreatureCreate) (*Creature, error)
	GetOne(query *CreatureGet) (*Creature, error)
	GetAll(query *CreatureGet) (*[]Creature, error)
	UpdateOne(creature *Creature) error
	DeleteOne(creature *Creature) error
}
