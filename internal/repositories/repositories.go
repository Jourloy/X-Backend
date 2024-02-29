package repositories

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID       string   `json:"id"`
	ApiKey   string   `json:"apiKey"`
	Colonies []Colony `json:"colonies"`
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

type Place struct {
	gorm.Model
	ID        string     `json:"id"`
	Resources []Resource `json:"resources"`
}

type IPlaceRepository interface {
	// Create создает место
	Create(place *Place)
	// GetOne возвращает первое место, попавшее под условие
	GetOne(id string) Place
	// UpdateOne обновляет место
	UpdateOne(place *Place)
	// DeleteOne удаляет место
	DeleteOne(place *Place)
}

type Resource struct {
	gorm.Model
	ID      string `json:"id"`
	Type    string `json:"type"`
	Amount  int    `json:"amount"`
	Weight  int    `json:"weight"`
	PlaceID string `json:"placeId"`
}

type IResourceRepository interface {
	// Create создает ресурс
	Create(resource *Resource, placeID string)
	// GetOne возвращает первый ресурс, попавший под условие
	GetOne(id string, placeID string) Resource
	// UpdateOne обновляет ресурс
	UpdateOne(resource *Resource)
	// DeleteOne удаляет ресурс
	DeleteOne(resource *Resource)
}

type Colony struct {
	gorm.Model
	ID         string   `json:"id"`
	Balance    int      `json:"balance"`
	MaxStorage int      `json:"maxStorage"`
	Storage    []Item   `json:"storage"`
	Workers    []Worker `json:"worker"`
	AccountID  string   `json:"accountId"`
	PlaceID    string   `json:"placeId"`
}

type IColonyRepository interface {
	// Create создает колонию
	Create(colony *Colony, accountID string, placeID string)
	// GetOne возвращает первую колонию, попавшую под условие
	GetOne(id string, accountID string) Colony
	// UpdateOne обновляет колонию
	UpdateOne(colony *Colony)
	// DeleteOne удаляет колонию
	DeleteOne(colony *Colony)
}

type Worker struct {
	gorm.Model
	ID          string `json:"id"`
	MaxStorage  int    `json:"maxStorage"`
	UsedStorage int    `json:"usedStorage"`
	Storage     []Item `json:"storage"`
	Location    string `json:"location"`
	ColonyID    string `json:"colonyId"`
	AccountID   string `json:"accountId"`
}

type IWorkerRepository interface {
	// Create создает рабочего
	Create(worker *Worker, colonyID string, accountID string)
	// GetOne возвращает первого рабочего, попавшего под условие
	GetOne(id string, accountID string) Worker
	// UpdateOne обновляет рабочего
	UpdateOne(worker *Worker)
	// DeleteOne удаляет рабочего
	DeleteOne(worker *Worker)
}

type Item struct {
	gorm.Model
	ID       string `json:"id"`
	Type     string `json:"type"`
	ParentID string `json:"workerId"`
}

type IItemRepository interface {
	// Create создает вещь
	Create(item *Item, parentID string)
	// GetOne возвращает первую вещь, попавшую под условие
	GetOne(id string, parentID string) Item
	// UpdateOne обновляет вещь
	UpdateOne(item *Item)
	// DeleteOne удаляет вещь
	DeleteOne(item *Item)
}
