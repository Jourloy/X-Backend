package resource

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	resource_service "github.com/jourloy/X-Backend/internal/modules/resource/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[resource]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service resource_service.ResourceService
}

// Init создает сервис ресурса
func Init() *Controller {
	service := resource_service.Init()
	logger.Info(`Controller initialized`)
	return &Controller{
		service: *service,
	}
}

// Create создает ресурс
func (s *Controller) Create(c *gin.Context) {
	// Парсинг body
	var body repositories.Resource
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

// GetOne получает ресурс по id
func (s *Controller) GetOne(c *gin.Context) {
	// Получение ID ресурса
	resourceID := c.Query(`resourceID`)
	if resourceID == `` {
		logger.Error(`resourceID is required`)
		c.JSON(400, gin.H{`error`: `resourceID is required`})
	}

	resp := s.service.GetOne(resourceID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `resource`: resp.Resource})
}

// GetAll возвращает все ресурсы
func (s *Controller) GetAll(c *gin.Context) {

	// Создание фильтров
	query := repositories.ResourceGetAll{}
	if q := c.Query(`type`); q != `` {
		query.Type = &q
	}
	if q := c.Query(`amount`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Amount = &n
	}
	if q := c.Query(`weight`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Weight = &n
	}
	if q := c.Query(`x`); q != `` {
		n, _ := strconv.Atoi(q)
		query.X = &n
	}
	if q := c.Query(`y`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Y = &n
	}
	if q := c.Query(`parentID`); q != `` {
		query.ParentID = &q
	}
	if q := c.Query(`parentType`); q != `` {
		query.ParentType = &q
	}
	if q := c.Query(`sectorId`); q != `` {
		query.SectorID = &q
	}
	if q := c.Query(`creatorId`); q != `` {
		query.CreatorID = &q
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
		`error`:     ``,
		`resources`: resp.Resources,
		`count`:     len(resp.Resources),
	})
}

// UpdateOne обновляет ресурс
func (s *Controller) UpdateOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Resource
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

// DeleteOne удаляет ресурс
func (s *Controller) DeleteOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Resource
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
