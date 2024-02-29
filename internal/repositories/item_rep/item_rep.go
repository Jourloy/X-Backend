package item_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-item]`,
		Level:  log.DebugLevel,
	})
)

type ItemRepository struct {
	db gorm.DB
}

// InitItemRepository создает репозиторий
func InitItemRepository(db gorm.DB) repositories.IItemRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Item{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &ItemRepository{
		db: db,
	}
}

// Create создает вещь
func (r *ItemRepository) Create(item *repositories.Item, parentID string) {
	r.db.Create(&repositories.Item{
		ID:       uuid.NewString(),
		Type:     item.Type,
		ParentID: parentID,
	})
}

// GetOne возвращает первую вещь, попавшую под условие
func (r *ItemRepository) GetOne(id string, parentID string) repositories.Item {
	var item = repositories.Item{
		ParentID: parentID,
		ID:       id,
	}
	r.db.First(&item)
	return item
}

// GetAll возвращает все вещи
func (r *ItemRepository) GetAll(q repositories.ItemFindAll) []repositories.Item {
	var item = repositories.Item{
		Type:     *q.Type,
		ParentID: *q.ParentID,
	}
	var items = []repositories.Item{}

	limit := -1
	if q.Limit != nil {
		limit = *q.Limit
	}

	r.db.Model(item).Limit(limit).Find(&items)
	return items
}

// UpdateOne обновляет вещь
func (r *ItemRepository) UpdateOne(item *repositories.Item) {
	r.db.Save(&item)
}

// DeleteOne удаляет вещь
func (r *ItemRepository) DeleteOne(item *repositories.Item) {
	r.db.Delete(&item)
}
