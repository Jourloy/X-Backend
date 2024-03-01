package sector

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"github.com/jourloy/X-Backend/internal/repositories"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[sector]`,
		Level:  log.DebugLevel,
	})
)

type SectorService struct {
	db    repositories.ISectorRepository
	cache redis.Client
}

// InitSectorService создает сервис сектора
func InitSectorService(db repositories.ISectorRepository, cache redis.Client) *SectorService {

	logger.Info(`SectorService initialized`)

	return &SectorService{
		db:    db,
		cache: cache,
	}
}

// Create создает сектор
func (s *SectorService) Create(c *gin.Context) {
	var body repositories.Sector
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	s.db.Create(&body)

	c.JSON(200, gin.H{`error`: ``})
}

// GetOne получает сектор по id
func (s *SectorService) GetOne(c *gin.Context) {
	s.db.GetOne(c.Param(`id`))
}

// GetAll возвращает все сектора
func (s *SectorService) GetAll(c *gin.Context) {

	// Создание фильтров
	query := repositories.SectorFindAll{}
	if q := c.Query(`limit`); q != `` {
		n, _ := strconv.Atoi(q)
		query.Limit = &n
	}

	// Получение секторов
	sectors := s.db.GetAll(query)

	c.JSON(200, gin.H{
		`error`:   ``,
		`sectors`: sectors,
		`count`:   len(sectors),
	})
}

// UpdateOne обновляет сектор
func (s *SectorService) UpdateOne(c *gin.Context) {
	var body repositories.Sector
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.db.GetOne(body.ID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.db.UpdateOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет сектор
func (s *SectorService) DeleteOne(c *gin.Context) {
	var body repositories.Sector
	if err := s.parseBody(c, &body); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	model := s.db.GetOne(body.ID)
	if model.ID != body.ID {
		logger.Error(`Model not found`)
		c.JSON(404, gin.H{`error`: `Model not found`})
	}

	s.db.DeleteOne(&body)
	c.JSON(200, gin.H{`error`: ``})
}

func (s *SectorService) parseBody(c *gin.Context, body interface{}) error {
	// Проверка body
	if c.Request.Body == nil {
		return errors.New(`body not found`)
	}

	// Чтение body
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	// Парсинг
	if err := json.Unmarshal(b, &body); err != nil {
		return err
	}

	return nil
}
