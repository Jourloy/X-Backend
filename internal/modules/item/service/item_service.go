package item_service

import (
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	item_rep "github.com/jourloy/X-Backend/internal/repositories/item"
)

type Service struct {
	itemRep repositories.ItemRepository
	cache   redis.Client
}

// Init создает сервис предмета
func Init() *Service {
	itemRep := item_rep.Repository

	return &Service{
		itemRep: itemRep,
		cache:   *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает предмет
func (s *Service) Create(body repositories.Item, parentID string, parentType string, accountID string) createResp {
	s.itemRep.Create(&body)
	return createResp{Err: nil}
}
