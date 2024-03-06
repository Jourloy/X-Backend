package account_rep

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/storage"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database-account]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.IAccountRepository

type AccountRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	// Автоматическая миграция
	if err := storage.Database.AutoMigrate(&repositories.Account{}); err != nil {
		logger.Fatal(`Migration failed`)
	}

	Repository = &AccountRepository{
		db: *storage.Database,
	}
}

// Create создает аккаунт
func (r *AccountRepository) Create(create *repositories.AccountCreate) (repositories.Account, error) {
	u := r.GetOne(repositories.Account{Username: create.Username})
	fmt.Println(u)

	user := repositories.Account{
		ID:       uuid.NewString(),
		ApiKey:   uuid.NewString(),
		Username: create.Username,
		Balance:  0,
	}
	res := r.db.Create(&user)
	return user, res.Error
}

// GetOne возвращает первый аккаунт, попавший под условие
func (r *AccountRepository) GetOne(account repositories.Account) repositories.Account {
	r.db.First(&account)
	return account
}

// UpdateOne обновляет аккаунт
func (r *AccountRepository) UpdateOne(account *repositories.Account) {
	r.db.Save(&account)
}

// DeleteOne удаляет аккаунт
func (r *AccountRepository) DeleteOne(account *repositories.Account) {
	r.db.Delete(&account)
}
