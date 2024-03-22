package repositories

import (
	"time"

	"gorm.io/gorm"
)

// Модель аккаунта
type Account struct {
	ID string `json:"id" gorm:"primarykey"`

	// Динамические поля, задаются пользователем
	Race     string `json:"race"`
	Username string `json:"username"`

	ApiKey  string `json:"apiKey"`
	Balance int    `json:"balance"`
	IsAdmin bool   `json:"-"`

	// Мета
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type AccountCreate struct {
	Username string `json:"username"`
	Race     string `json:"race"`
}

type AccountGet struct {
	ID       *string `json:"id"`
	ApiKey   *string `json:"apiKey"`
	Username *string `json:"username"`
	Balance  *int    `json:"balance"`
	Race     *string `json:"race"`
}

type AccountRepository interface {
	Create(create *AccountCreate) (*Account, error)
	GetOne(query *AccountGet) (*Account, error)
	UpdateOne(account *Account) error
	DeleteOne(account *Account) error
}

// Модель сектора
type Sector struct {
	ID string `json:"id"`

	// Глобальные координаты
	X int `json:"x"`
	Y int `json:"y"`

	// Узлы
	Nodes []Node `json:"nodes"`

	// Постройки
	Buildings []Building `json:"buildings"`
	Plans     []Plan     `json:"plans"`

	// Существа
	Creatures []Creature `json:"creatures"`

	// Ресурсы
	Deposits  []Deposit  `json:"deposits"`
	Resources []Resource `json:"resources" gorm:"foreignKey:ParentID"`

	// Предметы
	Items []Item `json:"items" gorm:"foreignKey:ParentID"`

	// Мета
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type SectorCreate struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// Структура поиска сектора
type SectorGet struct {
	ID    *string
	X     *int
	Y     *int
	Limit *int
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
	ID string `json:"id"`

	X int `json:"x"`
	Y int `json:"y"`

	Walkable  bool `json:"walkable"`
	Difficult int  `json:"difficult"`

	SectorID string `json:"sectorId"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Модель поиска узла
type NodeGetAll struct {
	X         *int
	Y         *int
	Walkable  *bool
	Difficult *int
	SectorID  string
	Limit     *int
}

// Репозиторий сектора
type NodeRepository interface {
	Create(node *Node)
	GetOne(node *Node)
	GetAll(dest *[]Node, query NodeGetAll)
	UpdateOne(node *Node)
	DeleteOne(node *Node)
}

// Модель залежи ресурсов
type Deposit struct {
	ID string `json:"id"`

	X int `json:"x"`
	Y int `json:"y"`

	Type   string `json:"type"`
	Amount int    `json:"amount"`

	SectorID string `json:"sectorId"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура поиска залежей
type DepositGetAll struct {
	Type   *string
	Amount *int
	X      *int
	Y      *int
	Limit  *int
}

// Репозиторий залежей
type IDepositRepository interface {
	Create(deposit *Deposit)
	GetOne(deposit *Deposit)
	GetAll(query DepositGetAll, sectorID string) []Deposit
	UpdateOne(deposit *Deposit)
	DeleteOne(deposit *Deposit)
}

// Модель ресурсов
type Resource struct {
	ID string `json:"id"`

	X int `json:"x"`
	Y int `json:"y"`

	Type   string `json:"type"`
	Amount int    `json:"amount"`
	Weight int    `json:"weight"`

	ParentID   string `json:"parentId"`
	ParentType string `json:"parentType"`
	SectorID   string `json:"sectorId"`
	CreatorID  string `json:"creatorId"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура поиска ресурсов
type ResourceGetAll struct {
	Type       *string
	Amount     *int
	Weight     *int
	X          *int
	Y          *int
	ParentID   *string
	ParentType *string
	SectorID   *string
	CreatorID  *string
	Limit      *int
}

// Репозиторий ресурсов
type IResourceRepository interface {
	Create(resource *Resource)
	GetOne(resource Resource)
	GetAll(query ResourceGetAll) []Resource
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
	Type       *string
	X          *int
	Y          *int
	ParentID   *string
	ParentType *string
	SectorID   *string
	CreatorID  *string
	Limit      *int
}

// Репозиторий предмета
type IItemRepository interface {
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
