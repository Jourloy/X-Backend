package warrior_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/warrior_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[warrior-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	warRep repositories.IWarriorRepository
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис воина
func Init() *Service {

	warRep := warrior_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		warRep: warRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает воина
func (s *Service) Create(body repositories.Warrior, accountID string) createResp {
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

	s.warRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err     error
	Warrior repositories.Warrior
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	warrior := s.warRep.GetOne(id, accountID)
	return getOneResp{
		Err:     nil,
		Warrior: warrior,
	}
}

type getAllResp struct {
	Err      error
	Warriors []repositories.Warrior
}

func (s *Service) GetAll(query repositories.WarriorGetAll, accountID string) getAllResp {
	// Получение воинов
	warriors := s.warRep.GetAll(query, accountID)
	return getAllResp{
		Err:      nil,
		Warriors: warriors,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(b repositories.Warrior, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = accountID

	s.warRep.UpdateOne(&b)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(b repositories.Warrior, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = accountID

	s.warRep.DeleteOne(&b)
	return deleteOneResp{
		Err: nil,
	}
}
