package deposit_rep

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
		Prefix: `[deposit-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IDepositRepository

type DepositRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Deposit{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &DepositRepository{
		db: *storage.Database,
	}
}

// Create создает ресурс
func (r *DepositRepository) Create(deposit *repositories.Deposit) {
	r.db.Create(&repositories.Resource{
		ID:       uuid.NewString(),
		Type:     deposit.Type,
		Amount:   deposit.Amount,
		SectorID: deposit.SectorID,
	})
}

// GetOne возвращает первый ресурс, попавший под условие
func (r *DepositRepository) GetOne(id string, sectorID string) repositories.Deposit {
	var deposit = repositories.Deposit{
		SectorID: sectorID,
		ID:       id,
	}
	r.db.First(&deposit)
	return deposit
}

// GetAll возвращает все ресурсы
func (r *DepositRepository) GetAll(query repositories.DepositGetAll, sectorID string) []repositories.Deposit {
	var deposit = repositories.Deposit{
		Type:     *query.Type,
		Amount:   *query.Amount,
		SectorID: sectorID,
	}
	var deposits = []repositories.Deposit{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(deposit).Limit(limit).Find(&deposits)
	return deposits
}

// UpdateOne обновляет ресурс
func (r *DepositRepository) UpdateOne(deposit *repositories.Deposit) {
	r.db.Save(&deposit)
}

// DeleteOne удаляет ресурс
func (r *DepositRepository) DeleteOne(deposit *repositories.Deposit) {
	r.db.Delete(&deposit)
}