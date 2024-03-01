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

// Init создает репозиторий
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
func (r *traderRepository) Create(trader *repositories.Trader, villageID string, accountId string) {
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
func (r *traderRepository) GetOne(id string, accountID string) repositories.Trader {
	var trader = repositories.Trader{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&trader)
	return trader
}

// GetAll возвращает всех торговцев
func (r *traderRepository) GetAll(accountID string, q repositories.TraderFindAll) []repositories.Trader {
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
func (r *traderRepository) UpdateOne(trader *repositories.Trader) {
	r.db.Save(&trader)
}

// DeleteOne удаляет торговца
func (r *traderRepository) DeleteOne(trader *repositories.Trader) {
	r.db.Delete(&trader)
}
