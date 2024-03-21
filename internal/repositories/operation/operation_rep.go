package operation_rep

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[operation-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.OperationRepository

type OperationRepository struct {
	db gorm.DB
}

// Init создает репозиторий операции
func Init() {
	go migration()

	Repository = &OperationRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Operation{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create cоздает операцию
func (r *OperationRepository) Create(create *repositories.OperationCreate) (*repositories.Operation, error) {
	operation := repositories.Operation{
		Price:      create.Price,
		Type:       create.Type,
		Amount:     create.Amount,
		Name:       create.Name,
		IsResource: create.IsResource,
		IsItem:     create.IsItem,
		BuildingID: create.BuildingID,
		SectorID:   create.SectorID,
		AccountID:  create.AccountID,
	}

	// ШАБЛОНЫ

	res := r.db.Create(&operation)
	if res.Error != nil {
		return nil, res.Error
	}

	return &operation, nil
}

// GetOne возвращает первую операцию, попавшую под условие
func (r *OperationRepository) GetOne(query *repositories.OperationGet) (*repositories.Operation, error) {
	operation := repositories.Operation{}

	res := r.db.First(&operation, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &operation, nil
}

// GetAll возвращает все операции, попавшие под условие
func (r *OperationRepository) GetAll(query *repositories.OperationGet) (*[]repositories.Operation, error) {
	operations := []repositories.Operation{}

	res := r.db.Find(&operations, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &operations, nil
}

// UpdateOne обновляет операцию
func (r *OperationRepository) UpdateOne(operation *repositories.Operation) error {
	res := r.db.Save(&operation)
	return res.Error
}

// DeleteOne удаляет операцию
func (r *OperationRepository) DeleteOne(operation *repositories.Operation) error {
	res := r.db.Delete(&operation, operation)
	return res.Error
}
