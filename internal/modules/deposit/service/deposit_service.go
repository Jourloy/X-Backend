package deposit_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/deposit_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
)

type Service struct {
	depRep repositories.IDepositRepository
	secRep repositories.ISectorRepository
	cache  redis.Client
}

// Init создает сервис залежи
func Init() *Service {

	depRep := deposit_rep.Repository
	secRep := sector_rep.Repository

	return &Service{
		depRep: depRep,
		secRep: secRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает рабочего
func (s *Service) Create(body repositories.Deposit) createResp {
	// Проверка существования сектора
	sector := repositories.Sector{ID: body.SectorID}
	s.secRep.GetOne(&sector)
	if sector.ID == `` {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.depRep.Create(&body)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err     error
	Deposit repositories.Deposit
}

func (s *Service) GetOne(id string, sectorID string) getOneResp {
	deposit := repositories.Deposit{ID: id, SectorID: sectorID}
	s.depRep.GetOne(&deposit)
	return getOneResp{
		Err:     nil,
		Deposit: deposit,
	}
}

type getAllResp struct {
	Err      error
	Deposits []repositories.Deposit
}

func (s *Service) GetAll(query repositories.DepositGetAll, accountID string) getAllResp {
	// Получение работников
	deposits := s.depRep.GetAll(query, accountID)
	return getAllResp{
		Err:      nil,
		Deposits: deposits,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Deposit) updateOneResp {
	s.depRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Deposit) deleteOneResp {
	s.depRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
