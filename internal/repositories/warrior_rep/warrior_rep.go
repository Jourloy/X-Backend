package warrior_rep

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var Repository repositories.IWarriorRepository

type warriorRepository struct {
	db gorm.DB
}

// Init создает репозиторий воина
func Init() {
	Repository = &warriorRepository{
		db: *storage.Database,
	}
}

// Create создает воина
func (r *warriorRepository) Create(warrior *repositories.WarriorCreate) {
	r.db.Create(&repositories.Warrior{
		ID:           uuid.NewString(),
		MaxStorage:   10,
		UsedStorage:  0,
		X:            warrior.X,
		Y:            warrior.Y,
		MaxHealth:    100,
		Health:       100,
		RequireCoins: 0.5,
		RequireFood:  0.5,
		Fatigue:      0,
		AccountID:    warrior.AccountID,
	})
}

// GetOne возвращает первого воина, попавшего под условие
func (r *warriorRepository) GetOne(warrior *repositories.Warrior) {
	r.db.First(&warrior, warrior)
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
