package repositories

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID       string    `json:"id"`
	ApiKey   string    `json:"apiKey"`
	Colonies []Village `json:"villages"`
}

type IAccountRepository interface {
	// Create создает аккаунт
	Create(account *Account)
	// GetOne возвращает первый аккаунт, попавший под условие
	GetOne(apiKey string) Account
	// UpdateOne обновляет аккаунт
	UpdateOne(account *Account)
	// DeleteOne удаляет аккаунт
	DeleteOne(account *Account)
}

type Sector struct {
	gorm.Model
	ID        string     `json:"id"`
	Resources []Resource `json:"resources"`
}

type SectorFindAll struct {
	Limit *int
}

type ISectorRepository interface {
	// Create создает сектор
	Create(sector *Sector)
	// GetOne возвращает первый сектор, попавший под условие
	GetOne(id string) Sector
	// GetAll возвращает все сектора
	GetAll(q SectorFindAll) []Sector
	// UpdateOne обновляет сектор
	UpdateOne(sector *Sector)
	// DeleteOne удаляет сектор
	DeleteOne(sector *Sector)
}

type Resource struct {
	gorm.Model
	ID       string `json:"id"`
	Type     string `json:"type"`
	Amount   int    `json:"amount"`
	Weight   int    `json:"weight"`
	SectorID string `json:"sectorId"`
}

type ResourceFindAll struct {
	Limit  *int
	Type   *string
	Amount *int
	Weight *int
}

type IResourceRepository interface {
	// Create создает ресурс
	Create(resource *Resource, sectorID string)
	// GetOne возвращает первый ресурс, попавший под условие
	GetOne(id string, sectorID string) Resource
	// GetAll возвращает все ресурсы
	GetAll(sectorID string, q ResourceFindAll) []Resource
	// UpdateOne обновляет ресурс
	UpdateOne(resource *Resource)
	// DeleteOne удаляет ресурс
	DeleteOne(resource *Resource)
}

type Village struct {
	gorm.Model
	ID          string    `json:"id"`
	Balance     int       `json:"balance"`
	MaxStorage  int       `json:"maxStorage"`
	UsedStorage int       `json:"usedStorage"`
	Storage     []Item    `json:"storage"`
	Workers     []Worker  `json:"worker"`
	Warriors    []Warrior `json:"warriors"`
	AccountID   string    `json:"accountId"`
	SectorID    string    `json:"sectorId"`
}

type VillageFindAll struct {
	Limit       *int
	Balance     *int
	MaxStorage  *int
	UsedStorage *int
}

type IVillageRepository interface {
	// Create создает поселение
	Create(village *Village, accountID string, sectorID string)
	// GetOne возвращает первое поселение, попавшее под условие
	GetOne(id string, accountID string) Village
	// GetAll возвращает все поселения
	GetAll(accountID string, q VillageFindAll) []Village
	// UpdateOne обновляет поселение
	UpdateOne(village *Village)
	// DeleteOne удаляет поселение
	DeleteOne(village *Village)
}

type Worker struct {
	gorm.Model
	ID            string `json:"id"`
	MaxStorage    int    `json:"maxStorage"`
	UsedStorage   int    `json:"usedStorage"`
	Location      string `json:"location"`
	ToArrival     int    `json:"toTarget"`
	FromDeparture int    `json:"fromDeparture"`
	Storage       []Item `json:"storage"`
	VillageID     string `json:"villageId"`
	AccountID     string `json:"accountId"`
}

type WorkerFindAll struct {
	MaxStorage    *int
	UsedStorage   *int
	Location      *string
	ToArrival     *int
	FromDeparture *int
	Limit         *int
}

type IWorkerRepository interface {
	// Create создает рабочего
	Create(worker *Worker, villageID string, accountID string)
	// GetOne возвращает первого рабочего, попавшего под условие
	GetOne(id string, accountID string) Worker
	// GetAll возвращает всех рабочих
	GetAll(accountID string, q WorkerFindAll) []Worker
	// UpdateOne обновляет рабочего
	UpdateOne(worker *Worker)
	// DeleteOne удаляет рабочего
	DeleteOne(worker *Worker)
}

type Item struct {
	gorm.Model
	ID        string `json:"id"`
	Type      string `json:"type"`
	ParentID  string `json:"workerId"`
	AccountID string `json:"accountId"`
}

type ItemFindAll struct {
	Limit    *int
	Type     *string
	ParentID *string
}

type IItemRepository interface {
	// Create создает вещь
	Create(item *Item, parentID string, aID string)
	// GetOne возвращает первую вещь, попавшую под условие
	GetOne(id string, aID string) Item
	// GetAll возвращает все вещи
	GetAll(q ItemFindAll, aID string) []Item
	// UpdateOne обновляет вещь
	UpdateOne(item *Item)
	// DeleteOne удаляет вещь
	DeleteOne(item *Item)
}

type Warrior struct {
	gorm.Model
	ID            string `json:"id"`
	MaxStorage    int    `json:"maxStorage"`
	UsedStorage   int    `json:"usedStorage"`
	Location      string `json:"location"`
	ToArrival     int    `json:"toTarget"`
	FromDeparture int    `json:"fromDeparture"`
	Health        int    `json:"health"`
	Storage       []Item `json:"storage"`
	VillageID     string `json:"villageId"`
	AccountID     string `json:"accountId"`
}

type WarriorFindAll struct {
	MaxStorage    *int
	UsedStorage   *int
	Location      *string
	ToArrival     *int
	FromDeparture *int
	Health        *int
	Limit         *int
}

type IWarriorRepository interface {
	// Create создает рабочего
	Create(warrior *Warrior, villageID string, accountID string)
	// GetOne возвращает первого рабочего, попавшего под условие
	GetOne(id string, accountID string) Warrior
	// GetAll возвращает всех рабочих
	GetAll(accountID string, q WarriorFindAll) []Warrior
	// UpdateOne обновляет рабочего
	UpdateOne(warrior *Warrior)
	// DeleteOne удаляет рабочего
	DeleteOne(warrior *Warrior)
}

type Trader struct {
	gorm.Model
	ID            string `json:"id"`
	MaxStorage    int    `json:"maxStorage"`
	UsedStorage   int    `json:"usedStorage"`
	Location      string `json:"location"`
	ToArrival     int    `json:"toTarget"`
	FromDeparture int    `json:"fromDeparture"`
	Storage       []Item `json:"storage"`
	VillageID     string `json:"villageId"`
	AccountID     string `json:"accountId"`
}

type TraderFindAll struct {
	MaxStorage    *int
	UsedStorage   *int
	Location      *string
	ToArrival     *int
	FromDeparture *int
	Limit         *int
}

type ITraderRepository interface {
	// Create создает торговца
	Create(trader *Trader, villageID string, accountID string)
	// GetOne возвращает первого торговца, попавшего под условие
	GetOne(id string, accountID string) Trader
	// GetAll возвращает всех торговцев
	GetAll(accountID string, q TraderFindAll) []Trader
	// UpdateOne обновляет торговца
	UpdateOne(trader *Trader)
	// DeleteOne удаляет торговца
	DeleteOne(trader *Trader)
}
