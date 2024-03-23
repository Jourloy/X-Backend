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
		Prefix: `[item-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IItemRepository

type ItemRepository struct {
	db gorm.DB
}

// Init создает репозиторий предмета
func Init() {
	go migration()

	Repository = &ItemRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Item{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create создает предмет
func (r *ItemRepository) Create(item *repositories.Item) {
	r.db.Create(&repositories.Item{
		ID:         uuid.NewString(),
		Type:       item.Type,
		X:          item.X,
		Y:          item.Y,
		ParentID:   item.ParentID,
		ParentType: item.ParentType,
		CreatorID:  item.CreatorID,
		SectorID:   item.SectorID,
	})
}

// GetOne возвращает первый пердмет, попавший под условие
func (r *ItemRepository) GetOne(item *repositories.Item) {
	r.db.First(&item, item)
}

// GetAll возвращает все предметы
func (r *ItemRepository) GetAll(query repositories.ItemGetAll) []repositories.Item {
	var item = repositories.Item{
		Type:       *query.Type,
		X:          *query.X,
		Y:          *query.Y,
		ParentID:   *query.ParentID,
		ParentType: *query.ParentType,
		CreatorID:  *query.CreatorID,
		SectorID:   *query.SectorID,
	}
	var items = []repositories.Item{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(item).Limit(limit).Find(&items)
	return items
}

// UpdateOne обновляет предмет
func (r *ItemRepository) UpdateOne(item *repositories.Item) {
	r.db.Save(&item)
}

// DeleteOne удаляет предмет
func (r *ItemRepository) DeleteOne(item *repositories.Item) {
	r.db.Delete(&item)
}
