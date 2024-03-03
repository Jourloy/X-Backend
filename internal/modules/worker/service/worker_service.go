package worker_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
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
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис рабочего
func Init() *Service {

	worRep := worker_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		worRep: worRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает рабочего
func (s *Service) Create(body repositories.Worker, accountID string) createResp {
	// Проверка существования аккаунта
	account := s.accRep.GetOne(accountID)
	if account.ID == `` {
		return createResp{Err: errors.New(`account not found`)}
	}

	// Проверка существования сектора
	sector := s.secRep.GetOne(body.SectorID)
	if sector.ID == `` {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.worRep.Create(&body, accountID)
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

func (s *Service) GetAll(query repositories.WorkerGetAll, accountID string) getAllResp {
	// Получение работников
	workers := s.worRep.GetAll(query, accountID)
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
