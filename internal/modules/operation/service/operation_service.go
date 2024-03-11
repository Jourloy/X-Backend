package operation_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/building_rep"
	"github.com/jourloy/X-Backend/internal/repositories/operation_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
)

type Service struct {
	accountRep   repositories.AccountRepository
	buildingRep  repositories.BuildingRepository
	sectorRep    repositories.SectorRepository
	operationRep repositories.OperationRepository
	cache        redis.Client
}

// Init создает сервис операций
func Init() *Service {

	accountRep := account_rep.Repository
	buildingRep := building_rep.Repository
	sectorRep := sector_rep.Repository
	operationRep := operation_rep.Repository

	return &Service{
		accountRep:   accountRep,
		buildingRep:  buildingRep,
		sectorRep:    sectorRep,
		operationRep: operationRep,
		cache:        *cache.Client,
	}
}

type createResp struct {
	Operation *repositories.Operation
	Err       error
	Code      int
}

// Create создает операцию
func (s *Service) Create(body repositories.OperationCreate) createResp {
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
	if sector, err := s.sectorRep.GetOne(&repositories.SectorGet{ID: &body.SectorID}); err != nil {
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

	// Проверка постройки
	if building, err := s.buildingRep.GetOne(&repositories.BuildingGet{ID: &body.BuildingID}); err != nil {
		return createResp{
			Err:  err,
			Code: 400,
		}
	} else if building == nil {
		return createResp{
			Err:  errors.New(`building not found`),
			Code: 400,
		}
	}

	// Создание операции
	operation, err := s.operationRep.Create(&body)

	return createResp{
		Code:      200,
		Err:       err,
		Operation: operation,
	}
}

type getOneResp struct {
	Err       error
	Operation repositories.Operation
}

// GetOne получает операцию, подходящую под условие
func (s *Service) GetOne(query *repositories.OperationGet) getOneResp {
	operation, err := s.operationRep.GetOne(query)
	return getOneResp{
		Err:       err,
		Operation: *operation,
	}
}

type getAllResp struct {
	Err        error
	Operations []repositories.Operation
}

// GetAll получает все операции, подходящие под условие
func (s *Service) GetAll(query *repositories.OperationGet) getAllResp {
	// Получение операций
	operations, err := s.operationRep.GetAll(query)
	return getAllResp{
		Err:        err,
		Operations: *operations,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет операцию
func (s *Service) UpdateOne(body repositories.Operation) updateOneResp {
	err := s.operationRep.UpdateOne(&body)
	return updateOneResp{Err: err}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет операцию
func (s *Service) DeleteOne(body repositories.Operation) deleteOneResp {
	err := s.operationRep.DeleteOne(&body)
	return deleteOneResp{Err: err}
}
