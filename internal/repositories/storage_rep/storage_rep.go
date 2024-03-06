package storage_rep

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
		Prefix: `[database-storage]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IStorageRepository

type StorageRepository struct {
	db gorm.DB
}

// Init создает репозиторий хранилища
func Init() {
	Repository = &StorageRepository{
		db: *storage.Database,
	}
}

// Create создает хранилище
func (r *StorageRepository) Create(storage *repositories.Storage, accountId string) {
	r.db.Create(&repositories.Storage{
		ID:            uuid.NewString(),
		MaxDurability: 1000,
		Durability:    1000,
		Level:         1,
		MaxStorage:    10000,
		UsedStorage:   0,
		X:             storage.X,
		Y:             storage.Y,
		AccountID:     accountId,
	})
}

// GetOne возвращает первое хранилище, попавшее под условие
func (r *StorageRepository) GetOne(id string, accountID string) repositories.Storage {
	var storage = repositories.Storage{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&storage)
	return storage
}

// GetAll возвращает все хранилища
func (r *StorageRepository) GetAll(query repositories.StorageGetAll, accountID string) []repositories.Storage {
	var storage = repositories.Storage{
		MaxDurability: *query.MaxDurability,
		Durability:    *query.Durability,
		Level:         *query.Level,
		MaxStorage:    *query.MaxStorage,
		UsedStorage:   *query.UsedStorage,
		X:             *query.X,
		Y:             *query.Y,
		AccountID:     accountID,
	}
	var storages = []repositories.Storage{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(storage).Limit(limit).Find(&storages)
	return storages
}

// UpdateOne обновляет хранилище
func (r *StorageRepository) UpdateOne(storage *repositories.Storage) {
	r.db.Save(&storage)
}

// DeleteOne удаляет хранилище
func (r *StorageRepository) DeleteOne(storage *repositories.Storage) {
	r.db.Delete(&storage)
}
