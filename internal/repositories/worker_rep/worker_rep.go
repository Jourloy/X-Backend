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

// Init создает репозиторий
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
func (r *WorkerRepository) Create(worker *repositories.Worker, villageID string, accountId string) {
	r.db.Create(&repositories.Worker{
		ID:            uuid.NewString(),
		Location:      worker.Location,
		MaxStorage:    100,
		UsedStorage:   0,
		FromDeparture: 0,
		ToArrival:     0,
		VillageID:     villageID,
		AccountID:     accountId,
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
func (r *WorkerRepository) GetAll(accountID string, q repositories.WorkerFindAll) []repositories.Worker {
	var worker = repositories.Worker{
		AccountID:     accountID,
		UsedStorage:   *q.UsedStorage,
		MaxStorage:    *q.MaxStorage,
		Location:      *q.Location,
		FromDeparture: *q.FromDeparture,
		ToArrival:     *q.ToArrival,
	}
	var workers = []repositories.Worker{}

	limit := -1
	if q.Limit != nil {
		limit = *q.Limit
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
