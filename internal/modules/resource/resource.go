package resource

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[resource]`,
		Level:  log.DebugLevel,
	})
)

type ResourceService struct {
	db    repositories.IResourceRepository
	cache redis.Client
}

// InitResourceService создает сервис ресурса
func InitResourceService(db repositories.IResourceRepository, cache redis.Client) *ResourceService {

	logger.Info(`ResourceService initialized`)

	return &ResourceService{
		db:    db,
		cache: cache,
	}
}

// Create создает ресурс
func (s *ResourceService) Create(c *gin.Context) {
	// Парсинг body
	var body repositories.Resource
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Получение sectorID
	q := c.Query(`sectorID`)
	if q == `` {
		logger.Error(`sectorID is required`)
		c.JSON(400, gin.H{`error`: `sectorID is required`})
	}

	// Создание
	s.db.Create(&body, q)

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает ресурс по id
func (s *ResourceService) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	s.db.GetOne(c.Param(`id`), accountID)
}

// GetAll возвращает все ресурсы
func (s *ResourceService) GetAll(c *gin.Context) {

	// Создание фильтров
	query := repositories.ResourceFindAll{}
	if q := c.Query(`limit`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Limit = &n
	}
	if q := c.Query(`type`); q != `` {
		query.Type = &q
	}
	if q := c.Query(`amount`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Amount = &n
	}
	if q := c.Query(`weight`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Weight = &n
	}

	sectorID := c.Query(`sectorID`)
	if sectorID == `` {
		logger.Error(`sectorID is required`)
		c.JSON(400, gin.H{`error`: `sectorID is required`})
	}

	// Получение ресурсов
	resources := s.db.GetAll(sectorID, query)
	c.JSON(200, gin.H{
		`error`:     ``,
		`resources`: resources,
		`count`:     len(resources),
	})
}

// UpdateOne обновляет ресурс
func (s *ResourceService) UpdateOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Resource
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	s.db.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет ресурс
func (s *ResourceService) DeleteOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Resource
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	s.db.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}
