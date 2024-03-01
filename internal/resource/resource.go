package resource

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/repositories"
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
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Получение placeID
	q := c.Query(`placeID`)
	if q == `` {
		logger.Error(`placeID is required`)
		c.JSON(400, gin.H{`error`: `placeID is required`})
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

	placeID := c.Query(`weight`)
	if placeID == `` {
		logger.Error(`placeID is required`)
		c.JSON(400, gin.H{`error`: `placeID is required`})
	}

	// Получение ресурсов
	resources := s.db.GetAll(placeID, query)
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
	if err := s.parseBody(c, &body); err != nil {
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
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	s.db.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

func (s *ResourceService) parseBody(c *gin.Context, body interface{}) error {
	// Проверка body
	if c.Request.Body == nil {
		return errors.New(`body not found`)
	}

	// Чтение body
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	// Парсинг
	if err := json.Unmarshal(b, &body); err != nil {
		return err
	}

	return nil
}
