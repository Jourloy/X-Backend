package item_template_rep

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var Repository repositories.IItemTemplateRepository

type ItemTemplateRepository struct {
	db gorm.DB
}

// Init создает репозиторий шаблона предмета
func Init() {
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
func (r *ItemTemplateRepository) GetOne(itemTemplate *repositories.ItemTemplate) {
	r.db.First(&itemTemplate)
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
