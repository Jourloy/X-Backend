package tower_service

import (
	"errors"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/tower_rep"
)

type Service struct {
	towRep repositories.ITowerRepository
	secRep repositories.SectorRepository
	accRep repositories.AccountRepository
	cache  redis.Client
}

// Init создает сервис башни
func Init() *Service {

	towRep := tower_rep.Repository
	secRep := sector_rep.Repository
	accRep := account_rep.Repository

	return &Service{
		towRep: towRep,
		secRep: secRep,
		accRep: accRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает башню
func (s *Service) Create(body repositories.Tower, accountID string) createResp {
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

	s.towRep.Create(&body, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err   error
	Tower repositories.Tower
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	tower := repositories.Tower{ID: id, AccountID: accountID}
	s.towRep.GetOne(&tower)
	return getOneResp{
		Err:   nil,
		Tower: tower,
	}
}

type getAllResp struct {
	Err    error
	Towers []repositories.Tower
}

func (s *Service) GetAll(query repositories.TowerGetAll, accountID string) getAllResp {
	// Получение башен
	towers := s.towRep.GetAll(query, accountID)
	return getAllResp{
		Err:    nil,
		Towers: towers,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Tower, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.towRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Tower, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.towRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
