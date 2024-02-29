package cache

import (
	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/config"
)

var Client *redis.Client

// InitCache подключается к кэшу
func InitCache() {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.CacheDSN,
		Password: ``,
		DB:       0,
	})

	log.Info(`Cache connected`)
}
