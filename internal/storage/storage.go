package storage

import (
	"os"
	"time"

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

	go runMigrations()
}

func runMigrations() {
	t := time.Now()
	// Автоматическая миграция
	if err := Database.AutoMigrate(
		&repositories.Account{},

		&repositories.Sector{},
		&repositories.Node{},
		&repositories.Deposit{},

		&repositories.Townhall{},
		&repositories.Storage{},
		&repositories.Tower{},
		&repositories.Wall{},
		&repositories.Market{},
		&repositories.Plan{},

		&repositories.Worker{},
		&repositories.Warrior{},
		&repositories.Trader{},
		&repositories.Scout{},

		&repositories.Item{},
		&repositories.Resource{},
	); err != nil {
		logger.Fatal(`Migration failed`)
	}
	logger.Debug(`Migration complete`, `latency`, time.Since(t))
}
