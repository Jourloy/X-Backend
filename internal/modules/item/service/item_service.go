package item_service

import (
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/item_rep"
	"github.com/jourloy/X-Backend/internal/repositories/trader_rep"
	"github.com/jourloy/X-Backend/internal/repositories/warrior_rep"
	"github.com/jourloy/X-Backend/internal/repositories/worker_rep"
)

type Service struct {
	iteRep repositories.IItemRepository
	warRep repositories.IWarriorRepository
	worRep repositories.IWorkerRepository
	traRep repositories.ITraderRepository
	cache  redis.Client
}

// Init создает сервис предмета
func Init() *Service {

	iteRep := item_rep.Repository
	worRep := worker_rep.Repository
	warRep := warrior_rep.Repository
	traRep := trader_rep.Repository

	return &Service{
		iteRep: iteRep,
		worRep: worRep,
		warRep: warRep,
		traRep: traRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает предмет
func (s *Service) Create(body repositories.Item, parentID string, parentType string, accountID string) createResp {
	/* if parentType == `worker` {
		w := s.worRep.GetOne(parentID, accountID)
		if w.ID == `` {
			return createResp{Err: errors.New(`worker not found`)}
		}
	} else if parentType == `warrior` {
		w := s.warRep.GetOne(parentID, accountID)
		if w.ID == `` {
			return createResp{Err: errors.New(`warrior not found`)}
		}
	} else if parentType == `trader` {
		w := s.warRep.GetOne(parentID, accountID)
		if w.ID == `` {
			return createResp{Err: errors.New(`trader not found`)}
		}
	} else {
		return createResp{Err: errors.New(`parentType is invalid`)}
	} */

	s.iteRep.Create(&body)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err  error
	Item repositories.Item
}

// GetOne получает предмет по id
func (s *Service) GetOne(id string) getOneResp {
	item := repositories.Item{ID: id}
	s.iteRep.GetOne(&item)
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
	items := s.iteRep.GetAll(query)

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
	s.iteRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет предмет
func (s *Service) DeleteOne(body repositories.Item) deleteOneResp {
	s.iteRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
