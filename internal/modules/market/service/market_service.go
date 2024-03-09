package market_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/market_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
)

type Service struct {
	marRep repositories.IMarketRepository
	secRep repositories.SectorRepository
	accRep repositories.AccountRepository
	cache  redis.Client
}

// Init создает сервис рынка
func Init() *Service {

	marRep := market_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	return &Service{
		marRep: marRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает рынок
func (s *Service) Create(body repositories.MarketCreate) createResp {
	// Проверка существования аккаунта
	account, err := s.accRep.GetOne(&repositories.AccountGet{ID: &body.AccountID})
	if err != nil {
		return createResp{Err: err}
	}
	if account == nil {
		return createResp{Err: errors.New(`account not found`)}
	}

	// Проверка существования сектора
	sector, err := s.secRep.GetOne(&repositories.SectorGet{ID: &body.SectorID})
	if err != nil {
		return createResp{Err: err}
	}

	if sector == nil {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.marRep.Create(&body)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err    error
	Market repositories.Market
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	market := repositories.Market{ID: id, AccountID: accountID}
	s.marRep.GetOne(&market)
	return getOneResp{
		Err:    nil,
		Market: market,
	}
}

type getAllResp struct {
	Err     error
	Markets []repositories.Market
}

func (s *Service) GetAll(query repositories.MarketGetAll, accountID string) getAllResp {
	// Получение рынков
	markets := s.marRep.GetAll(query, accountID)
	return getAllResp{
		Err:     nil,
		Markets: markets,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Market, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.marRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Market, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.marRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
