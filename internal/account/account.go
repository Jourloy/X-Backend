package account

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
		Prefix: `[account]`,
		Level:  log.DebugLevel,
	})
)

type AccountService struct {
	db    repositories.IAccountRepository
	cache redis.Client
}

// InitAccountService создает сервис аккаунта
func InitAccountService(db repositories.IAccountRepository, cache redis.Client) *AccountService {
	return &AccountService{
		db:    db,
		cache: cache,
	}
}

// Create создает аккаунт
func (s *AccountService) Create(c *gin.Context, accountID string) {
	var body repositories.Account
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	s.db.Create(&body, accountID)

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает аккаунт по id
func (s *AccountService) GetOne(c *gin.Context, accountID string) {
	s.db.GetOne(c.Param(`id`), accountID)
}

// UpdateOne обновляет аккаунт
func (s *AccountService) UpdateOne(c *gin.Context, accountID string) {
	var body repositories.Account
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

// DeleteOne удаляет аккаунт
func (s *AccountService) UpdateDelete(c *gin.Context, accountID string) {
	var body repositories.Account
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

func (s *AccountService) parseBody(c *gin.Context, body interface{}) error {
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
