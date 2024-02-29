package repositories

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID       string   `json:"id"`
	ApiKey   string   `json:"apiKey"`
	Colonies []Colony `json:"colonies"`
}

type IAccountRepository interface {
	// Create создает объект в БД
	Create(account *Account)
	// GetOne возвращает первый объект, попавший под условие
	GetOne(apiKey string) Account
	// UpdateOne обновляет объект в БД
	UpdateOne(account *Account)
	// DeleteOne удаляет объект из БД
	DeleteOne(account *Account)
}

type Place struct {
	gorm.Model
	ID        string     `json:"id"`
	Resources []Resource `json:"resources"`
}

type IPlaceRepository interface {
	// Create создает объект в БД
	Create(place *Place)
	// GetOne возвращает первый объект, попавший под условие
	GetOne(id string) Place
	// UpdateOne обновляет объект в БД
	UpdateOne(place *Place)
	// DeleteOne удаляет объект из БД
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
	// Create создает объект в БД
	Create(resource *Resource, placeID string)
	// GetOne возвращает первый объект, попавший под условие
	GetOne(id string, placeID string) Resource
	// UpdateOne обновляет объект в БД
	UpdateOne(resource *Resource)
	// DeleteOne удаляет объект из БД
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
	// Create создает объект в БД
	Create(colony *Colony, accountID string, placeID string)
	// GetOne возвращает первый объект, попавший под условие
	GetOne(id string, accountID string) Colony
	// UpdateOne обновляет объект в БД
	UpdateOne(colony *Colony)
	// DeleteOne удаляет объект из БД
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
	// Create создает объект в БД
	Create(worker *Worker, colonyID string, accountID string)
	// GetOne возвращает первый объект, попавший под условие
	GetOne(id string, accountID string) Worker
	// UpdateOne обновляет объект в БД
	UpdateOne(worker *Worker)
	// DeleteOne удаляет объект из БД
	DeleteOne(worker *Worker)
}

type Item struct {
	gorm.Model
	ID       string `json:"id"`
	Type     string `json:"type"`
	ParentID string `json:"workerId"`
}

type IItemRepository interface {
	// Create создает объект в БД
	Create(item *Item, parentID string)
	// GetOne возвращает первый объект, попавший под условие
	GetOne(id string, parentID string) Item
	// UpdateOne обновляет объект в БД
	UpdateOne(item *Item)
	// DeleteOne удаляет объект из БД
	DeleteOne(item *Item)
}
