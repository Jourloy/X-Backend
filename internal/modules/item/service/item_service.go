package item_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/item_rep"
	"github.com/jourloy/X-Backend/internal/repositories/trader_rep"
	"github.com/jourloy/X-Backend/internal/repositories/village_rep"
	"github.com/jourloy/X-Backend/internal/repositories/warrior_rep"
	"github.com/jourloy/X-Backend/internal/repositories/worker_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[item-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	iteRep repositories.IItemRepository
	warRep repositories.IWarriorRepository
	worRep repositories.IWorkerRepository
	vilRep repositories.IVillageRepository
	traRep repositories.ITraderRepository
	cache  redis.Client
}

// Init создает сервис вещи
func Init() *Service {

	iteRep := item_rep.Repository
	worRep := worker_rep.Repository
	warRep := warrior_rep.Repository
	traRep := trader_rep.Repository
	vilRep := village_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		iteRep: iteRep,
		worRep: worRep,
		warRep: warRep,
		traRep: traRep,
		vilRep: vilRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает вещь
func (s *Service) Create(b repositories.Item, pID string, pType string, aID string) createResp {
	if pType == `village` {
		v := s.vilRep.GetOne(pID, aID)
		if v.ID == `` {
			return createResp{Err: errors.New(`village not found`)}
		}
	} else if pType == `worker` {
		w := s.worRep.GetOne(pID, aID)
		if w.ID == `` {
			return createResp{Err: errors.New(`worker not found`)}
		}
	} else if pType == `warrior` {
		w := s.warRep.GetOne(pID, aID)
		if w.ID == `` {
			return createResp{Err: errors.New(`warrior not found`)}
		}
	} else if pType == `trader` {
		w := s.warRep.GetOne(pID, aID)
		if w.ID == `` {
			return createResp{Err: errors.New(`trader not found`)}
		}
	} else {
		return createResp{Err: errors.New(`pType is invalid`)}
	}

	s.iteRep.Create(&b, pID, aID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err  error
	Item repositories.Item
}

// GetOne получает вещь по id
func (s *Service) GetOne(id string, aID string) getOneResp {
	i := s.iteRep.GetOne(id, aID)
	return getOneResp{
		Err:  nil,
		Item: i,
	}
}

type getAllResp struct {
	Err   error
	Items []repositories.Item
}

// GetAll возвращает все вещи
func (s *Service) GetAll(q repositories.ItemFindAll, aID string) getAllResp {
	// Получение вещей
	items := s.iteRep.GetAll(q, aID)

	return getAllResp{
		Err:   nil,
		Items: items,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет вещь
func (s *Service) UpdateOne(b repositories.Item, aID string) updateOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = aID

	s.iteRep.UpdateOne(&b)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет вещь
func (s *Service) DeleteOne(b repositories.Item, aID string) deleteOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = aID

	s.iteRep.DeleteOne(&b)
	return deleteOneResp{
		Err: nil,
	}
}
