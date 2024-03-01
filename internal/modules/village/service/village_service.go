package village_service

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[village-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	vRep  repositories.IVillageRepository
	cache redis.Client
}

// InitVillageService создает сервис поселений
func InitVillageService(vRep repositories.IVillageRepository, cache redis.Client) *Service {

	logger.Info(`Service initialized`)

	return &Service{
		vRep:  vRep,
		cache: cache,
	}
}

type createResp struct {
	Err error
}

// Create создает поселение
func (s *Service) Create(b repositories.Village, aID string, sID string) createResp {
	s.vRep.Create(&b, aID, sID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err     error
	Village repositories.Village
}

func (s *Service) GetOne(id string, aID string) getOneResp {
	village := s.vRep.GetOne(id, aID)
	return getOneResp{
		Err:     nil,
		Village: village,
	}
}

type getAllResp struct {
	Err      error
	Villages []repositories.Village
}

func (s *Service) GetAll(q repositories.VillageFindAll, aID string) getAllResp {
	// Получение поселений
	villages := s.vRep.GetAll(aID, q)
	return getAllResp{
		Err:      nil,
		Villages: villages,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(b repositories.Village, aID string) updateOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = aID

	s.vRep.UpdateOne(&b)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(b repositories.Village, aID string) deleteOneResp {
	// Перезапись accountID для безопасности
	b.AccountID = aID

	s.vRep.DeleteOne(&b)
	return deleteOneResp{
		Err: nil,
	}
}
