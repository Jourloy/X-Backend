package warrior_rep

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
		Prefix: `[database-warrior]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IWarriorRepository

type warriorRepository struct {
	db gorm.DB
}

// Init создает репозиторий воина
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Warrior{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &warriorRepository{
		db: *storage.Database,
	}
}

// Create создает воина
func (r *warriorRepository) Create(warrior *repositories.Warrior, accountId string) {
	r.db.Create(&repositories.Warrior{
		ID:           uuid.NewString(),
		MaxStorage:   warrior.MaxStorage,
		UsedStorage:  warrior.UsedStorage,
		X:            warrior.X,
		Y:            warrior.Y,
		MaxHealth:    warrior.MaxHealth,
		Health:       warrior.Health,
		RequireCoins: 0.5,
		RequireFood:  0.5,
		Fatigue:      0,
		AccountID:    accountId,
	})
}

// GetOne возвращает первого воина, попавшего под условие
func (r *warriorRepository) GetOne(id string, accountID string) repositories.Warrior {
	var warrior = repositories.Warrior{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&warrior)
	return warrior
}

// GetAll возвращает всех воинов
func (r *warriorRepository) GetAll(query repositories.WarriorGetAll, accountID string) []repositories.Warrior {
	var warrior = repositories.Warrior{
		MaxStorage:   *query.MaxStorage,
		UsedStorage:  *query.UsedStorage,
		X:            *query.X,
		Y:            *query.Y,
		MaxHealth:    *query.MaxHealth,
		Health:       *query.Health,
		RequireCoins: *query.RequireCoins,
		RequireFood:  *query.RequireFood,
		Fatigue:      *query.Fatigue,
		AccountID:    accountID,
	}
	var warriors = []repositories.Warrior{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(warrior).Limit(limit).Find(&warriors)
	return warriors
}

// UpdateOne обновляет воина
func (r *warriorRepository) UpdateOne(warrior *repositories.Warrior) {
	r.db.Save(&warrior)
}

// DeleteOne удаляет воина
func (r *warriorRepository) DeleteOne(warrior *repositories.Warrior) {
	r.db.Delete(&warrior)
}
