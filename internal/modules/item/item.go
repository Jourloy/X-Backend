package item

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	item_service "github.com/jourloy/X-Backend/internal/modules/item/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[item]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service item_service.Service
}

// Init создает сервис предмета
func Init() *Controller {
	service := item_service.Init()

	return &Controller{
		service: *service,
	}
}

// Create создает предмет
func (s *Controller) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var b repositories.Item
	if err := tools.ParseBody(c, &b); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Получение ID родителя
	pID := c.Query(`parentID`)
	if pID == `` {
		logger.Error(`parentID is required`)
		c.JSON(400, gin.H{`error`: `parentID is required`})
	}

	// Получение тип родителя
	pType := c.Query(`parentType`)
	if pType == `` {
		logger.Error(`parentID is required`)
		c.JSON(400, gin.H{`error`: `parentID is required`})
	}

	resp := s.service.Create(b, pID, pType, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}
