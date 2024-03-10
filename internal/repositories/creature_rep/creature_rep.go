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

// Init создает репозиторий
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

	// Шаблоны
	switch create.Race {
	case `human`:
		requireFood := creature_templates.Human.RequireFood
		if create.IsWorker {
			requireFood += 0.2
		}
		if create.IsTrader {
			requireFood += 0.2
		}
		if create.IsWarrior {
			requireFood += 0.4
		}

		creature.MaxStorage = creature_templates.Human.MaxStorage
		creature.UsedStorage = creature_templates.Human.UsedStorage
		creature.FatiguePerStep = creature_templates.Human.FatiguePerStep
		creature.Fatigue = creature_templates.Human.Fatigue
		creature.MaxHealth = creature_templates.Human.MaxHealth
		creature.Health = creature_templates.Human.Health
		creature.RequireFood = requireFood
	case `swarm`:
		requireFood := creature_templates.Swarm.RequireFood
		if create.IsWorker {
			requireFood += 0.2
		}
		if create.IsTrader {
			requireFood += 0.2
		}
		if create.IsWarrior {
			requireFood += 0.4
		}

		creature.MaxStorage = creature_templates.Swarm.MaxStorage
		creature.UsedStorage = creature_templates.Swarm.UsedStorage
		creature.FatiguePerStep = creature_templates.Swarm.FatiguePerStep
		creature.Fatigue = creature_templates.Swarm.Fatigue
		creature.MaxHealth = creature_templates.Swarm.MaxHealth
		creature.Health = creature_templates.Swarm.Health
		creature.RequireFood = requireFood
	case `robot`:
		requireFood := creature_templates.Robot.RequireFood
		if create.IsWorker {
			requireFood += 0.2
		}
		if create.IsTrader {
			requireFood += 0.2
		}
		if create.IsWarrior {
			requireFood += 0.4
		}

		creature.MaxStorage = creature_templates.Robot.MaxStorage
		creature.UsedStorage = creature_templates.Robot.UsedStorage
		creature.FatiguePerStep = creature_templates.Robot.FatiguePerStep
		creature.Fatigue = creature_templates.Robot.Fatigue
		creature.MaxHealth = creature_templates.Robot.MaxHealth
		creature.Health = creature_templates.Robot.Health
		creature.RequireFood = requireFood
	}

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

	// Есои ошибка
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

	// Есои ошибка
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
