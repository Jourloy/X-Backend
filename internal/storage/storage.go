package storage

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/charmbracelet/log"
	"github.com/jourloy/X-Backend/internal/config"
	"github.com/jourloy/X-Backend/internal/repositories"
)

var Database *gorm.DB

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[database]`,
		Level:  log.DebugLevel,
	})
)

// InitDB подключается к базе данных
func InitDB() {
	db, err := gorm.Open(postgres.Open(config.DatabaseDSN), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(`Failed to connect database`)
	}

	Database = db

	// Автоматическая миграция
	if err := Database.AutoMigrate(
		&repositories.Worker{},
		&repositories.Warrior{},
		&repositories.Wall{},
		&repositories.Trader{},
		&repositories.Townhall{},
		&repositories.Tower{},
		&repositories.Storage{},
		&repositories.Sector{},
		&repositories.Scout{},
		&repositories.ResourceTemplate{},
		&repositories.Resource{},
		&repositories.Plan{},
		&repositories.Market{},
		&repositories.ItemTemplate{},
		&repositories.Item{},
		&repositories.Deposit{},
		&repositories.Account{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}

	logger.Info(`Database connected`)
}
