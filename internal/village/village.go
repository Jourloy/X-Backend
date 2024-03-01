package village

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
		Prefix: `[village]`,
		Level:  log.DebugLevel,
	})
)

type VillageService struct {
	cRep  repositories.IVillageRepository
	cache redis.Client
}

// InitVillageService создает сервис поселений
func InitVillageService(cRep repositories.IVillageRepository, cache redis.Client) *VillageService {

	logger.Info(`VillageService initialized`)

	return &VillageService{
		cRep:  cRep,
		cache: cache,
	}
}

// Create создает поселение
func (s *VillageService) Create(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Village
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Получение sectorID
	q := c.Query(`sectorID`)
	if q == `` {
		logger.Error(`sectorID is required`)
		c.JSON(400, gin.H{`error`: `sectorID is required`})
	}

	s.cRep.Create(&body, accountID, q)
	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает поселение по id
func (s *VillageService) GetOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)
	s.cRep.GetOne(c.Param(`id`), accountID)
}

// GetAll возвращает все поселения
func (s *VillageService) GetAll(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Создание фильтров
	query := repositories.VillageFindAll{}
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

	// Получение поселений
	villages := s.cRep.GetAll(accountID, query)
	c.JSON(200, gin.H{
		`error`:    ``,
		`villages`: villages,
		`count`:    len(villages),
	})
}

// UpdateOne обновляет поселение
func (s *VillageService) UpdateOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Village
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.cRep.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет поселение
func (s *VillageService) DeleteOne(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	// Парсинг body
	var body repositories.Village
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	// Перезапись accountID для безопасности
	body.AccountID = accountID

	s.cRep.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}
