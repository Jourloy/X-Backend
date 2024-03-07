package sector_rep

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
		Prefix: `[sector-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.ISectorRepository

type sectorRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	Repository = &sectorRepository{
		db: *storage.Database,
	}
}

// Create создает сектор
func (r *sectorRepository) Create(sector *repositories.Sector) {
	r.db.Create(&repositories.Sector{
		ID: uuid.NewString(),
		X:  sector.X,
		Y:  sector.Y,
	})
}

// GetOne возвращает сектор по его ID
func (r *sectorRepository) GetOne(sector *repositories.Sector) {
	r.db.First(&sector, sector)
}

// GetAll возвращает все сектора
func (r *sectorRepository) GetAll(query repositories.SectorGetAll) []repositories.Sector {
	var sector = repositories.Sector{
		X: *query.X,
		Y: *query.Y,
	}

	var sectors = []repositories.Sector{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(sector).Limit(limit).Find(&sectors)
	return sectors
}

// UpdateOne обновляет сектор
func (r *sectorRepository) UpdateOne(sector *repositories.Sector) {
	r.db.Save(&sector)
}

// DeleteOne удаляет сектор
func (r *sectorRepository) DeleteOne(sector *repositories.Sector) {
	r.db.Delete(&sector)
}
