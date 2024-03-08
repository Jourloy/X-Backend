package sector_rep

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
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
	sector.ID = uuid.NewString()
	r.db.Create(&sector)
}

// GetOne возвращает сектор по его ID
func (r *sectorRepository) GetOne(sector *repositories.Sector) {
	r.db.Preload(`Nodes`).Preload(`Deposits`).First(&sector, sector)
}

// GetAll возвращает все сектора
func (r *sectorRepository) GetAll(query repositories.SectorGetAll) []repositories.Sector {
	var sector = repositories.Sector{}
	var sectors = []repositories.Sector{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(sector).Preload(`Nodes`).Preload(`Deposits`).Limit(limit).Find(&sectors, query)
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
