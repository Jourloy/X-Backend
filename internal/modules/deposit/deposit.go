package deposit

import (
	"os"
	"strconv"

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
	var body repositories.Deposit
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

// GetOne получает залежб по ее ID
func (s *Controller) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Получение ID рабочего
	depositID := c.Query(`depositID`)
	if depositID == `` {
		logger.Error(`depositID is required`)
		c.JSON(400, gin.H{`error`: `depositID is required`})
	}

	resp := s.service.GetOne(depositID, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``, `deposit`: resp.Deposit})
}

// GetAll возвращает всех рабочих
func (s *Controller) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.DepositGetAll{}
	if q := c.Query(`type`); q != `` {
		query.Type = &q
	}
	if q := c.Query(`amount`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Amount = &n
	}

	resp := s.service.GetAll(query, accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{
		`error`:   ``,
		`workers`: resp.Deposits,
		`count`:   len(resp.Deposits),
	})
}

// UpdateOne обновляет залежь
func (s *Controller) UpdateOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Deposit
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

// DeleteOne удаляет залежь
func (s *Controller) DeleteOne(c *gin.Context) {
	// Парсинг body
	var body repositories.Deposit
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
