package sector_rep

import (
	"errors"
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

var Repository repositories.SectorRepository

type sectorRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	go migration()

	Repository = &sectorRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Sector{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create создает сектор
func (r *sectorRepository) Create(create *repositories.SectorCreate) (*repositories.Sector, error) {
	sec, err := r.GetOne(&repositories.SectorGet{X: &create.X, Y: &create.Y})
	if err != nil {
		return nil, err
	}

	if sec != nil {
		return nil, errors.New(`sector already exist`)
	}

	sector := repositories.Sector{
		ID: uuid.NewString(),
		X:  create.X,
		Y:  create.Y,
	}

	res := r.db.Create(&sector)
	if res.Error != nil {
		return nil, res.Error
	}

	return &sector, nil
}

// GetOne возвращает первый сектор, попавший под условие
func (r *sectorRepository) GetOne(query *repositories.SectorGet) (*repositories.Sector, error) {
	sector := &repositories.Sector{}

	// Поиск
	res := r.db.
		Preload(`Nodes`).
		Preload(`Deposits`).
		First(sector, query)

	// Если ничего не нашлось
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return sector, nil
}

// GetAll возвращает все сектора
func (r *sectorRepository) GetAll(query *repositories.SectorGet) (*[]repositories.Sector, error) {
	var sector = repositories.Sector{}
	var sectors = []repositories.Sector{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	res := r.db.Model(sector).Preload(`Nodes`).Preload(`Deposits`).Limit(limit).Find(&sectors, query)

	// Если ничего не нашлось
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &sectors, nil
}

// UpdateOne обновляет сектор
func (r *sectorRepository) UpdateOne(sector *repositories.Sector) error {
	res := r.db.Save(&sector)
	return res.Error
}

// DeleteOne удаляет сектор
func (r *sectorRepository) DeleteOne(sector *repositories.Sector) error {
	res := r.db.Delete(&sector)
	return res.Error
}
