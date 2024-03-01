package village_service

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/village_rep"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[village-service]`,
		Level:  log.DebugLevel,
	})
)

type Service struct {
	vilRep repositories.IVillageRepository
	cache  redis.Client
}

// Init создает сервис поселений
func Init() *Service {

	vilRep := village_rep.Repository

	logger.Info(`Service initialized`)

	return &Service{
		vilRep: vilRep,
		cache:  *cache.Client,
	}
}

type createResp struct {
	Err error
}

// Create создает поселение
func (s *Service) Create(body repositories.Village, accountID string, sID string) createResp {
	s.vilRep.Create(&body, accountID, sID)
	return createResp{Err: nil}
}

type getOneResp struct {
	Err     error
	Village repositories.Village
}

func (s *Service) GetOne(id string, accountID string) getOneResp {
	village := s.vilRep.GetOne(id, accountID)
	return getOneResp{
		Err:     nil,
		Village: village,
	}
}

type getAllResp struct {
	Err      error
	Villages []repositories.Village
}

func (s *Service) GetAll(query repositories.VillageFindAll, accountID string) getAllResp {
	// Получение поселений
	villages := s.vilRep.GetAll(accountID, query)
	return getAllResp{
		Err:      nil,
		Villages: villages,
	}
}

type updateOneResp struct {
	Err error
}

func (s *Service) UpdateOne(body repositories.Village, accountID string) updateOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.vilRep.UpdateOne(&body)
	return updateOneResp{
		Err: nil,
	}
}

type deleteOneResp struct {
	Err error
}

func (s *Service) DeleteOne(body repositories.Village, accountID string) deleteOneResp {
	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.vilRep.DeleteOne(&body)
	return deleteOneResp{
		Err: nil,
	}
}
