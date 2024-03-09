package repositories

import (
	"time"

	"gorm.io/gorm"
)

// Модель аккаунта
type Account struct {
	ID        string         `json:"id" gorm:"primarykey"`
	ApiKey    string         `json:"apiKey"`
	Username  string         `json:"username"`
	Balance   int            `json:"balance"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type AccountCreate struct {
	Username string `json:"username"`
}

type AccountGet struct {
	ID       *string `json:"id"`
	ApiKey   *string `json:"apiKey"`
	Username *string `json:"username"`
	Balance  *int    `json:"balance"`
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

	// Графы
	Nodes []Node `json:"nodes"`

	// Постройки
	Townhalls []Townhall `json:"townhalls"`
	Towers    []Tower    `json:"towers"`
	Storages  []Storage  `json:"storages"`
	Walls     []Wall     `json:"walls"`
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

	X         int  `json:"x"`
	Y         int  `json:"y"`
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
	ID        string         `json:"id"`
	Type      string         `json:"type"`
	Amount    int            `json:"amount"`
	X         int            `json:"x"`
	Y         int            `json:"y"`
	SectorID  string         `json:"sectorId"`
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
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Amount     int            `json:"amount"`
	Weight     int            `json:"weight"`
	X          int            `json:"x"`
	Y          int            `json:"y"`
	ParentID   string         `json:"parentId"`
	ParentType string         `json:"parentType"`
	SectorID   string         `json:"sectorId"`
	CreatorID  string         `json:"creatorId"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
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

// Модель шаблона ресурсов
type ResourceTemplate struct {
	ID        string         `json:"id"`
	Type      string         `json:"type"`
	Amount    int            `json:"amount"`
	Weight    int            `json:"weight"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура поиска шаблона ресурсов
type ResourceTemplateGetAll struct {
	Type   *string
	Amount *int
	Weight *int
	Limit  *int
}

// Репозиторий шаблона ресурсов
type IResourceTemplateRepository interface {
	Create(resourceTemplate *ResourceTemplate)
	GetOne(resourceTemplate ResourceTemplate) ResourceTemplate
	GetAll(query ResourceTemplateGetAll) []ResourceTemplate
	UpdateOne(resource *ResourceTemplate)
	DeleteOne(resource *ResourceTemplate)
}

// Модель предмета
type Item struct {
	ID         string         `json:"id" gorm:"primarykey"`
	Type       string         `json:"type"`
	X          int            `json:"x"`
	Y          int            `json:"y"`
	ParentID   string         `json:"parentId"`
	ParentType string         `json:"parentType"`
	SectorID   string         `json:"sectorId"`
	CreatorID  string         `json:"creatorId"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
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

// Модель шаблона предмета
type ItemTemplate struct {
	ID        string         `json:"id"`
	Type      string         `json:"type"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Структура поиска шаблона предмета
type ItemTemplateGetAll struct {
	Type  *string
	Limit *int
}

// Репозиторий шаблона предмета
type IItemTemplateRepository interface {
	Create(itemTemplate *ItemTemplate)
	GetOne(itemTemplate *ItemTemplate)
	GetAll(query ItemTemplateGetAll) []ItemTemplate
	UpdateOne(itemTemplate *ItemTemplate)
	DeleteOne(itemTemplate *ItemTemplate)
}

//////// Постройки ////////

// Модель главного здания
type Townhall struct {
	ID            string         `json:"id" gorm:"primarykey"`
	MaxDurability int            `json:"maxDurability"`
	Durability    int            `json:"durability"`
	MaxStorage    int            `json:"maxStorage"`
	UsedStorage   int            `json:"usedStorage"`
	X             int            `json:"x"`
	Y             int            `json:"y"`
	Storage       []Item         `json:"storage" gorm:"foreignKey:ParentID"`
	SectorID      string         `json:"sectorId"`
	AccountID     string         `json:"accountId"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Структура поиска главного здания
type TownhallGetAll struct {
	MaxDurability *int
	Durability    *int
	MaxStorage    *int
	UsedStorage   *int
	X             *int
	Y             *int
	Limit         *int
}

// Репозиторий главного здания
type ITownhallRepository interface {
	Create(townhall *Townhall, accountID string)
	GetOne(townhall *Townhall)
	GetAll(query TownhallGetAll, accountID string) []Townhall
	UpdateOne(townhall *Townhall)
	DeleteOne(townhall *Townhall)
}

// Модель башни
type Tower struct {
	ID            string         `json:"id"`
	MaxDurability int            `json:"maxDurability"`
	Durability    int            `json:"durability"`
	Level         int            `json:"level"`
	MaxStorage    int            `json:"maxStorage"`
	UsedStorage   int            `json:"usedStorage"`
	X             int            `json:"x"`
	Y             int            `json:"y"`
	Storage       []Item         `json:"storage" gorm:"foreignKey:ParentID"`
	SectorID      string         `json:"sectorId"`
	AccountID     string         `json:"accountId"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Структура поиска башни
type TowerGetAll struct {
	MaxDurability *int
	Durability    *int
	Level         *int
	MaxStorage    *int
	UsedStorage   *int
	X             *int
	Y             *int
	Limit         *int
}

// Репозиторий башни
type ITowerRepository interface {
	Create(tower *Tower, accountID string)
	GetOne(tower *Tower)
	GetAll(query TowerGetAll, accountID string) []Tower
	UpdateOne(tower *Tower)
	DeleteOne(tower *Tower)
}

// Модель хранилища
type Storage struct {
	ID            string         `json:"id"`
	MaxDurability int            `json:"maxDurability"`
	Durability    int            `json:"durability"`
	Level         int            `json:"level"`
	MaxStorage    int            `json:"maxStorage"`
	UsedStorage   int            `json:"usedStorage"`
	X             int            `json:"x"`
	Y             int            `json:"y"`
	Storage       []Item         `json:"storage" gorm:"foreignKey:ParentID"`
	SectorID      string         `json:"sectorId"`
	AccountID     string         `json:"accountId"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Структура поиска хранилища
type StorageGetAll struct {
	MaxDurability *int
	Durability    *int
	Level         *int
	MaxStorage    *int
	UsedStorage   *int
	X             *int
	Y             *int
	Limit         *int
}

// Репозиторий хранилища
type IStorageRepository interface {
	Create(storage *Storage, accountID string)
	GetOne(storage *Storage)
	GetAll(query StorageGetAll, accountID string) []Storage
	UpdateOne(storage *Storage)
	DeleteOne(storage *Storage)
}

// Модель стены
type Wall struct {
	ID            string         `json:"id"`
	MaxDurability int            `json:"maxDurability"`
	Durability    int            `json:"durability"`
	Level         int            `json:"level"`
	X             int            `json:"x"`
	Y             int            `json:"y"`
	Storage       []Item         `json:"storage" gorm:"foreignKey:ParentID"`
	SectorID      string         `json:"sectorId"`
	AccountID     string         `json:"accountId"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Структура поиска стены
type WallGetAll struct {
	MaxDurability *int
	Durability    *int
	Level         *int
	Y             *int
	X             *int
	Limit         *int
}

// Репозиторий  стены
type IWallRepository interface {
	Create(wall *Wall, accountID string)
	GetOne(wall *Wall)
	GetAll(query WallGetAll, accountID string) []Wall
	UpdateOne(wall *Wall)
	DeleteOne(wall *Wall)
}

// Модель планируемой постройки
type Plan struct {
	ID          string         `json:"id"`
	MaxProgress int            `json:"maxProgress"`
	Progress    int            `json:"progress"`
	Type        string         `json:"type"`
	X           int            `json:"x"`
	Y           int            `json:"y"`
	Storage     []Item         `json:"storage" gorm:"foreignKey:ParentID"`
	SectorID    string         `json:"sectorId"`
	AccountID   string         `json:"accountId"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
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

// Модель рынка
type Market struct {
	ID            string         `json:"id"`
	MaxDurability int            `json:"maxDurability"`
	Durability    int            `json:"durability"`
	Level         int            `json:"level"`
	MaxStorage    int            `json:"maxStorage"`
	UsedStorage   int            `json:"usedStorage"`
	X             int            `json:"x"`
	Y             int            `json:"y"`
	Resources     []Resource     `json:"resources" gorm:"foreignKey:ParentID"`
	Items         []Item         `json:"items" gorm:"foreignKey:ParentID"`
	SectorID      string         `json:"sectorId"`
	AccountID     string         `json:"accountId"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// Структура создания рынка
type MarketCreate struct {
	X         int    `json:"x"`
	Y         int    `json:"y"`
	SectorID  string `json:"sectorId"`
	AccountID string `json:"accountId"`
}

// Структура поиска рынка
type MarketGetAll struct {
	MaxDurability *int
	Durability    *int
	Level         *int
	MaxStorage    *int
	UsedStorage   *int
	X             *int
	Y             *int
	Limit         *int
}

// Репозиторий рынка
type IMarketRepository interface {
	Create(market *MarketCreate)
	GetOne(market *Market)
	GetAll(query MarketGetAll, accountID string) []Market
	UpdateOne(market *Market)
	DeleteOne(market *Market)
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
	Items []Item `json:"items" gorm:"foreignKey:ParentID"`

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
}

// Репозиторий существа
type CreatureRepository interface {
	Create(creatrue *CreatureCreate) (*Creature, error)
	GetOne(query *CreatureGet) (*Creature, error)
	GetAll(query *CreatureGet) (*[]Creature, error)
	UpdateOne(creature *Creature) error
	DeleteOne(creature *Creature) error
}
