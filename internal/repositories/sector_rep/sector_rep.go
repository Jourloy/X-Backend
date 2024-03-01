package sector_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-sector]`,
		Level:  log.DebugLevel,
	})
)

type SectorRepository struct {
	db gorm.DB
}

// InitSectorRepository создает репозиторий
func InitSectorRepository(db gorm.DB) repositories.ISectorRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Sector{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &SectorRepository{
		db: db,
	}
}

// Create создает место
func (r *SectorRepository) Create(sector *repositories.Sector) {
	r.db.Create(&repositories.Sector{
		ID: uuid.NewString(),
	})
}

// GetOne возвращает первое место, попавшее под условие
func (r *SectorRepository) GetOne(id string) repositories.Sector {
	var sector = repositories.Sector{
		ID: id,
	}
	r.db.First(&sector)
	return sector
}

// GetAll возвращает все места
func (r *SectorRepository) GetAll(q repositories.SectorFindAll) []repositories.Sector {
	var sector = repositories.Sector{}
	var sectors = []repositories.Sector{}

	limit := -1
	if q.Limit != nil {
		limit = *q.Limit
	}

	r.db.Model(sector).Limit(limit).Find(&sectors)
	return sectors
}

// UpdateOne обновляет место
func (r *SectorRepository) UpdateOne(sector *repositories.Sector) {
	r.db.Save(&sector)
}

// DeleteOne удаляет место
func (r *SectorRepository) DeleteOne(sector *repositories.Sector) {
	r.db.Delete(&sector)
}
