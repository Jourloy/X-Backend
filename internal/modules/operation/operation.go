package operation

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	operation_service "github.com/jourloy/X-Backend/internal/modules/operation/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[operation]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service operation_service.Service
}

// Init создает контроллер операции
func Init() *Controller {
	service := operation_service.Init()

	return &Controller{
		service: *service,
	}
}

// Create создает операцию
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.OperationCreate
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

	c.JSON(resp.Code, gin.H{`error`: ``, `operation`: resp.Operation})
}

// GetOne получает операцию по его ID
func (s *Controller) GetOne(c *gin.Context) {
	// Парсинг body
	var body repositories.OperationGet
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

	c.JSON(200, gin.H{`error`: ``, `worker`: resp.Operation})
}

// GetAll возвращает все операции
func (s *Controller) GetAll(c *gin.Context) {
	// Парсинг body
	var body repositories.OperationGet
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
		`error`:      ``,
		`operations`: resp.Operations,
		`count`:      len(resp.Operations),
	})
}

// UpdateOne обновляет операцию
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Operation
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

// DeleteOne удаляет операцию
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Operation
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
