package warrior

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
		Prefix: `[warrior]`,
		Level:  log.DebugLevel,
	})
)

type WarriorService struct {
	wRep  repositories.IWarriorRepository
	cRep  repositories.IColonyRepository
	cache redis.Client
}

// InitWarriorService создает сервис воина
func InitWarriorService(wRep repositories.IWarriorRepository, cRep repositories.IColonyRepository, cache redis.Client) *WarriorService {

	logger.Info(`WarriorService initialized`)

	return &WarriorService{
		wRep:  wRep,
		cRep:  cRep,
		cache: cache,
	}
}

// Create создает воина
func (s *WarriorService) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Warrior
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Проверка существования воина
	colonyID := c.Query(`colonyID`)
	if colonyID == `` {
		logger.Error(`colonyID is required`)
		c.JSON(400, gin.H{`error`: `colonyID is required`})
	}

	colony := s.cRep.GetOne(colonyID, accountID)
	if colony.ID == `` {
		logger.Error(`Colony not found`)
		c.JSON(404, gin.H{`error`: `Colony not found`})
	}

	s.wRep.Create(&body, colonyID, accountID)
	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает воина по его ID
func (s *WarriorService) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	s.wRep.GetOne(c.Query(`id`), accountID)
}

// GetAll возвращает всех воинов
func (s *WarriorService) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.WarriorFindAll{}
	if q := c.Query(`usedStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.UsedStorage = &n
	}
	if q := c.Query(`maxStorage`); q != `` {
		n, _ := strconv.Atoi(q)
		query.MaxStorage = &n
	}
	if q := c.Query(`location`); q != `` {
		query.Location = &q
	}
	if q := c.Query(`fromDeparture`); q != `` {
		n, _ := strconv.Atoi(q)
		query.FromDeparture = &n
	}
	if q := c.Query(`toArrival`); q != `` {
		n, _ := strconv.Atoi(q)
		query.ToArrival = &n
	}
	if q := c.Query(`limit`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Limit = &n
	}
	if q := c.Query(`health`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Health = &n
	}

	// Получение воинов
	warriors := s.wRep.GetAll(accountID, query)

	c.JSON(200, gin.H{
		`error`:    ``,
		`warriors`: warriors,
		`count`:    len(warriors),
	})
}

// UpdateOne обновляет воина
func (s *WarriorService) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Warrior
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.wRep.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет воина
func (s *WarriorService) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Warrior
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountId для безопасности
	body.AccountID = accountID

	s.wRep.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}
