package creature_rep

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
	creature_templates "github.com/jourloy/X-Backend/internal/templates/creatures"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[creature-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.CreatureRepository

type CreatureRepository struct {
	db gorm.DB
}

// Init создает репозиторий существа
func Init() {
	go migration()

	Repository = &CreatureRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Creature{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create cоздает существо
func (r *CreatureRepository) Create(create *repositories.CreatureCreate) (*repositories.Creature, error) {
	creature := repositories.Creature{
		X:         create.X,
		Y:         create.Y,
		Race:      create.Race,
		IsWorker:  create.IsWorker,
		IsWarrior: create.IsWarrior,
		IsTrader:  create.IsTrader,
		AccountID: create.AccountID,
		SectorID:  create.SectorID,
	}

	// Шаблон
	template := creature_templates.CreatureTemplate[create.Race]

	requireFood := template.RequireFood
	if create.IsWorker {
		requireFood += 0.2
	}
	if create.IsTrader {
		requireFood += 0.2
	}
	if create.IsWarrior {
		requireFood += 0.4
	}

	creature.MaxStorage = template.MaxStorage
	creature.UsedStorage = template.UsedStorage
	creature.FatiguePerStep = template.FatiguePerStep
	creature.Fatigue = template.Fatigue
	creature.MaxHealth = template.MaxHealth
	creature.Health = template.Health
	creature.RequireFood = requireFood

	res := r.db.Create(&creature)
	if res.Error != nil {
		return nil, res.Error
	}

	return &creature, nil
}

// GetOne возвращает первое существо, попавшее под условие
func (r *CreatureRepository) GetOne(query *repositories.CreatureGet) (*repositories.Creature, error) {
	creature := repositories.Creature{}

	res := r.db.First(&creature, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &creature, nil
}

// GetAll возвращает все существа, попавшие под условие
func (r *CreatureRepository) GetAll(query *repositories.CreatureGet) (*[]repositories.Creature, error) {
	creatures := []repositories.Creature{}

	res := r.db.Find(&creatures, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &creatures, nil
}

// UpdateOne обновляет существо
func (r *CreatureRepository) UpdateOne(creature *repositories.Creature) error {
	res := r.db.Save(&creature)
	return res.Error
}

// DeleteOne удаляет существо
func (r *CreatureRepository) DeleteOne(creature *repositories.Creature) error {
	res := r.db.Delete(&creature, creature)
	return res.Error
}
