package building_service

import (
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/building_rep"
)

type Service struct {
	buiRep repositories.BuildingRepository
	cache  redis.Client
}

// Init создает сервис постройки
func Init() *Service {

	buiRep := building_rep.Repository

	return &Service{
		buiRep: buiRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Building *repositories.Building
	Err      error
}

// Create создает постройку
func (s *Service) Create(body repositories.BuildingCreate) createResp {
	account, err := s.buiRep.Create(&body)
	return createResp{
		Err:      err,
		Building: account,
	}
}

type getOneResp struct {
	Err      error
	Building repositories.Building
}

// GetOne получает постройку по id
func (s *Service) GetOne(query *repositories.BuildingGet) getOneResp {
	building, err := s.buiRep.GetOne(query)
	return getOneResp{
		Err:      err,
		Building: *building,
	}
}

type getAllResp struct {
	Err       error
	Buildings []repositories.Building
}

func (s *Service) GetAll(query *repositories.BuildingGet) getAllResp {
	// Получение построек
	buildings, err := s.buiRep.GetAll(query)
	return getAllResp{
		Err:       err,
		Buildings: *buildings,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет постройку
func (s *Service) UpdateOne(body repositories.Building) updateOneResp {
	err := s.buiRep.UpdateOne(&body)
	return updateOneResp{Err: err}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет постройку
func (s *Service) DeleteOne(body repositories.Building) deleteOneResp {
	err := s.buiRep.DeleteOne(&body)
	return deleteOneResp{Err: err}
}
