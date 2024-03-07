package sector_service

import (
	"math/rand"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/repositories/deposit_rep"
	"github.com/jourloy/X-Backend/internal/repositories/node_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
)

type Service struct {
	secRep     repositories.ISectorRepository
	nodeRep    repositories.NodeRepository
	depositRep repositories.IDepositRepository
	cache      redis.Client
}

// Init создает сервис сектора
func Init() *Service {

	secRep := sector_rep.Repository
	nodeRep := node_rep.Repository
	depositRep := deposit_rep.Repository

	return &Service{
		secRep:     secRep,
		nodeRep:    nodeRep,
		depositRep: depositRep,
		cache:      *cache.Client,
	}
}

type CreateOptions struct {
	// Глобальные координаты
	X, Y int

	// Насколько сложная местность. Минимум 0, максимум 100
	Difficult int

	// Насколько непроходимая местность. Минимум 0, максимум 100
	Walkable int

	// Обилие ресурсов. Минимум 0, максимум 100
	Abundance int

	// Могут ли появится редкие ресурсы
	IsRare bool
}

type createResp struct {
	Err error
}

// Генерация сектора
func (s *Service) Create(body CreateOptions) createResp {
	sector := repositories.Sector{}
	s.secRep.Create(&sector)

	nodes := []repositories.Node{}

	// Создание узлов
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			node := repositories.Node{
				X:         x,
				Y:         y,
				Walkable:  true,
				Difficult: 0,
				SectorID:  sector.ID,
			}

			s.nodeRep.Create(&node)

			nodes = append(nodes, node)

			resourceCreateRand := rand.Intn(10)
			if resourceCreateRand > 5 {
				resourceTypeRand := rand.Intn(2)

				resourceType := `wood`
				if resourceTypeRand == 1 {
					resourceType = `stone`
				}

				deposit := repositories.Deposit{
					X:        x,
					Y:        y,
					Type:     resourceType,
					SectorID: sector.ID,
				}

				s.depositRep.Create(&deposit)
			}
		}
	}

	return createResp{
		Err: nil,
	}
}

type getOneResp struct {
	Err    error
	Sector repositories.Sector
}

// GetOne получает сектор по id
func (s *Service) GetOne(id string) getOneResp {
	sector := repositories.Sector{ID: id}
	s.secRep.GetOne(&sector)
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
