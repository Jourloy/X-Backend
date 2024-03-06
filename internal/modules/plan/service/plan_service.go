package plan_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/plan_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[plan-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	plaRep repositories.IPlanRepository
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис планируемой постройки
func Init() *Service {

	plaRep := plan_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		plaRep: plaRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает планируемую постройку
func (s *Service) Create(body repositories.Plan, accountID string) createResp {
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

	s.plaRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err  error
	Plan repositories.Plan
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	plan := s.plaRep.GetOne(id, accountID)
	return getOneResp{
		Err:  nil,
		Plan: plan,
	}
}

type getAllResp struct {
	Err   error
	Plans []repositories.Plan
}

func (s *Service) GetAll(query repositories.PlanGetAll, accountID string) getAllResp {
	// Получение планируемых построек
	plans := s.plaRep.GetAll(query, accountID)
	return getAllResp{
		Err:   nil,
		Plans: plans,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Plan, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.plaRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Plan, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.plaRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
