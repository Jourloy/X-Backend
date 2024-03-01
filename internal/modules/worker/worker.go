package worker

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	worker_service "github.com/jourloy/X-Backend/internal/modules/worker/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[worker]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service worker_service.Service
}

// InitWorkerService создает сервис рабочего
func InitWorkerService(wRep repositories.IWorkerRepository, cRep repositories.IVillageRepository, cache redis.Client) *Controller {

	logger.Info(`Controller initialized`)
	service := worker_service.InitWorkerService(wRep, cRep, cache)

	return &Controller{
		service: *service,
	}
}

// Create создает рабочего
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Worker
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Получение ID поселения
	villageID := c.Query(`villageID`)
	if villageID == `` {
		logger.Error(`villageID is required`)
		c.JSON(400, gin.H{`error`: `villageID is required`})
	}

	resp := s.service.Create(body, accountID, villageID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает рабочего по его ID
func (s *Controller) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Получение ID рабочего
	workerID := c.Query(`workerID`)
	if workerID == `` {
		logger.Error(`workerID is required`)
		c.JSON(400, gin.H{`error`: `workerID is required`})
	}

	resp := s.service.GetOne(workerID, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `worker`: resp.Worker})
}

// GetAll возвращает всех рабочих
func (s *Controller) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
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
	if q := c.Query(`limit`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Limit = &n
	}

	resp := s.service.GetAll(query, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{
		`error`:   ``,
		`workers`: resp.Workers,
		`count`:   len(resp.Workers),
	})
}

// UpdateOne обновляет рабочего
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Worker
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	resp := s.service.UpdateOne(body, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет рабочего
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Worker
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	resp := s.service.DeleteOne(body, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}
