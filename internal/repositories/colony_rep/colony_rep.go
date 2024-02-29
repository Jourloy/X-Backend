package colony_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-colony]`,
		Level:  log.DebugLevel,
	})
)

type ColonyRepository struct {
	db gorm.DB
}

// InitColonyRepository создает репозиторий
func InitColonyRepository(db gorm.DB) repositories.IColonyRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Colony{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &ColonyRepository{
		db: db,
	}
}

// Create создает колонию
func (r *ColonyRepository) Create(colony *repositories.Colony, accountID string, placeID string) {
	r.db.Create(&repositories.Colony{
		ID:         uuid.NewString(),
		Balance:    colony.Balance,
		MaxStorage: colony.MaxStorage,
		AccountID:  accountID,
		PlaceID:    placeID,
	})
}

// GetOne возвращает первую колонию, попавшую под условие
func (r *ColonyRepository) GetOne(id string, accountID string) repositories.Colony {
	var colony = repositories.Colony{
		AccountID: accountID,
		ID:        id,
	}
	r.db.First(&colony)
	return colony
}

// GetAll возвращает все колонии
func (r *ColonyRepository) GetAll(accountID string, q repositories.ColonyFindAll) []repositories.Colony {
	var colony = repositories.Colony{
		AccountID:   accountID,
		Balance:     *q.Balance,
		MaxStorage:  *q.MaxStorage,
		UsedStorage: *q.UsedStorage,
	}
	var colonys = []repositories.Colony{}

	limit := -1
	if q.Limit != nil {
		limit = *q.Limit
	}

	r.db.Model(colony).Limit(limit).Find(&colonys)
	return colonys
}

// UpdateOne обновляет колонию
func (r *ColonyRepository) UpdateOne(colony *repositories.Colony) {
	r.db.Save(&colony)
}

// DeleteOne удаляет колонию
func (r *ColonyRepository) DeleteOne(colony *repositories.Colony) {
	r.db.Delete(&colony)
}
