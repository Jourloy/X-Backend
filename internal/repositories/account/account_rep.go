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
		Prefix: `[account-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.AccountRepository

type AccountRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	go migration()

	Repository = &AccountRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Account{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create создает аккаунт
func (r *AccountRepository) Create(create *repositories.AccountCreate) (*repositories.Account, error) {
	// Проверка, есть ли уже такой аккаунт
	u, err := r.GetOne(&repositories.AccountGet{Username: &create.Username})
	if err != nil {
		return nil, err
	}

	// Если есть
	if u != nil {
		return nil, errors.New(`account already exist`)
	}

	// Создаем аккаунт
	user := repositories.Account{
		ID:       uuid.NewString(),
		ApiKey:   uuid.NewString(),
		Username: create.Username,
		Race:     create.Race,
		Balance:  0,
		IsAdmin:  false,
	}

	res := r.db.Create(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}

// GetOne возвращает первый аккаунт, попавший под условие
func (r *AccountRepository) GetOne(query *repositories.AccountGet) (*repositories.Account, error) {
	account := repositories.Account{}

	res := r.db.First(&account, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &account, nil
}

// UpdateOne обновляет аккаунт
func (r *AccountRepository) UpdateOne(account *repositories.Account) error {
	res := r.db.Save(&account)
	return res.Error
}

// DeleteOne удаляет аккаунт
func (r *AccountRepository) DeleteOne(account *repositories.Account) error {
	res := r.db.Delete(&account, account)
	return res.Error
}
