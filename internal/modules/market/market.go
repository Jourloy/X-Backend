package market

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	market_service "github.com/jourloy/X-Backend/internal/modules/market/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[market]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service market_service.Service
}

// Init создает контроллер рынка
func Init() *Controller {
	service := market_service.Init()

	logger.Info(`Market controller initialized`)

	return &Controller{
		service: *service,
	}
}

// Create создает рынок
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.MarketCreate
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	body.AccountID = accountID

	resp := s.service.Create(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает рынок по его ID
func (s *Controller) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Получение ID рынка
	marketID := c.Query(`marketID`)
	if marketID == `` {
		logger.Error(`marketID is required`)
		c.JSON(400, gin.H{`error`: `marketID is required`})
	}

	resp := s.service.GetOne(marketID, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `market`: resp.Market})
}

// GetAll возвращает все рынки
func (s *Controller) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.MarketGetAll{}
	if q := c.Query(`maxDurability`); q != `` {
		n, _ := strconv.Atoi(q)
		query.MaxDurability = &n
	}
	if q := c.Query(`durability`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Durability = &n
	}
	if q := c.Query(`level`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Level = &n
	}
	if q := c.Query(`maxStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.MaxStorage = &n
	}
	if q := c.Query(`usedStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.UsedStorage = &n
	}
	if q := c.Query(`x`); q != `` {
		n, _ := strconv.Atoi(q)
		query.X = &n
	}
	if q := c.Query(`y`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Y = &n
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
		`markets`: resp.Markets,
		`count`:   len(resp.Markets),
	})
}

// UpdateOne обновляет рынок
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Market
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

// DeleteOne удаляет рынок
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Market
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
