package trader_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-trader]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.ITraderRepository

type traderRepository struct {
	db gorm.DB
}

// Init создает репозиторий торговца
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Trader{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &traderRepository{
		db: *storage.Database,
	}
}

// Create создает торговца
func (r *traderRepository) Create(trader *repositories.Trader, accountId string) {
	r.db.Create(&repositories.Trader{
		ID:           uuid.NewString(),
		MaxStorage:   trader.MaxStorage,
		UsedStorage:  trader.UsedStorage,
		X:            trader.X,
		Y:            trader.Y,
		MaxHealth:    trader.MaxHealth,
		Health:       trader.Health,
		RequireCoins: 0.5,
		RequireFood:  0.5,
		Fatigue:      0,
		AccountID:    accountId,
	})
}

// GetOne возвращает первого торговца, попавшего под условие
func (r *traderRepository) GetOne(id string, accountID string) repositories.Trader {
	var trader = repositories.Trader{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&trader)
	return trader
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
