package plan_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	account_rep "github.com/jourloy/X-Backend/internal/modules/account/repository"
	plan_rep "github.com/jourloy/X-Backend/internal/modules/plan/repository"
	sector_rep "github.com/jourloy/X-Backend/internal/modules/sector/repository"
	"github.com/jourloy/X-Backend/internal/repositories"
)

type Service struct {
	plaRep repositories.IPlanRepository
	secRep repositories.SectorRepository
	accRep repositories.AccountRepository
	cache  redis.Client
}

// Init создает сервис планируемой постройки
func Init() *Service {

	plaRep := plan_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

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
func (s *Service) Create(body repositories.PlanCreate) createResp {
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

	s.plaRep.Create(&body)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err  error
	Plan repositories.Plan
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	plan := repositories.Plan{ID: id, AccountID: accountID}
	s.plaRep.GetOne(&plan)
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
