package worker_rep

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
		Prefix: `[database-worker]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IWorkerRepository

type WorkerRepository struct {
	db gorm.DB
}

// Init создает репозиторий рабочего
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Worker{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &WorkerRepository{
		db: *storage.Database,
	}
}

// Create создает рабочего
func (r *WorkerRepository) Create(worker *repositories.Worker, accountId string) {
	r.db.Create(&repositories.Worker{
		ID:           uuid.NewString(),
		MaxStorage:   worker.MaxStorage,
		UsedStorage:  worker.UsedStorage,
		X:            worker.X,
		Y:            worker.Y,
		MaxHealth:    worker.MaxHealth,
		Health:       worker.Health,
		RequireCoins: 0.5,
		RequireFood:  0.5,
		Fatigue:      0,
		AccountID:    accountId,
	})
}

// GetOne возвращает первого рабочего, попавшего под условие
func (r *WorkerRepository) GetOne(id string, accountID string) repositories.Worker {
	var worker = repositories.Worker{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&worker)
	return worker
}

// GetAll возвращает всех рабочих
func (r *WorkerRepository) GetAll(query repositories.WorkerGetAll, accountID string) []repositories.Worker {
	var worker = repositories.Worker{
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
	var workers = []repositories.Worker{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(worker).Limit(limit).Find(&workers)
	return workers
}

// UpdateOne обновляет рабочего
func (r *WorkerRepository) UpdateOne(worker *repositories.Worker) {
	r.db.Save(&worker)
}

// DeleteOne удаляет рабочего
func (r *WorkerRepository) DeleteOne(worker *repositories.Worker) {
	r.db.Delete(&worker)
}
