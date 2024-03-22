package storage

import (
	"github.com/charmbracelet/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/jourloy/X-Backend/internal/config"
)

var Database *gorm.DB

// InitDB подключается к базе данных
func InitDB() {
	db, err := gorm.Open(postgres.Open(config.DatabaseDSN), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(`Failed to connect database`)
	}
	Database = db
}
