package colony

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
		Prefix: `[colony]`,
		Level:  log.DebugLevel,
	})
)

type ColonyService struct {
	cRep  repositories.IColonyRepository
	cache redis.Client
}

// InitColonyService создает сервис колонии
func InitColonyService(cRep repositories.IColonyRepository, cache redis.Client) *ColonyService {

	logger.Info(`ColonyService initialized`)

	return &ColonyService{
		cRep:  cRep,
		cache: cache,
	}
}

// Create создает колонию
func (s *ColonyService) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)
	var body repositories.Colony
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	s.cRep.Create(&body, accountID, c.Query(`placeID`))

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает колонию по id
func (s *ColonyService) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)
	s.cRep.GetOne(c.Param(`id`), accountID)
}

// GetAll возвращает все колонии
func (s *ColonyService) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	query := repositories.ColonyFindAll{}

	if q := c.Query(`limit`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Limit = &n
	}
	if q := c.Query(`balance`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Balance = &n
	}
	if q := c.Query(`maxStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.MaxStorage = &n
	}
	if q := c.Query(`usedStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.UsedStorage = &n
	}

	colonies := s.cRep.GetAll(accountID, query)

	c.JSON(200, gin.H{
		`error`:    ``,
		`colonies`: colonies,
		`count`:    len(colonies),
	})
}

// UpdateOne обновляет колонию
func (s *ColonyService) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)
	var body repositories.Colony
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.cRep.GetOne(body.ID, accountID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.cRep.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет колонию
func (s *ColonyService) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)
	var body repositories.Colony
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.cRep.GetOne(body.ID, accountID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.cRep.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

func (s *ColonyService) parseBody(c *gin.Context, body interface{}) error {
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
