package account_service

import (
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
)

type Service struct {
	aRep  repositories.IAccountRepository
	cache redis.Client
}

// Init создает сервис аккаунта
func Init() *Service {

	aRep := account_rep.Repository

	return &Service{
		aRep:  aRep,
		cache: *cache.Client,
	}
}

type createResp struct {
	Account *repositories.Account
	Err     error
}

// Create создает аккаунт
func (s *Service) Create(body repositories.AccountCreate) createResp {
	account, err := s.aRep.Create(&body)
	return createResp{
		Err:     err,
		Account: account,
	}
}

type getOneResp struct {
	Err     error
	Account repositories.Account
}

// GetOne получает аккаунт по id
func (s *Service) GetOne(accountID string) getOneResp {
	user := repositories.Account{ID: accountID}
	s.aRep.GetOne(&user)
	return getOneResp{
		Err:     nil,
		Account: user,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет аккаунт
func (s *Service) UpdateOne(body repositories.Account) updateOneResp {
	s.aRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет аккаунт
func (s *Service) DeleteOne(body repositories.Account) deleteOneResp {
	s.aRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
