package scout

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	scout_service "github.com/jourloy/X-Backend/internal/modules/scout/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[scout]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service scout_service.Service
}

// Init создает контроллер разведчика
func Init() *Controller {
	service := scout_service.Init()
	logger.Info(`Controller initialized`)
	return &Controller{
		service: *service,
	}
}

// Create создает разведчика
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Scout
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	resp := s.service.Create(body, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает разведчика по его ID
func (s *Controller) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Получение ID разведчика
	scoutID := c.Query(`scoutID`)
	if scoutID == `` {
		logger.Error(`scoutID is required`)
		c.JSON(400, gin.H{`error`: `scoutID is required`})
	}

	resp := s.service.GetOne(scoutID, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `scout`: resp.Scout})
}

// GetAll возвращает всех разведчиков
func (s *Controller) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.ScoutGetAll{}
	if q := c.Query(`usedStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.UsedStorage = &n
	}
	if q := c.Query(`maxStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.MaxStorage = &n
	}
	if q := c.Query(`x`); q != `` {
		n, _ := strconv.Atoi(q)
		query.X = &n
	}
	if q := c.Query(`y`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Y = &n
	}
	if q := c.Query(`maxHealth`); q != `` {
		n, _ := strconv.Atoi(q)
		query.MaxHealth = &n
	}
	if q := c.Query(`health`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Health = &n
	}
	if q := c.Query(`requireCoins`); q != `` {
		n, _ := strconv.ParseFloat(q, 64)
		query.RequireCoins = &n
	}
	if q := c.Query(`requireFood`); q != `` {
		n, _ := strconv.ParseFloat(q, 64)
		query.RequireFood = &n
	}
	if q := c.Query(`fatigue`); q != `` {
		n, _ := strconv.ParseFloat(q, 64)
		query.Fatigue = &n
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
		`error`:  ``,
		`scouts`: resp.Scouts,
		`count`:  len(resp.Scouts),
	})
}

// UpdateOne обновляет разведчика
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Scout
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

// DeleteOne удаляет разведчика
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Scout
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
