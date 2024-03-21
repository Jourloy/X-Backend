package creature_service

import (
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	creature_rep "github.com/jourloy/X-Backend/internal/modules/creature/repository"
	"github.com/jourloy/X-Backend/internal/repositories"
)

type Service struct {
	creRep repositories.CreatureRepository
	cache  redis.Client
}

// Init создает сервис существа
func Init() *Service {
	creature_rep.Init()

	creRep := creature_rep.Repository

	return &Service{
		creRep: creRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Creature *repositories.Creature
	Err      error
}

// Create создает существо
func (s *Service) Create(body repositories.CreatureCreate) createResp {
	account, err := s.creRep.Create(&body)
	return createResp{
		Err:      err,
		Creature: account,
	}
}

type getOneResp struct {
	Err      error
	Creature repositories.Creature
}

// GetOne получает существо по id
func (s *Service) GetOne(query *repositories.CreatureGet) getOneResp {
	creature, err := s.creRep.GetOne(query)
	return getOneResp{
		Err:      err,
		Creature: *creature,
	}
}

type getAllResp struct {
	Err       error
	Creatures []repositories.Creature
}

func (s *Service) GetAll(query *repositories.CreatureGet) getAllResp {
	// Получение существ
	creatures, err := s.creRep.GetAll(query)
	return getAllResp{
		Err:       err,
		Creatures: *creatures,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет существо
func (s *Service) UpdateOne(body repositories.Creature) updateOneResp {
	err := s.creRep.UpdateOne(&body)
	return updateOneResp{Err: err}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет существо
func (s *Service) DeleteOne(body repositories.Creature) deleteOneResp {
	err := s.creRep.DeleteOne(&body)
	return deleteOneResp{Err: err}
}
