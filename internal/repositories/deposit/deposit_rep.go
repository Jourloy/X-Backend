package deposit_rep

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
		Prefix: `[deposit-database]`,
		Level:  log.DebugLevel,
	})
)

var Repository repositories.DepositRepository

type DepositRepository struct {
	db gorm.DB
}

// Init создает репозиторий
func Init() {
	go migration()

	Repository = &DepositRepository{
		db: *storage.Database,
	}
}

func migration() {
	if err := storage.Database.AutoMigrate(
		&repositories.Deposit{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
}

// Create создает аккаунт
func (r *DepositRepository) Create(create *repositories.DepositCreate) (*repositories.Deposit, error) {
	// Проверка, есть ли уже такой аккаунт
	u, err := r.GetOne(&repositories.DepositGet{Username: &create.Username})
	if err != nil {
		return nil, err
	}

	// Если есть
	if u != nil {
		return nil, errors.New(`deposit already exist`)
	}

	// Создаем аккаунт
	user := repositories.Deposit{
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
func (r *DepositRepository) GetOne(query *repositories.DepositGet) (*repositories.Deposit, error) {
	deposit := repositories.Deposit{}

	res := r.db.First(&deposit, query)

	// Если ничего не нашли
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// Если ошибка
	if res.Error != nil {
		return nil, res.Error
	}

	return &deposit, nil
}

// UpdateOne обновляет аккаунт
func (r *DepositRepository) UpdateOne(deposit *repositories.Deposit) error {
	res := r.db.Save(&deposit)
	return res.Error
}

// DeleteOne удаляет аккаунт
func (r *DepositRepository) DeleteOne(deposit *repositories.Deposit) error {
	res := r.db.Delete(&deposit, deposit)
	return res.Error
}
