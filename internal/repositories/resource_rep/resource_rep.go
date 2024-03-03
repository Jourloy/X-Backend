package resource_rep

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
		Prefix: `[database-resource]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IResourceRepository

type ResourceRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Resource{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &ResourceRepository{
		db: *storage.Database,
	}
}

// Create создает ресурс
func (r *ResourceRepository) Create(resource *repositories.Resource, sectorID string) {
	r.db.Create(&repositories.Resource{
		ID: uuid.NewString(),

		// Добавить данные

		SectorID: sectorID,
	})
}

// GetOne возвращает первый ресурс, попавший под условие
func (r *ResourceRepository) GetOne(id string, sectorID string) repositories.Resource {
	var resource = repositories.Resource{
		SectorID: sectorID,
		ID:       id,
	}
	r.db.First(&resource)
	return resource
}

// GetAll возвращает все ресурсы
func (r *ResourceRepository) GetAll(query repositories.ResourceGetAll, sectorID string) []repositories.Resource {
	var resource = repositories.Resource{

		// Добавить данные

		SectorID: sectorID,
	}
	var resources = []repositories.Resource{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(resource).Limit(limit).Find(&resources)
	return resources
}

// UpdateOne обновляет ресурс
func (r *ResourceRepository) UpdateOne(resource *repositories.Resource) {
	r.db.Save(&resource)
}

// DeleteOne удаляет ресурс
func (r *ResourceRepository) DeleteOne(resource *repositories.Resource) {
	r.db.Delete(&resource)
}
