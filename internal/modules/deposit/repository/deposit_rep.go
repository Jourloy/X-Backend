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

// Init создает репозиторий залежей
func Init() {
	go migration()

	Repository = &DepositRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Deposit{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create создает залежь
func (r *DepositRepository) Create(deposit *repositories.Deposit) {
	deposit.ID = uuid.NewString()
	r.db.Create(&deposit)
}

// GetOne возвращает первую залежь, попавшую под условие
func (r *DepositRepository) GetOne(deposit *repositories.Deposit) {
	r.db.First(&deposit)
}

// GetAll возвращает все залежи
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

// UpdateOne обновляет залежь
func (r *DepositRepository) UpdateOne(deposit *repositories.Deposit) {
	r.db.Save(&deposit)
}

// DeleteOne удаляет залежь
func (r *DepositRepository) DeleteOne(deposit *repositories.Deposit) {
	r.db.Delete(&deposit)
}
