package resource_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-resource]`,
		Level:  log.DebugLevel,
	})
)

type ResourceRepository struct {
	db gorm.DB
}

// InitResourceRepository создает репозиторий
func InitResourceRepository(db gorm.DB) repositories.IResourceRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Resource{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &ResourceRepository{
		db: db,
	}
}

// Create создает ресурс
func (r *ResourceRepository) Create(resource *repositories.Resource, placeID string) {
	r.db.Create(&repositories.Resource{
		ID:      uuid.NewString(),
		Type:    resource.Type,
		Amount:  resource.Amount,
		Weight:  resource.Weight,
		PlaceID: placeID,
	})
}

// GetOne возвращает первый ресурс, попавший под условие
func (r *ResourceRepository) GetOne(id string, placeID string) repositories.Resource {
	var resource = repositories.Resource{
		PlaceID: placeID,
		ID:      id,
	}
	r.db.First(&resource)
	return resource
}

// GetAll возвращает все ресурсы
func (r *ResourceRepository) GetAll(placeID string, q repositories.ResourceFindAll) []repositories.Resource {
	var resource = repositories.Resource{
		PlaceID: placeID,
		Type:    *q.Type,
		Amount:  *q.Amount,
		Weight:  *q.Weight,
	}
	var resources = []repositories.Resource{}

	limit := -1
	if q.Limit != nil {
		limit = *q.Limit
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
