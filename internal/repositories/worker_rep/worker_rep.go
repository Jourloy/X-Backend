package worker_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-worker]`,
		Level:  log.DebugLevel,
	})
)

type WorkerRepository struct {
	db gorm.DB
}

// InitWorkerRepository создает репозиторий
func InitWorkerRepository(db gorm.DB) repositories.IWorkerRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Worker{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &WorkerRepository{
		db: db,
	}
}

// Create создает объект в БД
func (r *WorkerRepository) Create(worker *repositories.Worker, colonyID string) {
	r.db.Create(&repositories.Worker{
		ID:         uuid.NewString(),
		Location:   worker.Location,
		MaxStorage: worker.MaxStorage,
		ColonyID:   colonyID,
	})
}

// GetOne возвращает первый объект, попавший под условие
func (r *WorkerRepository) GetOne(id string, colonyID string) repositories.Worker {
	var worker = repositories.Worker{
		ColonyID: colonyID,
		ID:       id,
	}
	r.db.First(&worker)
	return worker
}

// UpdateOne обновляет объект в БД
func (r *WorkerRepository) UpdateOne(worker *repositories.Worker) {
	r.db.Save(&worker)
}

// DeleteOne удаляет объект из БД
func (r *WorkerRepository) DeleteOne(worker *repositories.Worker) {
	r.db.Delete(&worker)
}
