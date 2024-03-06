package resourceTemplate_rep

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
		Prefix: `[database-resourceTemplate]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IResourceTemplateRepository

type ResourceTemplateRepository struct {
	db gorm.DB
}

// Init создает шаблон ресурсов
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.ResourceTemplate{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &ResourceTemplateRepository{
		db: *storage.Database,
	}
}

// Create создает шаблон ресурсов
func (r *ResourceTemplateRepository) Create(resourceTemplate *repositories.ResourceTemplate) {
	r.db.Create(&repositories.ResourceTemplate{
		ID:     uuid.NewString(),
		Type:   resourceTemplate.Type,
		Amount: resourceTemplate.Amount,
		Weight: resourceTemplate.Weight,
	})
}

// GetOne возвращает первый шаблон ресурсов, попавший под условие
func (r *ResourceTemplateRepository) GetOne(resourceTemplate repositories.ResourceTemplate) repositories.ResourceTemplate {
	r.db.First(&resourceTemplate)
	return resourceTemplate
}

// GetAll возвращает все шаблоны ресурсов
func (r *ResourceTemplateRepository) GetAll(query repositories.ResourceTemplateGetAll) []repositories.ResourceTemplate {
	var resourceTemplate = repositories.ResourceTemplate{
		Type:   *query.Type,
		Amount: *query.Amount,
		Weight: *query.Weight,
	}
	var resourceTemplates = []repositories.ResourceTemplate{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(resourceTemplate).Limit(limit).Find(&resourceTemplates)
	return resourceTemplates
}

// UpdateOne обновляет шаблон ресурсов
func (r *ResourceTemplateRepository) UpdateOne(resourceTemplate *repositories.ResourceTemplate) {
	r.db.Save(&resourceTemplate)
}

// DeleteOne удаляет шаблон ресурсов
func (r *ResourceTemplateRepository) DeleteOne(resourceTemplate *repositories.ResourceTemplate) {
	r.db.Delete(&resourceTemplate)
}
