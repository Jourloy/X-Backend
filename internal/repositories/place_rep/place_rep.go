package place_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-place]`,
		Level:  log.DebugLevel,
	})
)

type PlaceRepository struct {
	db gorm.DB
}

// InitPlaceRepository создает репозиторий
func InitPlaceRepository(db gorm.DB) repositories.IPlaceRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Place{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &PlaceRepository{
		db: db,
	}
}

// Create создает место
func (r *PlaceRepository) Create(place *repositories.Place) {
	r.db.Create(&repositories.Place{
		ID: uuid.NewString(),
	})
}

// GetOne возвращает первое место, попавшее под условие
func (r *PlaceRepository) GetOne(id string) repositories.Place {
	var place = repositories.Place{
		ID: id,
	}
	r.db.First(&place)
	return place
}

// UpdateOne обновляет место
func (r *PlaceRepository) UpdateOne(place *repositories.Place) {
	r.db.Save(&place)
}

// DeleteOne удаляет место
func (r *PlaceRepository) DeleteOne(place *repositories.Place) {
	r.db.Delete(&place)
}
