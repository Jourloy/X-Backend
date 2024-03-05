package itemTemplate_rep

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
		Prefix: `[database-itemTemplate]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IItemTemplateRepository

type ItemTemplateRepository struct {
	db gorm.DB
}

// Init создает репозиторий шаблона предмета
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.ItemTemplate{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &ItemTemplateRepository{
		db: *storage.Database,
	}
}

// Create создает шаблон предмета
func (r *ItemTemplateRepository) Create(itemTemplate *repositories.ItemTemplate) {
	r.db.Create(&repositories.ItemTemplate{
		ID:   uuid.NewString(),
		Type: itemTemplate.Type,
	})
}

// GetOne возвращает первый шаблон предмета, попавший под условие
func (r *ItemTemplateRepository) GetOne(id string) repositories.ItemTemplate {
	var itemTemplate = repositories.ItemTemplate{
		ID: id,
	}
	r.db.First(&itemTemplate)
	return itemTemplate
}

// GetAll возвращает все шаблоны предмета
func (r *ItemTemplateRepository) GetAll(query repositories.ItemTemplateGetAll) []repositories.ItemTemplate {
	var itemTemplate = repositories.ItemTemplate{
		Type: *query.Type,
	}
	var itemTemplates = []repositories.ItemTemplate{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(itemTemplate).Limit(limit).Find(&itemTemplates)
	return itemTemplates
}

// UpdateOne обновляет шаблон предмета
func (r *ItemTemplateRepository) UpdateOne(itemTemplate *repositories.ItemTemplate) {
	r.db.Save(&itemTemplate)
}

// DeleteOne удаляет шаблон предмета
func (r *ItemTemplateRepository) DeleteOne(itemTemplate *repositories.ItemTemplate) {
	r.db.Delete(&itemTemplate)
}
