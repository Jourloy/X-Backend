package trader_rep

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var Repository repositories.ITraderRepository

type traderRepository struct {
	db gorm.DB
}

// Init создает репозиторий торговца
func Init() {
	Repository = &traderRepository{
		db: *storage.Database,
	}
}

// Create создает торговца
func (r *traderRepository) Create(trader *repositories.TraderCreate) {
	r.db.Create(&repositories.Trader{
		ID:           uuid.NewString(),
		MaxStorage:   100,
		UsedStorage:  0,
		X:            trader.X,
		Y:            trader.Y,
		MaxHealth:    100,
		Health:       100,
		RequireCoins: 0.5,
		RequireFood:  0.5,
		Fatigue:      0,
		AccountID:    trader.AccountID,
	})
}

// GetOne возвращает первого торговца, попавшего под условие
func (r *traderRepository) GetOne(trader *repositories.Trader) {
	r.db.First(&trader, trader)
}

func (r *traderRepository) GetAll(query repositories.TraderGetAll, accountID string) []repositories.Trader {
	var trader = repositories.Trader{
		MaxStorage:   *query.MaxStorage,
		UsedStorage:  *query.UsedStorage,
		X:            *query.X,
		Y:            *query.Y,
		MaxHealth:    *query.MaxHealth,
		Health:       *query.Health,
		RequireCoins: *query.RequireCoins,
		RequireFood:  *query.RequireFood,
		Fatigue:      *query.Fatigue,
		AccountID:    accountID,
	}
	var traders = []repositories.Trader{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(trader).Limit(limit).Find(&traders)
	return traders
}

// UpdateOne обновляет торговца
func (r *traderRepository) UpdateOne(trader *repositories.Trader) {
	r.db.Save(&trader)
}

// DeleteOne удаляет торговца
func (r *traderRepository) DeleteOne(trader *repositories.Trader) {
	r.db.Delete(&trader)
}
