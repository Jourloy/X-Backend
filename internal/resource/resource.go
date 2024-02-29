package resource

import (
	"encoding/json"
	"errors"
	"io"
	"os"

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

// InitResourceService создает сервис колоний
func InitResourceService(db repositories.IResourceRepository, cache redis.Client) *ResourceService {
	return &ResourceService{
		db:    db,
		cache: cache,
	}
}

func (s *ResourceService) Create(c *gin.Context, accountID string) {
	var body repositories.Resource
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	s.db.Create(&body, accountID)

	c.JSON(200, gin.H{`error`: ``})
}

func (s *ResourceService) GetOne(c *gin.Context, accountID string) {
	s.db.GetOne(c.Param(`id`), accountID)
}

func (s *ResourceService) UpdateOne(c *gin.Context, accountID string) {
	var body repositories.Resource
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.db.GetOne(body.ID, accountID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.db.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

func (s *ResourceService) UpdateDelete(c *gin.Context, accountID string) {
	var body repositories.Resource
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.db.GetOne(body.ID, accountID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
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
