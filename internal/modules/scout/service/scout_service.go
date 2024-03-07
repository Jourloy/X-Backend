package scout_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/scout_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
)

type Service struct {
	scoRep repositories.IScoutRepository
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис разведчика
func Init() *Service {

	scoRep := scout_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	return &Service{
		scoRep: scoRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает разведчика
func (s *Service) Create(body repositories.Scout, accountID string) createResp {
	// Проверка существования аккаунта
	account := repositories.Account{ID: accountID}
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

	s.scoRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err   error
	Scout repositories.Scout
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	scout := s.scoRep.GetOne(id, accountID)
	return getOneResp{
		Err:   nil,
		Scout: scout,
	}
}

type getAllResp struct {
	Err    error
	Scouts []repositories.Scout
}

func (s *Service) GetAll(query repositories.ScoutGetAll, accountID string) getAllResp {
	// Получение разведчиков
	scouts := s.scoRep.GetAll(query, accountID)
	return getAllResp{
		Err:    nil,
		Scouts: scouts,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Scout, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.scoRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Scout, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.scoRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
