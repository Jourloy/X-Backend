package wall

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	wall_service "github.com/jourloy/X-Backend/internal/modules/wall/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[wall]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service wall_service.Service
}

// Init создает контроллер стены
func Init() *Controller {
	service := wall_service.Init()

	logger.Info(`Wall controller initialized`)

	return &Controller{
		service: *service,
	}
}

// Create создает стену
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Wall
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

// GetOne получает стену по его ID
func (s *Controller) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Получение ID рабочего
	wallID := c.Query(`wallID`)
	if wallID == `` {
		logger.Error(`wallID is required`)
		c.JSON(400, gin.H{`error`: `wallID is required`})
	}

	resp := s.service.GetOne(wallID, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `wall`: resp.Wall})
}

// GetAll возвращает все стены
func (s *Controller) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.WallGetAll{}
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
	if q := c.Query(`y`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Y = &n
	}
	if q := c.Query(`x`); q != `` {
		n, _ := strconv.Atoi(q)
		query.X = &n
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
		`error`: ``,
		`walls`: resp.Walls,
		`count`: len(resp.Walls),
	})
}

// UpdateOne обновляет стену
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Wall
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

// DeleteOne удаляет стену
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Wall
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
