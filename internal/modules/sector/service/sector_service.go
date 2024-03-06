package sector_service

import (
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
)

type Service struct {
	secRep repositories.ISectorRepository
	cache  redis.Client
}

// Init создает сервис сектора
func Init() *Service {

	secRep := sector_rep.Repository

	return &Service{
		secRep: secRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает сектор
func (s *Service) Create(body repositories.Sector) createResp {
	s.secRep.Create(&body)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err    error
	Sector repositories.Sector
}

// GetOne получает сектор по id
func (s *Service) GetOne(id string) getOneResp {
	sector := s.secRep.GetOne(id)
	return getOneResp{
		Err:    nil,
		Sector: sector,
	}
}

type getAllResp struct {
	Err     error
	Sectors []repositories.Sector
}

// GetAll возвращает все сектора
func (s *Service) GetAll(query repositories.SectorGetAll) getAllResp {
	// Получение секторов
	sectors := s.secRep.GetAll(query)
	return getAllResp{
		Err:     nil,
		Sectors: sectors,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет сектор
func (s *Service) UpdateOne(body repositories.Sector) updateOneResp {
	s.secRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет сектор
func (s *Service) DeleteOne(body repositories.Sector) deleteOneResp {
	s.secRep.UpdateOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
