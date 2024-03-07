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

// GetOne получает предмет по id
func (s *Controller) GetOne(c *gin.Context) {
	// Получение ID рабочего
	id := c.Query(`itemID`)
	if id == `` {
		logger.Error(`itemID is required`)
		c.JSON(400, gin.H{`error`: `itemID is required`})
	}

	resp := s.service.GetOne(id)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `item`: resp.Item})
}

// GetAll возвращает все пердметы
func (s *Controller) GetAll(c *gin.Context) {
	// Создание фильтров
	query := repositories.ItemGetAll{}
	if q := c.Query(`type`); q != `` {
		query.Type = &q
	}
	if q := c.Query(`x`); q != `` {
		n, _ := strconv.Atoi(q)
		query.X = &n
	}
	if q := c.Query(`y`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Y = &n
	}
	if q := c.Query(`parentId`); q != `` {
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
		`error`: ``,
		`items`: resp.Items,
		`count`: len(resp.Items),
	})
}

// UpdateOne обновляет пердмет
func (s *Controller) UpdateOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Item
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

// DeleteOne удаляет предмет
func (s *Controller) DeleteOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Item
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
