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

// Create создает объект в БД
func (r *ItemRepository) Create(item *repositories.Item, parentID string) {
	r.db.Create(&repositories.Item{
		ID:       uuid.NewString(),
		Type:     item.Type,
		ParentID: parentID,
	})
}

// GetOne возвращает первый объект, попавший под условие
func (r *ItemRepository) GetOne(id string, parentID string) repositories.Item {
	var item = repositories.Item{
		ParentID: parentID,
		ID:       id,
	}
	r.db.First(&item)
	return item
}

// UpdateOne обновляет объект в БД
func (r *ItemRepository) UpdateOne(item *repositories.Item) {
	r.db.Save(&item)
}

// DeleteOne удаляет объект из БД
func (r *ItemRepository) DeleteOne(item *repositories.Item) {
	r.db.Delete(&item)
}
