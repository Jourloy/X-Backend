package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/charmbracelet/log"
	"github.com/jourloy/X-Backend/internal/config"
)

var Database *gorm.DB

// InitDB подключается к базе данных
func InitDB() {
	db, err := gorm.Open(postgres.Open(config.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic(`failed to connect database`)
	}

	Database = db

	log.Info(`Database connected`)
}
