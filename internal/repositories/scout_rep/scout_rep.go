package scout_rep

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
		Prefix: `[database-scout]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IScoutRepository

type ScoutRepository struct {
	db gorm.DB
}

// Init создает репозиторий разведчика
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Scout{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &ScoutRepository{
		db: *storage.Database,
	}
}

// Create создает разведчика
func (r *ScoutRepository) Create(scout *repositories.Scout, accountId string) {
	r.db.Create(&repositories.Scout{
		ID:           uuid.NewString(),
		MaxStorage:   scout.MaxStorage,
		UsedStorage:  scout.UsedStorage,
		X:            scout.X,
		Y:            scout.Y,
		MaxHealth:    scout.MaxHealth,
		Health:       scout.Health,
		RequireCoins: 0.5,
		RequireFood:  0.5,
		Fatigue:      0,
		AccountID:    accountId,
	})
}

// GetOne возвращает первого разведчика, попавшего под условие
func (r *ScoutRepository) GetOne(id string, accountID string) repositories.Scout {
	var scout = repositories.Scout{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&scout)
	return scout
}

// GetAll возвращает всех разведчиков
func (r *ScoutRepository) GetAll(query repositories.ScoutGetAll, accountID string) []repositories.Scout {
	var scout = repositories.Scout{
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
	var scouts = []repositories.Scout{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(scout).Limit(limit).Find(&scouts)
	return scouts
}

// UpdateOne обновляет разведчика
func (r *ScoutRepository) UpdateOne(scout *repositories.Scout) {
	r.db.Save(&scout)
}

// DeleteOne удаляет разведчика
func (r *ScoutRepository) DeleteOne(scout *repositories.Scout) {
	r.db.Delete(&scout)
}
