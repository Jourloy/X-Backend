package item_rep

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
		Prefix: `[database-item]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IItemRepository

type ItemRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Item{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &ItemRepository{
		db: *storage.Database,
	}
}

// Create создает вещь
func (r *ItemRepository) Create(item *repositories.Item, parentID string, accountID string) {
	r.db.Create(&repositories.Item{
		ID: uuid.NewString(),

		// Добавить данные

	})
}

// GetOne возвращает первую вещь, попавшую под условие
func (r *ItemRepository) GetOne(id string, accountID string) repositories.Item {
	var item = repositories.Item{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&item)
	return item
}

// GetAll возвращает все вещи
func (r *ItemRepository) GetAll(query repositories.ItemGetAll, accountID string) []repositories.Item {
	var item = repositories.Item{

		// Добавить данные

		AccountID: accountID,
	}
	var items = []repositories.Item{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
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
