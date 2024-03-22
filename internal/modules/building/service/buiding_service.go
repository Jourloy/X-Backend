package building_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	account_rep "github.com/jourloy/X-Backend/internal/repositories/account"
	building_rep "github.com/jourloy/X-Backend/internal/repositories/building"
	node_rep "github.com/jourloy/X-Backend/internal/repositories/node"
	sector_rep "github.com/jourloy/X-Backend/internal/repositories/sector"
)

type Service struct {
	accountRep  repositories.AccountRepository
	sectorRep   repositories.SectorRepository
	buildingRep repositories.BuildingRepository
	nodeRep     repositories.NodeRepository
	cache       redis.Client
}

// Init создает сервис постройки
func Init() *Service {
	accountRep := account_rep.Repository
	sectorRep := sector_rep.Repository
	buildingRep := building_rep.Repository
	nodeRep := node_rep.Repository

	return &Service{
		accountRep:  accountRep,
		sectorRep:   sectorRep,
		buildingRep: buildingRep,
		nodeRep:     nodeRep,
		cache:       *cache.Client,
	}
}

type createResp struct {
	Building *repositories.Building
	Err      error
	Code     int
}

// Create создает постройку
func (s *Service) Create(body repositories.BuildingCreate) createResp {
	// Проверка аккаунта
	if account, err := s.accountRep.GetOne(&repositories.AccountGet{ID: &body.AccountID}); err != nil {
		return createResp{
			Err:  err,
			Code: 400,
		}
	} else if account == nil {
		return createResp{
			Err:  errors.New(`account not found`),
			Code: 404,
		}
	}

	// Проверка сектора
	if sector, err := s.sectorRep.GetOne(&repositories.SectorGet{ID: &body.AccountID}); err != nil {
		return createResp{
			Err:  err,
			Code: 400,
		}
	} else if sector == nil {
		return createResp{
			Err:  errors.New(`sector not found`),
			Code: 404,
		}
	}

	// Создание постройки
	building, err := s.buildingRep.Create(&body)

	return createResp{
		Code:     200,
		Err:      err,
		Building: building,
	}
}

type getOneResp struct {
	Err      error
	Building repositories.Building
}

// GetOne получает постройку, подходящую под условие
func (s *Service) GetOne(query *repositories.BuildingGet) getOneResp {
	building, err := s.buildingRep.GetOne(query)
	return getOneResp{
		Err:      err,
		Building: *building,
	}
}

type getAllResp struct {
	Err       error
	Buildings []repositories.Building
}

// GetAll получает все постройки, подходящие под условие
func (s *Service) GetAll(query *repositories.BuildingGet) getAllResp {
	// Получение построек
	buildings, err := s.buildingRep.GetAll(query)
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
	err := s.buildingRep.UpdateOne(&body)
	return updateOneResp{Err: err}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет постройку
func (s *Service) DeleteOne(body repositories.Building) deleteOneResp {
	err := s.buildingRep.DeleteOne(&body)
	return deleteOneResp{Err: err}
}

type placeTownHallResponse struct {
	Err      error
	Code     int
	Building *repositories.Building
}

func (s *Service) PlaceTownHall(body repositories.BuildingCreate) placeTownHallResponse {
	// Проверка аккаунта
	if account, err := s.accountRep.GetOne(&repositories.AccountGet{ID: &body.AccountID}); err != nil {
		return placeTownHallResponse{
			Err:  err,
			Code: 400,
		}
	} else if account == nil {
		return placeTownHallResponse{
			Err:  errors.New(`account not found`),
			Code: 404,
		}
	}

	// Проверка сектора
	if sector, err := s.sectorRep.GetOne(&repositories.SectorGet{ID: &body.SectorID}); err != nil {
		return placeTownHallResponse{
			Err:  err,
			Code: 400,
		}
	} else if sector == nil {
		return placeTownHallResponse{
			Err:  errors.New(`sector not found`),
			Code: 404,
		}
	}

	// Проверка узла
	if node, err := s.nodeRep.GetOne(&repositories.NodeGet{SectorID: &body.SectorID, X: &body.X, Y: &body.Y}); err != nil {
		return placeTownHallResponse{
			Err:  err,
			Code: 400,
		}
	} else if node == nil {
		return placeTownHallResponse{
			Err:  errors.New(`node not found`),
			Code: 404,
		}
	} else if !node.Walkable {
		return placeTownHallResponse{
			Err:  errors.New(`cannot place here`),
			Code: 400,
		}
	}

	// Проверка наличия townhall
	t := `townhall`
	if townhall, err := s.buildingRep.GetOne(&repositories.BuildingGet{Type: &t}); err != nil {
		return placeTownHallResponse{
			Err:  err,
			Code: 400,
		}
	} else if townhall != nil {
		return placeTownHallResponse{
			Err:  errors.New(`townhall already exist in this sector`),
			Code: 400,
		}
	}

	building, err := s.buildingRep.Create(&body)

	return placeTownHallResponse{
		Code:     200,
		Err:      err,
		Building: building,
	}
}
