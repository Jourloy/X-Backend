package wall_rep

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
		Prefix: `[database-wall]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IWallRepository

type WallRepository struct {
	db gorm.DB
}

// Init создает репозиторий стены
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Wall{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &WallRepository{
		db: *storage.Database,
	}
}

// Create создает стену
func (r *WallRepository) Create(wall *repositories.Wall, accountId string) {
	r.db.Create(&repositories.Wall{
		ID:            uuid.NewString(),
		MaxDurability: 1000,
		Durability:    1000,
		X:             wall.X,
		Y:             wall.Y,
		Level:         1,
		AccountID:     accountId,
	})
}

// GetOne возвращает первую стену, попавую под условие
func (r *WallRepository) GetOne(id string, accountID string) repositories.Wall {
	var wall = repositories.Wall{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&wall)
	return wall
}

// GetAll возвращает все стены
func (r *WallRepository) GetAll(query repositories.WallGetAll, accountID string) []repositories.Wall {
	var wall = repositories.Wall{
		MaxDurability: *query.MaxDurability,
		Durability:    *query.Durability,
		Level:         *query.Level,
		X:             *query.X,
		Y:             *query.Y,
		AccountID:     accountID,
	}
	var walls = []repositories.Wall{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(wall).Limit(limit).Find(&walls)
	return walls
}

// UpdateOne обновляет стену
func (r *WallRepository) UpdateOne(wall *repositories.Wall) {
	r.db.Save(&wall)
}

// DeleteOne удаляет стену
func (r *WallRepository) DeleteOne(wall *repositories.Wall) {
	r.db.Delete(&wall)
}
