package storage

import (
	"os"

	"github.com/charmbracelet/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/config"
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
		Logger:                 nil,
	})
	if err != nil {
		log.Fatal(`Failed to connect database`)
	}
	Database = db
}
