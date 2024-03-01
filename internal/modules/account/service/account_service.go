package account_service

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[account-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	aRep  repositories.IAccountRepository
	cache redis.Client
}

// InitAccountService создает сервис аккаунта
func InitAccountService() *Service {

	aRep := account_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		aRep:  aRep,
		cache: *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает аккаунт
func (s *Service) Create(b repositories.Account) createResp {
	s.aRep.Create(&b)
	return createResp{
		Err: nil,
	}
}

type getOneResp struct {
	Err     error
	Account repositories.Account
}

// GetOne получает аккаунт по id
func (s *Service) GetOne(aID string) getOneResp {
	a := s.aRep.GetOne(aID)
	return getOneResp{
		Err:     nil,
		Account: a,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет аккаунт
func (s *Service) UpdateOne(b repositories.Account) updateOneResp {
	s.aRep.UpdateOne(&b)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет аккаунт
func (s *Service) DeleteOne(b repositories.Account) deleteOneResp {
	s.aRep.DeleteOne(&b)
	return deleteOneResp{
		Err: nil,
	}
}
