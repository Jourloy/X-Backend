package account_rep

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
	Repository = &AccountRepository{
		db: *storage.Database,
	}
}

// Create создает аккаунт
func (r *AccountRepository) Create(create *repositories.AccountCreate) (*repositories.Account, error) {
	u := repositories.Account{Username: create.Username}
	if err := r.GetOne(&u); err != nil {
		return nil, err
	}

	if u.ID != `` {
		return nil, errors.New(`account already exist`)
	}

	user := repositories.Account{
		ID:       uuid.NewString(),
		ApiKey:   uuid.NewString(),
		Username: create.Username,
		Balance:  0,
	}

	res := r.db.Create(&user)
	return &user, res.Error
}

// GetOne возвращает первый аккаунт, попавший под условие
func (r *AccountRepository) GetOne(account *repositories.Account) error {
	res := r.db.First(&account, account)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return res.Error
}

// UpdateOne обновляет аккаунт
func (r *AccountRepository) UpdateOne(account *repositories.Account) {
	r.db.Save(&account)
}

// DeleteOne удаляет аккаунт
func (r *AccountRepository) DeleteOne(account *repositories.Account) {
	r.db.Delete(&account)
}
