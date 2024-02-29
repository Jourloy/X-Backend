package colony

import (
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[colony]`,
		Level:  log.DebugLevel,
	})
)

type ColonyService struct {
	cRep  repositories.IColonyRepository
	cache redis.Client
}

// InitColonyService создает сервис колонии
func InitColonyService(cRep repositories.IColonyRepository, cache redis.Client) *ColonyService {

	logger.Info(`ColonyService initialized`)

	return &ColonyService{
		cRep:  cRep,
		cache: cache,
	}
}

// Create создает колонию
func (s *ColonyService) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Colony
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Получение placeID
	q := c.Query(`placeID`)
	if q == `` {
		logger.Error(`placeID is required`)
		c.JSON(400, gin.H{`error`: `placeID is required`})
	}

	s.cRep.Create(&body, accountID, q)
	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает колонию по id
func (s *ColonyService) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)
	s.cRep.GetOne(c.Param(`id`), accountID)
}

// GetAll возвращает все колонии
func (s *ColonyService) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.ColonyFindAll{}
	if q := c.Query(`balance`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Balance = &n
	}
	if q := c.Query(`maxStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.MaxStorage = &n
	}
	if q := c.Query(`usedStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.UsedStorage = &n
	}
	if q := c.Query(`limit`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Limit = &n
	}

	// Получение работников
	colonies := s.cRep.GetAll(accountID, query)
	c.JSON(200, gin.H{
		`error`:    ``,
		`colonies`: colonies,
		`count`:    len(colonies),
	})
}

// UpdateOne обновляет колонию
func (s *ColonyService) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Colony
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.cRep.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет колонию
func (s *ColonyService) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Colony
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.cRep.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}
