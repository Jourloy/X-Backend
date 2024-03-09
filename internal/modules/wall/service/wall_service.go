package wall_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/wall_rep"
)

type Service struct {
	walRep repositories.IWallRepository
	secRep repositories.SectorRepository
	accRep repositories.AccountRepository
	cache  redis.Client
}

// Init создает сервис стены
func Init() *Service {

	walRep := wall_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	return &Service{
		walRep: walRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error `json:"error"`
}

// Create создает стену
func (s *Service) Create(body repositories.Wall, accountID string) createResp {
	// Проверка существования аккаунта
	account, err := s.accRep.GetOne(&repositories.AccountGet{ID: &accountID})
	if err != nil {
		return createResp{Err: err}
	}
	if account == nil {
		return createResp{Err: errors.New(`account not found`)}
	}

	// Проверка существования сектора
	sector, err := s.secRep.GetOne(&repositories.SectorGet{ID: &body.SectorID})
	if err != nil {
		return createResp{Err: err}
	}

	if sector == nil {
		return createResp{Err: errors.New(`sector not found`)}
	}

	s.walRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err  error
	Wall repositories.Wall
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	wall := repositories.Wall{ID: id, AccountID: accountID}
	s.walRep.GetOne(&wall)
	return getOneResp{
		Err:  nil,
		Wall: wall,
	}
}

type getAllResp struct {
	Err   error
	Walls []repositories.Wall
}

func (s *Service) GetAll(query repositories.WallGetAll, accountID string) getAllResp {
	// Получение стен
	walls := s.walRep.GetAll(query, accountID)
	return getAllResp{
		Err:   nil,
		Walls: walls,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Wall, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.walRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Wall, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.walRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
