package warrior_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-warrior]`,
		Level:  log.DebugLevel,
	})
)

type WarriorRepository struct {
	db gorm.DB
}

// InitWarriorRepository создает репозиторий
func InitWarriorRepository(db gorm.DB) repositories.IWarriorRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Warrior{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &WarriorRepository{
		db: db,
	}
}

// Create создает рабочего
func (r *WarriorRepository) Create(warrior *repositories.Warrior, villageID string, accountId string) {
	r.db.Create(&repositories.Warrior{
		ID:            uuid.NewString(),
		Location:      warrior.Location,
		MaxStorage:    100,
		Health:        100,
		UsedStorage:   0,
		FromDeparture: 0,
		ToArrival:     0,
		VillageID:     villageID,
		AccountID:     accountId,
	})
}

// GetOne возвращает первого рабочего, попавшего под условие
func (r *WarriorRepository) GetOne(id string, accountID string) repositories.Warrior {
	var warrior = repositories.Warrior{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&warrior)
	return warrior
}

// GetAll возвращает всех рабочих
func (r *WarriorRepository) GetAll(accountID string, q repositories.WarriorFindAll) []repositories.Warrior {
	var warrior = repositories.Warrior{
		AccountID:     accountID,
		UsedStorage:   *q.UsedStorage,
		MaxStorage:    *q.MaxStorage,
		Location:      *q.Location,
		FromDeparture: *q.FromDeparture,
		ToArrival:     *q.ToArrival,
	}
	var warriors = []repositories.Warrior{}

	limit := -1
	if q.Limit != nil {
		limit = *q.Limit
	}

	r.db.Model(warrior).Limit(limit).Find(&warriors)
	return warriors
}

// UpdateOne обновляет рабочего
func (r *WarriorRepository) UpdateOne(warrior *repositories.Warrior) {
	r.db.Save(&warrior)
}

// DeleteOne удаляет рабочего
func (r *WarriorRepository) DeleteOne(warrior *repositories.Warrior) {
	r.db.Delete(&warrior)
}
