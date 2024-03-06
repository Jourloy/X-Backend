package resource_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/resource_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[resource-service]`,
		Level:  log.DebugLevel,
	})
)

type ResourceService struct {
	resRep repositories.IResourceRepository
	secRep repositories.ISectorRepository
	accRep repositories.IAccountRepository
	cache  redis.Client
}

// Init создает сервис ресурса
func Init() *ResourceService {

	resRep := resource_rep.Repository
	secRep := sector_rep.Repository

	logger.Info(`ResourceService initialized`)

	return &ResourceService{
		resRep: resRep,
		secRep: secRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает ресурс
func (s *ResourceService) Create(body repositories.Resource) createResp {
	// Проверка существования аккаунта
	account := s.accRep.GetOne(repositories.Account{ID: body.CreatorID})
	if account.ID == `` {
		return createResp{Err: errors.New(`account not found`)}
	}

	// Проверка существования сектора
	sector := s.secRep.GetOne(body.SectorID)
	if sector.ID == `` {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.resRep.Create(&body)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err      error
	Resource repositories.Resource
}

// GetOne получает ресурс по id
func (s *ResourceService) GetOne(resourceID string) getOneResp {
	resource := s.resRep.GetOne(resourceID)
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
	resources := s.resRep.GetAll(query)
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
	s.resRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет ресурс
func (s *ResourceService) DeleteOne(body repositories.Resource) deleteOneResp {
	s.resRep.UpdateOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
