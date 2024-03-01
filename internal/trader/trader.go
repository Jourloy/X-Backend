package trader

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
		Prefix: `[trader]`,
		Level:  log.DebugLevel,
	})
)

type TraderService struct {
	wRep  repositories.ITraderRepository
	cRep  repositories.IVillageRepository
	cache redis.Client
}

// InitTraderService создает сервис торговца
func InitTraderService(wRep repositories.ITraderRepository, cRep repositories.IVillageRepository, cache redis.Client) *TraderService {

	logger.Info(`TraderService initialized`)

	return &TraderService{
		wRep:  wRep,
		cRep:  cRep,
		cache: cache,
	}
}

// Create создает торговца
func (s *TraderService) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Trader
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Проверка существования поселения
	villageID := c.Query(`villageID`)
	if villageID == `` {
		logger.Error(`villageID is required`)
		c.JSON(400, gin.H{`error`: `villageID is required`})
	}

	village := s.cRep.GetOne(villageID, accountID)
	if village.ID == `` {
		logger.Error(`Village not found`)
		c.JSON(404, gin.H{`error`: `Village not found`})
	}

	s.wRep.Create(&body, villageID, accountID)
	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает торговца по его ID
func (s *TraderService) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	s.wRep.GetOne(c.Query(`id`), accountID)
}

// GetAll возвращает всех торговцев
func (s *TraderService) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.TraderFindAll{}
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

	// Получение торговцев
	traders := s.wRep.GetAll(accountID, query)

	c.JSON(200, gin.H{
		`error`:   ``,
		`traders`: traders,
		`count`:   len(traders),
	})
}

// UpdateOne обновляет торговца
func (s *TraderService) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Trader
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.wRep.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет торговца
func (s *TraderService) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Trader
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountId для безопасности
	body.AccountID = accountID

	s.wRep.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}
