package sector_service

import (
	"math/rand"

	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/repositories"
	deposit_rep "github.com/jourloy/X-Backend/internal/repositories/deposit"
	node_rep "github.com/jourloy/X-Backend/internal/repositories/node"
	sector_rep "github.com/jourloy/X-Backend/internal/repositories/sector"
)

type Service struct {
	sectorRep  repositories.SectorRepository
	nodeRep    repositories.NodeRepository
	depositRep repositories.IDepositRepository
	cache      redis.Client
}

// Init создает сервис сектора
func Init() *Service {
	sectorRep := sector_rep.Repository
	nodeRep := node_rep.Repository
	depositRep := deposit_rep.Repository

	return &Service{
		sectorRep:  sectorRep,
		nodeRep:    nodeRep,
		depositRep: depositRep,
		cache:      *cache.Client,
	}
}

type CreateOptions struct {
	// Глобальные координаты
	X int `json:"x"`
	Y int `json:"y"`

	// Насколько сложная местность. Минимум 0, максимум 100
	Difficult int `json:"difficult"`

	// Насколько непроходимая местность. Минимум 0, максимум 100
	Walkable int `json:"walkable"`

	// Обилие ресурсов. Минимум 0, максимум 100
	Abundance int `json:"abundance"`

	// Могут ли появится редкие ресурсы
	IsRare bool `json:"isRare"`
}

type createResp struct {
	Err error
}

// Генерация сектора
func (s *Service) Create(body CreateOptions) createResp {
	sector, err := s.sectorRep.Create(&repositories.SectorCreate{X: body.X, Y: body.Y})
	if err != nil {
		return createResp{Err: err}
	}

	go s.generateNodes(sector.ID)
	go s.generateDeposits(sector.ID)

	return createResp{
		Err: nil,
	}
}

func (s *Service) generateNodes(sectorID string) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			node := repositories.Node{
				X:         x,
				Y:         y,
				Walkable:  true,
				Difficult: 0,
				SectorID:  sectorID,
			}

			s.nodeRep.Create(&node)
		}
	}
}

func (s *Service) generateDeposits(sectorID string) {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
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
					SectorID: sectorID,
				}

				s.depositRep.Create(&deposit)
			}
		}
	}
}

type getOneResp struct {
	Err    error
	Sector repositories.Sector
}

// GetOne получает сектор по id
func (s *Service) GetOne(id string) getOneResp {
	sector, err := s.sectorRep.GetOne(&repositories.SectorGet{ID: &id})
	if err != nil {
		return getOneResp{Err: err}
	}

	return getOneResp{Sector: *sector}
}

type getAllResp struct {
	Err     error
	Sectors []repositories.Sector
}

// GetAll возвращает все сектора
func (s *Service) GetAll(query repositories.SectorGet) getAllResp {
	// Получение секторов
	sectors, err := s.sectorRep.GetAll(&query)
	if err != nil {
		return getAllResp{Err: err}
	}

	return getAllResp{
		Err:     nil,
		Sectors: *sectors,
	}
}

type updateOneResp struct {
	Err error
}

// UpdateOne обновляет сектор
func (s *Service) UpdateOne(body repositories.Sector) updateOneResp {
	err := s.sectorRep.UpdateOne(&body)
	return updateOneResp{Err: err}
}

type deleteOneResp struct {
	Err error
}

// DeleteOne удаляет сектор
func (s *Service) DeleteOne(body repositories.Sector) deleteOneResp {
	err := s.sectorRep.UpdateOne(&body)
	return deleteOneResp{Err: err}
}
