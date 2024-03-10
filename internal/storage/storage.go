package storage

import (
	"github.com/charmbracelet/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jourloy/X-Backend/internal/config"
)

var Database *gorm.DB

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
