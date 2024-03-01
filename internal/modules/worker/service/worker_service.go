package worker_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/village_rep"
	"github.com/jourloy/X-Backend/internal/repositories/worker_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[worker-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	worRep repositories.IWorkerRepository
	vilRep repositories.IVillageRepository
	cache  redis.Client
}

// Init создает сервис рабочего
func Init() *Service {

	worRep := worker_rep.Repository
	vilRep := village_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		worRep: worRep,
		vilRep: vilRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает рабочего
func (s *Service) Create(body repositories.Worker, accountID string, vID string) createResp {
	// Проверка существования поселения
	village := s.vilRep.GetOne(vID, accountID)
	if village.ID == `` {
		return createResp{Err: errors.New(`village not found`)}
	}

	s.worRep.Create(&body, vID, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err    error
	Worker repositories.Worker
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	worker := s.worRep.GetOne(id, accountID)
	return getOneResp{
		Err:    nil,
		Worker: worker,
	}
}

type getAllResp struct {
	Err     error
	Workers []repositories.Worker
}

func (s *Service) GetAll(query repositories.WorkerFindAll, accountID string) getAllResp {
	// Получение работников
	workers := s.worRep.GetAll(accountID, query)
	return getAllResp{
		Err:     nil,
		Workers: workers,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Worker, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.worRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Worker, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.worRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
