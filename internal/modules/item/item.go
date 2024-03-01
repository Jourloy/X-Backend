package item

import (
	"os"
	"strconv"

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

// InitItemService создает сервис вещи
func InitItemService() *Controller {
	service := item_service.Init()
	logger.Info(`Controller initialized`)
	return &Controller{
		service: *service,
	}
}

// Create создает вещь
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

// GetOne получает вещь по id
func (s *Controller) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Получение ID рабочего
	id := c.Query(`itemID`)
	if id == `` {
		logger.Error(`itemID is required`)
		c.JSON(400, gin.H{`error`: `itemID is required`})
	}

	resp := s.service.GetOne(id, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `item`: resp.Item})
}

// GetAll возвращает все вещи
func (s *Controller) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.ItemFindAll{}
	if q := c.Query(`limit`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Limit = &n
	}
	if q := c.Query(`type`); q != `` {
		query.Type = &q
	}
	if q := c.Query(`parentID`); q != `` {
		query.ParentID = &q
	}

	resp := s.service.GetAll(query, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{
		`error`: ``,
		`items`: resp.Items,
		`count`: len(resp.Items),
	})
}

// UpdateOne обновляет вещь
func (s *Controller) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Item
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

// DeleteOne удаляет вещь
func (s *Controller) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Item
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
