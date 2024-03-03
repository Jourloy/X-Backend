package repositories

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID     string `json:"id"`
	ApiKey string `json:"apiKey"`
}

type IAccountRepository interface {
	Create(account *Account)
	GetOne(apiKey string) Account
	UpdateOne(account *Account)
	DeleteOne(account *Account)
}

// Модель сектора
type Sector struct {
	gorm.Model
	ID string `json:"id"`

	X int `json:"x"`
	Y int `json:"y"`

	// Постройки

	Townhalls []Townhall `json:"townhall"`
	Towers    []Tower    `json:"tower"`
	Storages  []Storage  `json:"storages"`
	Walls     []Wall     `json:"walls"`
	Plans     []Plan     `json:"plans"`

	// Существа

	Workers  []Worker  `json:"workes"`
	Warriors []Warrior `json:"warriors"`
	Traders  []Trader  `json:"traders"`

	// Ресурсы

	Deposits  []Deposit  `json:"deposits"`
	Resources []Resource `json:"resources"`

	// Предметы

	Items []Item `json:"items"`
}

// Модель поиска сектора
type SectorGetAll struct {
	X     *int
	Y     *int
	Limit *int
}

// Репозиторий сектора
type ISectorRepository interface {
	Create(sector *Sector)
	GetOne(id string) Sector
	GetAll(query SectorGetAll) []Sector
	UpdateOne(sector *Sector)
	DeleteOne(sector *Sector)
}

// Модель залежи ресурсов
type Deposit struct {
	gorm.Model
	ID       string `json:"id"`
	Type     string `json:"type"`
	Amount   int    `json:"amount"`
	SectorID string `json:"sectorId"`
}

// Модель ресурсов
type Resource struct {
	gorm.Model
	ID         string `json:"id"`
	Type       string `json:"type"`
	Amount     int    `json:"amount"`
	Weight     int    `json:"weight"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	ParentID   string `json:"parentId"`
	ParentType string `json:"parentType"`
	SectorID   string `json:"sectorId"`
	AccountID  string `json:"accountId"`
	CreatorID  string `json:"creatorId"`
}

// Структура поиска ресурсов
type ResourceGetAll struct {

	// Переделать здесь и в сервисе

	Limit *int
}

// Репозиторий ресурсов
type IResourceRepository interface {
	Create(resource *Resource, sectorID string)
	GetOne(id string, sectorID string) Resource
	GetAll(query ResourceGetAll, sectorID string) []Resource
	UpdateOne(resource *Resource)
	DeleteOne(resource *Resource)
}

// Модель вещи
type Item struct {
	gorm.Model
	ID         string `json:"id"`
	Type       string `json:"type"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	ParentID   string `json:"workerId"`
	ParentType string `json:"parentType"`
	AccountID  string `json:"accountId"`
	CreatorID  string `json:"creatorId"`
}

// Структура поиска вещи
type ItemGetAll struct {

	// Переделать здесь и в сервисе

	Limit *int
}

// Репозиторий вещи
type IItemRepository interface {
	Create(item *Item, parentID string, accountID string)
	GetOne(id string, accountID string) Item
	GetAll(query ItemGetAll, accountID string) []Item
	UpdateOne(item *Item)
	DeleteOne(item *Item)
}

//////// Постройки ////////

// Модель главного здания
type Townhall struct {
	gorm.Model
	ID            string `json:"id"`
	MaxDurability int    `json:"maxDurability"`
	Durability    int    `json:"durability"`
	MaxStorage    int    `json:"maxStorage"`
	UsedStorage   int    `json:"usedStorage"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Storage       []Item `json:"storage"`
	SectorID      string `json:"sectorId"`
	AccountID     string `json:"accountId"`
}

// Модель башни
type Tower struct {
	gorm.Model
	ID            string `json:"id"`
	MaxDurability int    `json:"maxDurability"`
	Durability    int    `json:"durability"`
	Level         int    `json:"level"`
	MaxStorage    int    `json:"maxStorage"`
	UsedStorage   int    `json:"usedStorage"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Storage       []Item `json:"storage"`
	SectorID      string `json:"sectorId"`
	AccountID     string `json:"accountId"`
}

