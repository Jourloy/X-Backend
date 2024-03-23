package resource

import (
	"os"

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

	return &Controller{
		service: *service,
	}
}

// Create создает ресурс
func (s *Controller) Create(c *gin.Context) {
	// Парсинг body
	var body repositories.ResourceCreate
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`parse body error`)
		c.JSON(400, gin.H{`error`: `parse body error`})
	}

	resp := s.service.Create(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}
