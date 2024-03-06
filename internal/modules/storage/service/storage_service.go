package storage_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/storage_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[storage-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	stoRep repositories.IStorageRepository
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис хранилища
func Init() *Service {

	stoRep := storage_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		stoRep: stoRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает хранилище
func (s *Service) Create(body repositories.Storage, accountID string) createResp {
	// Проверка существования аккаунта
	account := s.accRep.GetOne(repositories.Account{ID: accountID})
	if account.ID == `` {
		return createResp{Err: errors.New(`account not found`)}
	}

	// Проверка существования сектора
	sector := s.secRep.GetOne(body.SectorID)
	if sector.ID == `` {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.stoRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err     error
	Storage repositories.Storage
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	storage := s.stoRep.GetOne(id, accountID)
	return getOneResp{
		Err:     nil,
		Storage: storage,
	}
}

type getAllResp struct {
	Err      error
	Storages []repositories.Storage
}

func (s *Service) GetAll(query repositories.StorageGetAll, accountID string) getAllResp {
	// Получение хранилищ
	storages := s.stoRep.GetAll(query, accountID)
	return getAllResp{
		Err:      nil,
		Storages: storages,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Storage, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.stoRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Storage, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.stoRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
