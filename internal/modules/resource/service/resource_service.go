package resource_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	resource_rep "github.com/jourloy/X-Backend/internal/repositories/resource"
	sector_rep "github.com/jourloy/X-Backend/internal/repositories/sector"
)

type ResourceService struct {
	resourceRep repositories.ResourceRepository
	sectorRep   repositories.SectorRepository
	accRep      repositories.AccountRepository
	cache       redis.Client
}

// Init создает сервис ресурса
func Init() *ResourceService {
	resourceRep := resource_rep.Repository
	sectorRep := sector_rep.Repository

	return &ResourceService{
		resourceRep: resourceRep,
		sectorRep:   sectorRep,
		cache:       *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает ресурс
func (s *ResourceService) Create(body repositories.ResourceCreate) createResp {
	// Проверка существования аккаунта
	account, err := s.accRep.GetOne(&repositories.AccountGet{ID: &body.CreatorID})
	if err != nil {
		return createResp{Err: err}
	}
	if account == nil {
		return createResp{Err: errors.New(`account not found`)}
	}

	// Проверка существования сектора
	sector, err := s.sectorRep.GetOne(&repositories.SectorGet{ID: &body.SectorID})
	if err != nil {
		return createResp{Err: err}
	}

	if sector == nil {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.resourceRep.Create(body)
	return createResp{Err: nil}
}
