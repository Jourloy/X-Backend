package resource_rep

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var Repository repositories.IResourceRepository

type ResourceRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	Repository = &ResourceRepository{
		db: *storage.Database,
	}
}

// Create создает ресурс
func (r *ResourceRepository) Create(resource *repositories.Resource) {
	resource.ID = uuid.NewString()
	r.db.Create(&resource)
}

// GetOne возвращает первый ресурс, попавший под условие
func (r *ResourceRepository) GetOne(resource repositories.Resource) {
	r.db.First(&resource)
}

// GetAll возвращает все ресурсы
func (r *ResourceRepository) GetAll(query repositories.ResourceGetAll) []repositories.Resource {
	resource := repositories.Resource{}
	resources := []repositories.Resource{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(resource).Limit(limit).Find(&resources, query)
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
