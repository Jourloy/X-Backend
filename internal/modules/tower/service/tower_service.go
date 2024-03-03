package tower_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/tower_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[tower-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	towRep repositories.ITowerRepository
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис башни
func Init() *Service {

	towRep := tower_rep.Repository
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

// Create создает башню
func (s *Service) Create(body repositories.Tower, accountID string) createResp {
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

	s.towRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err   error
	Tower repositories.Tower
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	tower := s.towRep.GetOne(id, accountID)
	return getOneResp{
		Err:   nil,
		Tower: tower,
	}
}

type getAllResp struct {
	Err    error
	Towers []repositories.Tower
}

func (s *Service) GetAll(query repositories.TowerGetAll, accountID string) getAllResp {
	// Получение башен
	towers := s.towRep.GetAll(query, accountID)
	return getAllResp{
		Err:    nil,
		Towers: towers,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Tower, accountID string) updateOneResp {
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

func (s *Service) DeleteOne(body repositories.Tower, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.towRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
