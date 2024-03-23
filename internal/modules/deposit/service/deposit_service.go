package deposit_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	deposit_rep "github.com/jourloy/X-Backend/internal/repositories/deposit"
	sector_rep "github.com/jourloy/X-Backend/internal/repositories/sector"
)

type Service struct {
	depRep repositories.DepositRepository
	secRep repositories.SectorRepository
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
func (s *Service) Create(body repositories.DepositCreate) createResp {
	// Проверка существования сектора
	sector, err := s.secRep.GetOne(&repositories.SectorGet{ID: &body.SectorID})
	if err != nil {
		return createResp{Err: err}
	}

	if sector == nil {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.depRep.Create(body)
	return createResp{Err: nil}
}
