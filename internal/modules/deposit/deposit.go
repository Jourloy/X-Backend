package deposit

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	deposit_service "github.com/jourloy/X-Backend/internal/modules/deposit/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[deposit]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service deposit_service.Service
}

// Init создает контроллер залежи
func Init() *Controller {
	service := deposit_service.Init()

	return &Controller{
		service: *service,
	}
}

// Create создает залежь
func (s *Controller) Create(c *gin.Context) {
	// Парсинг body
	var body repositories.DepositCreate
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
