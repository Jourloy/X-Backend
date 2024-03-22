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
	resourceRep repositories.IResourceRepository
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
func (s *ResourceService) Create(body repositories.Resource) createResp {
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

	s.resourceRep.Create(&body)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err      error
	Resource repositories.Resource
}

// GetOne получает ресурс по id
func (s *ResourceService) GetOne(resourceID string) getOneResp {
	resource := repositories.Resource{ID: resourceID}
	s.resourceRep.GetOne(resource)
	return getOneResp{
		Err:      nil,
		Resource: resource,
	}
}

type getAllResp struct {
	Err       error
	Resources []repositories.Resource
}

// GetAll возвращает все ресурсы
func (s *ResourceService) GetAll(query repositories.ResourceGetAll) getAllResp {
	resources := s.resourceRep.GetAll(query)
	return getAllResp{
		Err:       nil,
		Resources: resources,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет ресурс
func (s *ResourceService) UpdateOne(body repositories.Resource) updateOneResp {
	s.resourceRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет ресурс
func (s *ResourceService) DeleteOne(body repositories.Resource) deleteOneResp {
	s.resourceRep.UpdateOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
