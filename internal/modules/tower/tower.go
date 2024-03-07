package tower

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	tower_service "github.com/jourloy/X-Backend/internal/modules/tower/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[tower]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service tower_service.Service
}

// Init создает контроллер башни
func Init() *Controller {
	service := tower_service.Init()

	return &Controller{
		service: *service,
	}
}

// Create создает башню
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Tower
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

// GetOne получает башню по его ID
func (s *Controller) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Получение ID башни
	towerID := c.Query(`towerID`)
	if towerID == `` {
		logger.Error(`towerID is required`)
		c.JSON(400, gin.H{`error`: `towerID is required`})
	}

	resp := s.service.GetOne(towerID, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `tower`: resp.Tower})
}

// GetAll возвращает все башни
func (s *Controller) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.TowerGetAll{}
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
		`error`:  ``,
		`towers`: resp.Towers,
		`count`:  len(resp.Towers),
	})
}

// UpdateOne обновляет башню
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Tower
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

// DeleteOne удаляет башню
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Tower
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
