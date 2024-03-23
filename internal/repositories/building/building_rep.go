package building_rep

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
	building_templates "github.com/jourloy/X-Backend/internal/templates/buildings"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[building-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.BuildingRepository

type BuildingRepository struct {
	db gorm.DB
}

// Init создает репозиторий постройки
func Init() {
	go migration()

	Repository = &BuildingRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Building{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create cоздает постройку
func (r *BuildingRepository) Create(create *repositories.BuildingCreate) (*repositories.Building, error) {
	building := repositories.Building{
		X:         create.X,
		Y:         create.Y,
		Type:      create.Type,
		AccountID: create.AccountID,
		SectorID:  create.SectorID,
	}

	// Шаблон
	template := building_templates.BuildingTemplates[create.Type]

	building.MaxDurability = template.MaxDurability
	building.Durability = template.Durability
	building.MaxStorage = template.MaxStorage
	building.UsedStorage = template.UsedStorage
	building.Level = template.Level
	building.AttackRange = template.AttackRange
	building.CanTrade = template.CanTrade

	res := r.db.Create(&building)
	if res.Error != nil {
		return nil, res.Error
	}

	return &building, nil
}

// GetOne возвращает первую постройку, попавшую под условие
func (r *BuildingRepository) GetOne(query *repositories.BuildingGet) (*repositories.Building, error) {
	building := repositories.Building{}

	res := r.db.First(&building, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &building, nil
}

// GetAll возвращает все постройки, попавшие под условие
func (r *BuildingRepository) GetAll(query *repositories.BuildingGet) (*[]repositories.Building, error) {
	buildings := []repositories.Building{}

	res := r.db.Find(&buildings, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &buildings, nil
}

// UpdateOne обновляет постройку
func (r *BuildingRepository) UpdateOne(building *repositories.Building) error {
	res := r.db.Save(&building)
	return res.Error
}

// DeleteOne удаляет постройку
func (r *BuildingRepository) DeleteOne(building *repositories.Building) error {
	res := r.db.Delete(&building, building)
	return res.Error
}
