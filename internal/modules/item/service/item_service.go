package item_service

import (
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	item_rep "github.com/jourloy/X-Backend/internal/repositories/item"
)

type Service struct {
	itemRep repositories.IItemRepository
	cache   redis.Client
}

// Init создает сервис предмета
func Init() *Service {
	go item_rep.Init()

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

type getOneResp struct {
	Err  error
	Item repositories.Item
}

// GetOne получает предмет по id
func (s *Service) GetOne(id string) getOneResp {
	item := repositories.Item{ID: id}
	s.itemRep.GetOne(&item)
	return getOneResp{
		Err:  nil,
		Item: item,
	}
}

type getAllResp struct {
	Err   error
	Items []repositories.Item
}

// GetAll возвращает все предметы
func (s *Service) GetAll(query repositories.ItemGetAll) getAllResp {
	// Получение предметов
	items := s.itemRep.GetAll(query)

	return getAllResp{
		Err:   nil,
		Items: items,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет предмет
func (s *Service) UpdateOne(body repositories.Item) updateOneResp {
	s.itemRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет предмет
func (s *Service) DeleteOne(body repositories.Item) deleteOneResp {
	s.itemRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
