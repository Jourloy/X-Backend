package place

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
		Prefix: `[place]`,
		Level:  log.DebugLevel,
	})
)

type PlaceService struct {
	db    repositories.IPlaceRepository
	cache redis.Client
}

// InitPlaceService создает сервис колоний
func InitPlaceService(db repositories.IPlaceRepository, cache redis.Client) *PlaceService {
	return &PlaceService{
		db:    db,
		cache: cache,
	}
}

func (s *PlaceService) Create(c *gin.Context) {
	var body repositories.Place
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	s.db.Create(&body)

	c.JSON(200, gin.H{`error`: ``})
}

func (s *PlaceService) GetOne(c *gin.Context) {
	s.db.GetOne(c.Param(`id`))
}

func (s *PlaceService) UpdateOne(c *gin.Context) {
	var body repositories.Place
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.db.GetOne(body.ID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.db.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

func (s *PlaceService) DeleteOne(c *gin.Context) {
	var body repositories.Place
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.db.GetOne(body.ID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.db.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

func (s *PlaceService) parseBody(c *gin.Context, body interface{}) error {
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