// Модель хранилища
type Storage struct {
	gorm.Model
	ID            string `json:"id"`
	MaxDurability int    `json:"maxDurability"`
	Durability    int    `json:"durability"`
	Level         int    `json:"level"`
	MaxStorage    int    `json:"maxStorage"`
	UsedStorage   int    `json:"usedStorage"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Storage       []Item `json:"storage"`
	SectorID      string `json:"sectorId"`
	AccountID     string `json:"accountId"`
}

// Модель стены
type Wall struct {
	gorm.Model
	ID            string `json:"id"`
	MaxDurability int    `json:"maxDurability"`
	Durability    int    `json:"durability"`
	Level         int    `json:"level"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Storage       []Item `json:"storage"`
	SectorID      string `json:"sectorId"`
	AccountID     string `json:"accountId"`
}

// Модель планируемой постройки
type Plan struct {
	gorm.Model
	ID          string `json:"id"`
	MaxProgress int    `json:"maxProgress"`
	Progress    int    `json:"progress"`
	Type        string `json:"type"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	Storage     []Item `json:"storage"`
	SectorID    string `json:"sectorId"`
	AccountID   string `json:"accountId"`
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
	Create(plan *Plan, accountID string)
	GetOne(id string, accountID string) Plan
	GetAll(query PlanGetAll, accountID string) []Plan
	UpdateOne(plan *Plan)
	DeleteOne(plan *Plan)
}

//////// Существа ////////

// Модель рабочего
type Worker struct {
	gorm.Model
	ID          string `json:"id"`
	MaxStorage  int    `json:"maxStorage"`
	UsedStorage int    `json:"usedStorage"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	MaxHealth   int    `json:"maxHealth"`
	Health      int    `json:"health"`
	Storage     []Item `json:"storage"`
	SectorID    string `json:"sectorId"`
	AccountID   string `json:"accountId"`
}

// Структура поиска рабочего
type WorkerGetAll struct {
	MaxStorage  *int
	UsedStorage *int
	X           *int
	Y           *int
	MaxHealth   *int
	Health      *int
	Limit       *int
}

// Репозиторий рабочего
type IWorkerRepository interface {
	Create(worker *Worker, accountID string)
	GetOne(id string, accountID string) Worker
	GetAll(query WorkerGetAll, accountID string) []Worker
	UpdateOne(worker *Worker)
	DeleteOne(worker *Worker)
}

// Модель воина
type Warrior struct {
	gorm.Model
	ID          string `json:"id"`
	MaxStorage  int    `json:"maxStorage"`
	UsedStorage int    `json:"usedStorage"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	MaxHealth   int    `json:"maxHealth"`
	Health      int    `json:"health"`
	Storage     []Item `json:"storage"`
	SectorID    string `json:"sectorId"`
	AccountID   string `json:"accountId"`
}

// Структура поиска воинов
type WarriorGetAll struct {
	MaxStorage  *int
	UsedStorage *int
	X           *int
	Y           *int
	MaxHealth   *int
	Health      *int
	Limit       *int
}

// Репозиторий воина
type IWarriorRepository interface {
	Create(warrior *Warrior, accountID string)
	GetOne(id string, accountID string) Warrior
	GetAll(query WarriorGetAll, accountID string) []Warrior
	UpdateOne(warrior *Warrior)
	DeleteOne(warrior *Warrior)
}

// Модель торговца
type Trader struct {
	gorm.Model
	ID          string `json:"id"`
	MaxStorage  int    `json:"maxStorage"`
	UsedStorage int    `json:"usedStorage"`
	X           int    `json:"x"`
	Y           int    `json:"y"`
	MaxHealth   int    `json:"maxHealth"`
	Health      int    `json:"health"`
	Storage     []Item `json:"storage"`
	SectorID    string `json:"sectorId"`
	AccountID   string `json:"accountId"`
}

// Структура поиска торговца
type TraderGetAll struct {
	MaxStorage  *int
	UsedStorage *int
	X           *int
	Y           *int
	MaxHealth   *int
	Health      *int
	Limit       *int
}

// Репозиторий торговца
type ITraderRepository interface {
	Create(trader *Trader, accountID string)
	GetOne(id string, accountID string) Trader
	GetAll(query TraderGetAll, accountID string) []Trader
	UpdateOne(trader *Trader)
	DeleteOne(trader *Trader)
}
