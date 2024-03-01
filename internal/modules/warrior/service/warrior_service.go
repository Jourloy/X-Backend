package warrior_service

import (
	"errors"
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/village_rep"
	"github.com/jourloy/X-Backend/internal/repositories/warrior_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[warrior-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	warRep repositories.IWarriorRepository
	vilRep repositories.IVillageRepository
	cache  redis.Client
}

// Init создает сервис воина
func Init() *Service {

	warRep := warrior_rep.Repository
	vilRep := village_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		warRep: warRep,
		vilRep: vilRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает воина
func (s *Service) Create(b repositories.Warrior, vID string, accountID string) createResp {
	// Проверка существования воина
	village := s.vilRep.GetOne(vID, accountID)
	if village.ID == `` {
		return createResp{Err: errors.New(`village not found`)}
	}

	s.warRep.Create(&b, vID, accountID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err     error
	Warrior repositories.Warrior
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	warrior := s.warRep.GetOne(id, accountID)
	return getOneResp{
		Err:     nil,
		Warrior: warrior,
	}
}

type getAllResp struct {
	Err      error
	Warriors []repositories.Warrior
}

func (s *Service) GetAll(q repositories.WarriorFindAll, accountID string) getAllResp {
	// Получение воинов
	warriors := s.warRep.GetAll(accountID, q)
	return getAllResp{
		Err:      nil,
		Warriors: warriors,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(b repositories.Warrior, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = accountID

	s.warRep.UpdateOne(&b)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(b repositories.Warrior, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = accountID

	s.warRep.DeleteOne(&b)
	return deleteOneResp{
		Err: nil,
	}
}
