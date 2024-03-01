package warrior_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[warrior-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	wRep  repositories.IWarriorRepository
	cRep  repositories.IVillageRepository
	cache redis.Client
}

// InitWarriorService создает сервис воина
func InitWarriorService(wRep repositories.IWarriorRepository, cRep repositories.IVillageRepository, cache redis.Client) *Service {

	logger.Info(`WarriorService initialized`)

	return &Service{
		wRep:  wRep,
		cRep:  cRep,
		cache: cache,
	}
}

type createResp struct {
	Err error
}

// Create создает воина
func (s *Service) Create(b repositories.Warrior, vID string, aID string) createResp {
	// Проверка существования воина
	village := s.cRep.GetOne(vID, aID)
	if village.ID == `` {
		return createResp{Err: errors.New(`village not found`)}
	}

	s.wRep.Create(&b, vID, aID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err     error
	Warrior repositories.Warrior
}

func (s *Service) GetOne(id string, aID string) getOneResp {
	warrior := s.wRep.GetOne(id, aID)
	return getOneResp{
		Err:     nil,
		Warrior: warrior,
	}
}

type getAllResp struct {
	Err      error
	Warriors []repositories.Warrior
}

func (s *Service) GetAll(q repositories.WarriorFindAll, aID string) getAllResp {
	// Получение воинов
	warriors := s.wRep.GetAll(aID, q)
	return getAllResp{
		Err:      nil,
		Warriors: warriors,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(b repositories.Warrior, aID string) updateOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = aID

	s.wRep.UpdateOne(&b)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(b repositories.Warrior, aID string) deleteOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = aID

	s.wRep.DeleteOne(&b)
	return deleteOneResp{
		Err: nil,
	}
}
