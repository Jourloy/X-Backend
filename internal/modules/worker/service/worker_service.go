package worker_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[worker-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	wRep  repositories.IWorkerRepository
	cRep  repositories.IVillageRepository
	cache redis.Client
}

// InitWorkerService создает сервис рабочего
func InitWorkerService(wRep repositories.IWorkerRepository, cRep repositories.IVillageRepository, cache redis.Client) *Service {

	logger.Info(`Service initialized`)

	return &Service{
		wRep:  wRep,
		cRep:  cRep,
		cache: cache,
	}
}

type createResp struct {
	Err error
}

// Create создает рабочего
func (s *Service) Create(b repositories.Worker, aID string, vID string) createResp {
	// Проверка существования поселения
	village := s.cRep.GetOne(vID, aID)
	if village.ID == `` {
		return createResp{Err: errors.New(`village not found`)}
	}

	s.wRep.Create(&b, vID, aID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err    error
	Worker repositories.Worker
}

func (s *Service) GetOne(id string, aID string) getOneResp {
	worker := s.wRep.GetOne(id, aID)
	return getOneResp{
		Err:    nil,
		Worker: worker,
	}
}

type getAllResp struct {
	Err     error
	Workers []repositories.Worker
}

func (s *Service) GetAll(q repositories.WorkerFindAll, aID string) getAllResp {
	// Получение работников
	workers := s.wRep.GetAll(aID, q)
	return getAllResp{
		Err:     nil,
		Workers: workers,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(b repositories.Worker, aID string) updateOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = aID

	s.wRep.UpdateOne(&b)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(b repositories.Worker, aID string) deleteOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = aID

	s.wRep.DeleteOne(&b)
	return deleteOneResp{
		Err: nil,
	}
}
