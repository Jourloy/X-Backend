package building

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	building_service "github.com/jourloy/X-Backend/internal/modules/building/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[building]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service building_service.Service
}

// Init создает контроллер постройки
func Init() *Controller {
	service := building_service.Init()

	return &Controller{
		service: *service,
	}
}

// Create создает постройку
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.BuildingCreate
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`parse body error`)
		c.JSON(400, gin.H{`error`: `parse body error`})
		return
	}

	body.AccountID = accountID

	resp := s.service.Create(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(resp.Code, gin.H{`error`: resp.Err.Error()})
		return
	}

	c.JSON(resp.Code, gin.H{`error`: ``, `building`: resp.Building})
}

// GetOne получает постройку по его ID
func (s *Controller) GetOne(c *gin.Context) {
	// Парсинг body
	var body repositories.BuildingGet
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`parse body error`)
		c.JSON(400, gin.H{`error`: `parse body error`})
		return
	}

	resp := s.service.GetOne(&body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
		return
	}

	c.JSON(200, gin.H{`error`: ``, `worker`: resp.Building})
}

// GetAll возвращает все постройки
func (s *Controller) GetAll(c *gin.Context) {
	// Парсинг body
	var body repositories.BuildingGet
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`parse body error`)
		c.JSON(400, gin.H{`error`: `parse body error`})
		return
	}

	resp := s.service.GetAll(&body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
		return
	}

	c.JSON(200, gin.H{
		`error`:     ``,
		`buildings`: resp.Buildings,
		`count`:     len(resp.Buildings),
	})
}

// UpdateOne обновляет постройку
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Building
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`parse body error`)
		c.JSON(400, gin.H{`error`: `parse body error`})
		return
	}

	body.AccountID = accountID

	resp := s.service.UpdateOne(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
		return
	}

	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет постройку
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Building
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`parse body error`)
		c.JSON(400, gin.H{`error`: `parse body error`})
		return
	}

	body.AccountID = accountID

	resp := s.service.DeleteOne(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
		return
	}

	c.JSON(200, gin.H{`error`: ``})
}
