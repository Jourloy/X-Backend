package townhall_rep

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
		Prefix: `[database-townhall]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.ITownhallRepository

type TownhallRepository struct {
	db gorm.DB
}

// Init создает репозиторий главного здания
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Townhall{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &TownhallRepository{
		db: *storage.Database,
	}
}

// Create создает главное здание
func (r *TownhallRepository) Create(townhall *repositories.Townhall, accountId string) {
	r.db.Create(&repositories.Townhall{
		ID:            uuid.NewString(),
		MaxDurability: townhall.MaxDurability,
		Durability:    townhall.Durability,
		MaxStorage:    townhall.MaxStorage,
		UsedStorage:   townhall.UsedStorage,
		X:             townhall.X,
		Y:             townhall.Y,
		AccountID:     accountId,
	})
}

// GetOne возвращает первое главное здание, попавшее под условие
func (r *TownhallRepository) GetOne(id string, accountID string) repositories.Townhall {
	var townhall = repositories.Townhall{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&townhall)
	return townhall
}

// GetAll возвращает все главные здания
func (r *TownhallRepository) GetAll(query repositories.TownhallGetAll, accountID string) []repositories.Townhall {
	var townhall = repositories.Townhall{
		MaxDurability: *query.MaxDurability,
		Durability:    *query.Durability,
		MaxStorage:    *query.MaxStorage,
		UsedStorage:   *query.UsedStorage,
		X:             *query.X,
		Y:             *query.Y,
		AccountID:     accountID,
	}
	var townhalls = []repositories.Townhall{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(townhall).Limit(limit).Find(&townhalls)
	return townhalls
}

// UpdateOne обновляет главное здание
func (r *TownhallRepository) UpdateOne(townhall *repositories.Townhall) {
	r.db.Save(&townhall)
}

// DeleteOne удаляет главное здание
func (r *TownhallRepository) DeleteOne(townhall *repositories.Townhall) {
	r.db.Delete(&townhall)
}
