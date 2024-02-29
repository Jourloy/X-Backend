package worker

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
		Prefix: `[worker]`,
		Level:  log.DebugLevel,
	})
)

type WorkerService struct {
	wRep  repositories.IWorkerRepository
	cRep  repositories.IColonyRepository
	cache redis.Client
}

// InitWorkerService создает сервис рабочего
func InitWorkerService(wRep repositories.IWorkerRepository, cRep repositories.IColonyRepository, cache redis.Client) *WorkerService {

	logger.Info(`WorkerService initialized`)

	return &WorkerService{
		wRep:  wRep,
		cRep:  cRep,
		cache: cache,
	}
}

// Create создает рабочего
func (s *WorkerService) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Worker
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Проверка существования колонии
	colonyID := c.Query(`colonyID`)
	colony := s.cRep.GetOne(colonyID, accountID)
	if colony.ID == `` {
		logger.Error(`Colony not found`)
		c.JSON(404, gin.H{`error`: `Colony not found`})
	}

	// Создание
	s.wRep.Create(&body, colonyID, accountID)

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает рабочего по его ID
func (s *WorkerService) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	s.wRep.GetOne(c.Query(`id`), accountID)
}

// GetAll возвращает всех рабочих
func (s *WorkerService) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	query := repositories.WorkerFindAll{}

	if q := c.Query(`usedStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.UsedStorage = &n
	}
	if q := c.Query(`maxStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.MaxStorage = &n
	}
	if q := c.Query(`location`); q != `` {
		query.Location = &q
	}
	if q := c.Query(`fromDeparture`); q != `` {
		n, _ := strconv.Atoi(q)
		query.FromDeparture = &n
	}
	if q := c.Query(`toArrival`); q != `` {
		n, _ := strconv.Atoi(q)
		query.ToArrival = &n
	}

	workers := s.wRep.GetAll(accountID, query)

	c.JSON(200, gin.H{
		`error`:   ``,
		`workers`: workers,
		`count`:   len(workers),
	})
}

// UpdateOne обновляет рабочего
func (s *WorkerService) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Worker
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountId для безопасности
	body.AccountID = accountID

	s.wRep.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет рабочего
func (s *WorkerService) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Worker
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountId для безопасности
	body.AccountID = accountID

	s.wRep.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// Переводит body в структуру
func (s *WorkerService) parseBody(c *gin.Context, body interface{}) error {
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
