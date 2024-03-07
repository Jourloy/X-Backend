package trader_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/trader_rep"
)

type Service struct {
	traRep repositories.ITraderRepository
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис торговца
func Init() *Service {

	traRep := trader_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	return &Service{
		traRep: traRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает торговца
func (s *Service) Create(body repositories.TraderCreate) createResp {
	// Проверка существования аккаунта
	account := repositories.Account{ID: body.AccountID}
	s.accRep.GetOne(&account)
	if account.Username == `` {
		return createResp{Err: errors.New(`account not found`)}
	}

	// Проверка существования сектора
	sector := repositories.Sector{ID: body.SectorID}
	s.secRep.GetOne(&sector)
	if sector.ID == `` {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.traRep.Create(&body)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err    error
	Trader repositories.Trader
}

// GetOne получает торговца по его ID
func (s *Service) GetOne(id string, accountID string) getOneResp {
	trader := repositories.Trader{ID: id, AccountID: accountID}
	s.traRep.GetOne(&trader)
	return getOneResp{
		Err:    nil,
		Trader: trader,
	}
}

type getAllResp struct {
	Err     error
	Traders []repositories.Trader
}

// GetAll возвращает всех торговцев
func (s *Service) GetAll(query repositories.TraderGetAll, accountID string) getAllResp {
	// Получение работников
	traders := s.traRep.GetAll(query, accountID)
	return getAllResp{
		Err:     nil,
		Traders: traders,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет торговца
func (s *Service) UpdateOne(body repositories.Trader, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.traRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет торговца
func (s *Service) DeleteOne(body repositories.Trader, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.traRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
