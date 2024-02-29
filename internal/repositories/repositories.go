package repositories

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID     string `json:"id"`
	ApiKey string `json:"apiKey"`
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

type Colony struct {
	gorm.Model
	ID        string   `json:"id"`
	Balance   int      `json:"balance"`
	Storage   int      `json:"storage"`
	Workers   []Worker `json:"worker"`
	AccountID string   `json:"accountId"`
}

type Worker struct {
	gorm.Model
	ID         string `json:"id"`
	MaxStorage int    `json:"maxStorage"`
	Storage    []Item `json:"storage"`
	Location   string `json:"location"`
	ColonyID   string `json:"colonyId"`
}

type Item struct {
	gorm.Model
	ID       string `json:"id"`
	Type     string `json:"type"`
	WorkerID string `json:"workerId"`
}

type Place struct {
	gorm.Model
	ID        string     `json:"id"`
	Resources []Resource `json:"resources"`
}

type Resource struct {
	gorm.Model
	ID      string `json:"id"`
	Type    string `json:"type"`
	Amount  int    `json:"amount"`
	Weight  int    `json:"weight"`
	PlaceID string `json:"placeId"`
}
