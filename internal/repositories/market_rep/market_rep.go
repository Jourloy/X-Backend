package market_rep

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var Repository repositories.IMarketRepository

type MarketRepository struct {
	db gorm.DB
}

// Init создает репозиторий рынка
func Init() {
	Repository = &MarketRepository{
		db: *storage.Database,
	}
}

// Create создает рынок
func (r *MarketRepository) Create(market *repositories.MarketCreate) {
	r.db.Create(&repositories.Market{
		ID:            uuid.NewString(),
		MaxDurability: 1000,
		Durability:    1000,
		Level:         1,
		MaxStorage:    1000,
		UsedStorage:   0,
		X:             market.X,
		Y:             market.Y,
		AccountID:     market.AccountID,
	})
}

// GetOne возвращает первый рынок, попавший под условие
func (r *MarketRepository) GetOne(market *repositories.Market) {
	r.db.First(&market, market)
}

// GetAll возвращает все рынки
func (r *MarketRepository) GetAll(query repositories.MarketGetAll, accountID string) []repositories.Market {
	var market = repositories.Market{
		MaxDurability: *query.MaxDurability,
		Durability:    *query.Durability,
		Level:         *query.Level,
		MaxStorage:    *query.MaxStorage,
		UsedStorage:   *query.UsedStorage,
		X:             *query.X,
		Y:             *query.Y,
		AccountID:     accountID,
	}
	var markets = []repositories.Market{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(market).Limit(limit).Find(&markets)
	return markets
}

// UpdateOne обновляет рынок
func (r *MarketRepository) UpdateOne(market *repositories.Market) {
	r.db.Save(&market)
}

// DeleteOne удаляет рынок
func (r *MarketRepository) DeleteOne(market *repositories.Market) {
	r.db.Delete(&market)
}
