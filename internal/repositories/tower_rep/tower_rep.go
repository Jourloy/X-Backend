package tower_rep

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
		Prefix: `[database-tower]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.ITowerRepository

type TowerRepository struct {
	db gorm.DB
}

// Init создает репозиторий башни
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Tower{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &TowerRepository{
		db: *storage.Database,
	}
}

// Create создает башню
func (r *TowerRepository) Create(tower *repositories.Tower, accountId string) {
	r.db.Create(&repositories.Tower{
		ID:            uuid.NewString(),
		MaxDurability: 500,
		Durability:    500,
		Level:         1,
		MaxStorage:    500,
		UsedStorage:   0,
		X:             tower.X,
		Y:             tower.Y,
		AccountID:     accountId,
	})
}

// GetOne возвращает первую башню, попавшую под условие
func (r *TowerRepository) GetOne(id string, accountID string) repositories.Tower {
	var tower = repositories.Tower{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&tower)
	return tower
}

// GetAll возвращает все башни
func (r *TowerRepository) GetAll(query repositories.TowerGetAll, accountID string) []repositories.Tower {
	var tower = repositories.Tower{
		MaxDurability: *query.MaxDurability,
		Durability:    *query.Durability,
		Level:         *query.Level,
		MaxStorage:    *query.MaxStorage,
		UsedStorage:   *query.UsedStorage,
		X:             *query.X,
		Y:             *query.Y,
		AccountID:     accountID,
	}
	var towers = []repositories.Tower{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(tower).Limit(limit).Find(&towers)
	return towers
}

// UpdateOne обновляет башню
func (r *TowerRepository) UpdateOne(tower *repositories.Tower) {
	r.db.Save(&tower)
}

// DeleteOne удаляет башню
func (r *TowerRepository) DeleteOne(tower *repositories.Tower) {
	r.db.Delete(&tower)
}
