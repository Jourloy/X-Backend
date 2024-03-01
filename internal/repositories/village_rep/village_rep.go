package village_rep

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-village]`,
		Level:  log.DebugLevel,
	})
)

type VillageRepository struct {
	db gorm.DB
}

// InitVillageRepository создает репозиторий
func InitVillageRepository(db gorm.DB) repositories.IVillageRepository {
	// Автоматическая миграция
	if err := db.AutoMigrate(&repositories.Village{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	return &VillageRepository{
		db: db,
	}
}

// Create создает поселение
func (r *VillageRepository) Create(village *repositories.Village, accountID string, sectorID string) {
	r.db.Create(&repositories.Village{
		ID:         uuid.NewString(),
		Balance:    village.Balance,
		MaxStorage: 1000000,
		AccountID:  accountID,
		SectorID:   sectorID,
	})
}

// GetOne возвращает первое поселение, попавшее под условие
func (r *VillageRepository) GetOne(id string, accountID string) repositories.Village {
	var village = repositories.Village{
		AccountID: accountID,
		ID:        id,
	}
	r.db.First(&village)
	return village
}

// GetAll возвращает все поселения
func (r *VillageRepository) GetAll(accountID string, q repositories.VillageFindAll) []repositories.Village {
	var village = repositories.Village{
		AccountID:   accountID,
		Balance:     *q.Balance,
		MaxStorage:  *q.MaxStorage,
		UsedStorage: *q.UsedStorage,
	}
	var villages = []repositories.Village{}

	limit := -1
	if q.Limit != nil {
		limit = *q.Limit
	}

	r.db.Model(village).Limit(limit).Find(&villages)
	return villages
}

// UpdateOne обновляет поселение
func (r *VillageRepository) UpdateOne(village *repositories.Village) {
	r.db.Save(&village)
}

// DeleteOne удаляет поселение
func (r *VillageRepository) DeleteOne(village *repositories.Village) {
	r.db.Delete(&village)
}
