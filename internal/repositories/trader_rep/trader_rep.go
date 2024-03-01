package trader_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-trader]`,
		Level:  log.DebugLevel,
	})
)

type TraderRepository struct {
	db gorm.DB
}

// InitTraderRepository создает репозиторий
func InitTraderRepository(db gorm.DB) repositories.ITraderRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Trader{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &TraderRepository{
		db: db,
	}
}

// Create создает торговца
func (r *TraderRepository) Create(trader *repositories.Trader, villageID string, accountId string) {
	r.db.Create(&repositories.Trader{
		ID:            uuid.NewString(),
		Location:      trader.Location,
		MaxStorage:    200,
		UsedStorage:   0,
		FromDeparture: 0,
		ToArrival:     0,
		VillageID:     villageID,
		AccountID:     accountId,
	})
}

// GetOne возвращает первого торговца, попавшего под условие
func (r *TraderRepository) GetOne(id string, accountID string) repositories.Trader {
	var trader = repositories.Trader{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&trader)
	return trader
}

// GetAll возвращает всех торговцев
func (r *TraderRepository) GetAll(accountID string, q repositories.TraderFindAll) []repositories.Trader {
	var trader = repositories.Trader{
		AccountID:     accountID,
		UsedStorage:   *q.UsedStorage,
		MaxStorage:    *q.MaxStorage,
		Location:      *q.Location,
		FromDeparture: *q.FromDeparture,
		ToArrival:     *q.ToArrival,
	}
	var traders = []repositories.Trader{}

	limit := -1
	if q.Limit != nil {
		limit = *q.Limit
	}

	r.db.Model(trader).Limit(limit).Find(&traders)
	return traders
}

// UpdateOne обновляет торговца
func (r *TraderRepository) UpdateOne(trader *repositories.Trader) {
	r.db.Save(&trader)
}

// DeleteOne удаляет торговца
func (r *TraderRepository) DeleteOne(trader *repositories.Trader) {
	r.db.Delete(&trader)
}
