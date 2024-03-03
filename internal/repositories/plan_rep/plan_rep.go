package plan_rep

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
		Prefix: `[database-plan]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IPlanRepository

type planRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Plan{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &planRepository{
		db: *storage.Database,
	}
}

// Create создает рабочего
func (r *planRepository) Create(plan *repositories.Plan, accountId string) {
	r.db.Create(&repositories.Plan{
		ID:          uuid.NewString(),
		MaxProgress: plan.MaxProgress,
		Progress:    plan.Progress,
		X:           plan.X,
		Y:           plan.Y,
		Type:        plan.Type,
		AccountID:   accountId,
	})
}

// GetOne возвращает первого рабочего, попавшего под условие
func (r *planRepository) GetOne(id string, accountID string) repositories.Plan {
	var plan = repositories.Plan{
		ID:        id,
		AccountID: accountID,
	}
	r.db.First(&plan)
	return plan
}

// GetAll возвращает всех рабочих
func (r *planRepository) GetAll(query repositories.PlanGetAll, accountID string) []repositories.Plan {
	var plan = repositories.Plan{
		MaxProgress: *query.MaxProgress,
		Progress:    *query.Progress,
		Type:        *query.Type,
		Y:           *query.Y,
		X:           *query.X,
		AccountID:   accountID,
	}
	var plans = []repositories.Plan{}

	limit := -1
	if query.Limit != nil {
		limit = *query.Limit
	}

	r.db.Model(plan).Limit(limit).Find(&plans)
	return plans
}

// UpdateOne обновляет рабочего
func (r *planRepository) UpdateOne(plan *repositories.Plan) {
	r.db.Save(&plan)
}

// DeleteOne удаляет рабочего
func (r *planRepository) DeleteOne(plan *repositories.Plan) {
	r.db.Delete(&plan)
}