# X-Backend

## Описание

Backend для API игры

## Установка

```bash
$ go mod tidy
```

##### .env

Не забудь создать `.env` из `.env.template`

## Запуск

```bash
$ go run ./cmd/http/main.go
```

## Тестирование

```bash
$ go test ./...
```

## Как добавлять модель

У модели должна быть структура

```golang
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
```

Она **всегда** содержит 2 обязательных поля

```golang
gorm.Model
ID          string `json:"id"`
```

Далее идут динамические поля, затем "дети" и в конце "родители"

```golang
// Динамическое поле
MaxStorage  int    `json:"maxStorage"`

// Дети
Storage     []Item `json:"storage"`

// Родители
AccountID   string `json:"accountId"`
```

`AccountID` должен присутствовать у всего, что так или иначе относится к пользователю

Далее нужно создать структуру поиска. В нее помещаются все динамические поля

```golang
// Структура поиска торговца
type TraderGetAll struct {}
```

И в конце создается репозиторий, который содержит основные функции. Если у модели есть `AccountID`, то `Create`, `GetOne` и `GetAll` должны принимать его

```golang
// Репозиторий торговца
type ITraderRepository interface {
	Create(trader *Trader, accountID string)
	GetOne(id string, accountID string) Trader
	GetAll(query TraderGetAll, accountID string) []Trader
	UpdateOne(trader *Trader)
	DeleteOne(trader *Trader)
}
```
