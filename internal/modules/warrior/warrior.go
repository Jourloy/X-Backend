package warrior

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	warrior_service "github.com/jourloy/X-Backend/internal/modules/warrior/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[warrior]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service warrior_service.Service
}

// InitWarriorService создает сервис воина
func InitWarriorService(wRep repositories.IWarriorRepository, cRep repositories.IVillageRepository, cache redis.Client) *Controller {

	logger.Info(`Controller initialized`)
	service := warrior_service.InitWarriorService(wRep, cRep, cache)

	return &Controller{
		service: *service,
	}
}

// Create создает воина
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Warrior
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

// GetOne получает воина по его ID
func (s *Controller) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Получение ID рабочего
	warriorID := c.Query(`warriorID`)
	if warriorID == `` {
		logger.Error(`warriorID is required`)
		c.JSON(400, gin.H{`error`: `warriorID is required`})
	}

	resp := s.service.GetOne(warriorID, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `worker`: resp.Warrior})
}

// GetAll возвращает всех воинов
func (s *Controller) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.WarriorFindAll{}
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
	if q := c.Query(`health`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Health = &n
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
		`error`:    ``,
		`warriors`: resp.Warriors,
		`count`:    len(resp.Warriors),
	})
}

// UpdateOne обновляет воина
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Warrior
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

// DeleteOne удаляет воина
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Warrior
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
