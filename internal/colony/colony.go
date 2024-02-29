package colony

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type ColonyService struct {
	db    gorm.DB
	cache redis.Client
}

// InitColonyService создает сервис колоний
func InitColonyService(db gorm.DB, cache redis.Client) *ColonyService {
	return &ColonyService{
		db:    db,
		cache: cache,
	}
}

func Create(c gin.Context) {

}
