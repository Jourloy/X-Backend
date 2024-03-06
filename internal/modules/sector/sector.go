package sector

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	sector_service "github.com/jourloy/X-Backend/internal/modules/sector/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[sector]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service sector_service.Service
}

// InitSectorService создает сервис сектора
func InitSectorService() *Controller {

	service := sector_service.InitSectorService()

	logger.Info(`Controller initialized`)

	return &Controller{
		service: *service,
	}
}

// Create создает сектор
func (s *Controller) Create(c *gin.Context) {
	// Парсинг body
	var body repositories.Sector
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	resp := s.service.Create(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает сектор по id
func (s *Controller) GetOne(c *gin.Context) {
	// Получение ID сектора
	sectorID := c.Query(`sectorID`)
	if sectorID == `` {
		logger.Error(`sectorID is required`)
		c.JSON(400, gin.H{`error`: `sectorID is required`})
	}

	resp := s.service.GetOne(sectorID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `sector`: resp.Sector})
}

// GetAll возвращает все сектора
func (s *Controller) GetAll(c *gin.Context) {

	// Создание фильтров
	query := repositories.SectorGetAll{}
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

	resp := s.service.GetAll(query)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{
		`error`:   ``,
		`sectors`: resp.Sectors,
		`count`:   len(resp.Sectors),
	})
}

// UpdateOne обновляет сектор
func (s *Controller) UpdateOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Sector
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	resp := s.service.UpdateOne(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет сектор
func (s *Controller) DeleteOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Sector
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	resp := s.service.DeleteOne(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}
