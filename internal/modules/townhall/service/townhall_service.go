package townhall_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/townhall_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[townhall-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	towRep repositories.ITownhallRepository
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис главного здания
func Init() *Service {

	towRep := townhall_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		towRep: towRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает главное здание
func (s *Service) Create(body repositories.Townhall, accountID string) createResp {
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

	s.towRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err      error
	Townhall repositories.Townhall
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	townhall := s.towRep.GetOne(id, accountID)
	return getOneResp{
		Err:      nil,
		Townhall: townhall,
	}
}

type getAllResp struct {
	Err       error
	Townhalls []repositories.Townhall
}

func (s *Service) GetAll(query repositories.TownhallGetAll, accountID string) getAllResp {
	// Получение главных зданий
	townhalls := s.towRep.GetAll(query, accountID)
	return getAllResp{
		Err:       nil,
		Townhalls: townhalls,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Townhall, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.towRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Townhall, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.towRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
