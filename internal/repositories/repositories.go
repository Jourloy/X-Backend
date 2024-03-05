package repositories

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID      string `json:"id"`
	ApiKey  string `json:"apiKey"`
	Balance int    `json:"balance"`
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

// Структура поиска залежей
type DepositGetAll struct {
	Type   *string
	Amount *int
	Limit  *int
}

// Репозиторий ресурсов
type IDepositRepository interface {
	Create(deposit *Deposit)
	GetOne(id string, sectorID string) Deposit
	GetAll(query DepositGetAll, sectorID string) []Deposit
	UpdateOne(deposit *Deposit)
	DeleteOne(deposit *Deposit)
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
	CreatorID  string `json:"creatorId"`
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
	GetOne(id string) Resource
	GetAll(query ResourceGetAll) []Resource
	UpdateOne(resource *Resource)
	DeleteOne(resource *Resource)
}

// Модель шаблона ресурсов
type ResourceTemplate struct {
	gorm.Model
	ID     string `json:"id"`
	Type   string `json:"type"`
	Amount int    `json:"amount"`
	Weight int    `json:"weight"`
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
	GetOne(id string) ResourceTemplate
	GetAll(query ResourceTemplateGetAll) []ResourceTemplate
	UpdateOne(resource *ResourceTemplate)
	DeleteOne(resource *ResourceTemplate)
}

// Модель предмета
type Item struct {
	gorm.Model
	ID         string `json:"id"`
	Type       string `json:"type"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	ParentID   string `json:"workerId"`
	ParentType string `json:"parentType"`
	SectorID   string `json:"sectorId"`
	CreatorID  string `json:"creatorId"`
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
	GetOne(id string) Item
	GetAll(query ItemGetAll) []Item
	UpdateOne(item *Item)
	DeleteOne(item *Item)
}

// Модель шаблона предмета
type ItemTemplate struct {
	gorm.Model
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Структура поиска шаблона предмета
type ItemTemplateGetAll struct {
	Type  *string
	Limit *int
}

// Репозиторий шаблона предмета
type IItemTemplateRepository interface {
	Create(itemTemplate *ItemTemplate)
	GetOne(id string) ItemTemplate
	GetAll(query ItemTemplateGetAll) []ItemTemplate
	UpdateOne(itemTemplate *ItemTemplate)
	DeleteOne(itemTemplate *ItemTemplate)
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
	GetOne(id string, accountID string) Townhall
	GetAll(query TownhallGetAll, accountID string) []Townhall
	UpdateOne(townhall *Townhall)
	DeleteOne(townhall *Townhall)
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
	GetOne(id string, accountID string) Tower
	GetAll(query TowerGetAll, accountID string) []Tower
	UpdateOne(tower *Tower)
	DeleteOne(tower *Tower)
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
	GetOne(id string, accountID string) Storage
	GetAll(query StorageGetAll, accountID string) []Storage
	UpdateOne(storage *Storage)
	DeleteOne(storage *Storage)
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
	GetOne(id string, accountID string) Wall
	GetAll(query WallGetAll, accountID string) []Wall
	UpdateOne(wall *Wall)
	DeleteOne(wall *Wall)
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

// Модель рынка
type Market struct {
	gorm.Model
	ID            string     `json:"id"`
	MaxDurability int        `json:"maxDurability"`
	Durability    int        `json:"durability"`
	Level         int        `json:"level"`
	MaxStorage    int        `json:"maxStorage"`
	UsedStorage   int        `json:"usedStorage"`
	X             int        `json:"x"`
	Y             int        `json:"y"`
	Resources     []Resource `json:"resources"`
	Items         []Item     `json:"items"`
	SectorID      string     `json:"sectorId"`
	AccountID     string     `json:"accountId"`
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
	Create(market *Market, accountID string)
	GetOne(id string, accountID string) Market
	GetAll(query MarketGetAll, accountID string) []Market
	UpdateOne(market *Market)
	DeleteOne(market *Market)
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
