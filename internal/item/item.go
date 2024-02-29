package item

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
		Prefix: `[item]`,
		Level:  log.DebugLevel,
	})
)

type ItemService struct {
	iRep  repositories.IItemRepository
	wRep  repositories.IWorkerRepository
	cache redis.Client
}

// InitItemService создает сервис колоний
func InitItemService(iRep repositories.IItemRepository, wRep repositories.IWorkerRepository, cache redis.Client) *ItemService {
	return &ItemService{
		iRep:  iRep,
		wRep:  wRep,
		cache: cache,
	}
}

func (s *ItemService) Create(c *gin.Context) {
	var body repositories.Item
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	c.JSON(200, gin.H{`error`: ``})
}

func (s *ItemService) GetOne(c *gin.Context) {
	s.iRep.GetOne(c.Param(`id`), c.Param(`parentId`))
}

func (s *ItemService) UpdateOne(c *gin.Context, accountID string) {
	var body repositories.Item
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.iRep.GetOne(body.ID, accountID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.iRep.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

func (s *ItemService) DeleteOne(c *gin.Context, accountID string) {
	var body repositories.Item
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.iRep.GetOne(body.ID, accountID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.iRep.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

func (s *ItemService) parseBody(c *gin.Context, body interface{}) error {
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
