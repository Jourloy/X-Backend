package trader_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/trader_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[trader]`,
		Level:  log.DebugLevel,
	})
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

	logger.Info(`Service initialized`)

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
func (s *Service) Create(body repositories.Trader, accountID string) createResp {
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

	s.traRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err    error
	Trader repositories.Trader
}

// GetOne получает торговца по его ID
func (s *Service) GetOne(id string, accountID string) getOneResp {
	trader := s.traRep.GetOne(id, accountID)
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
