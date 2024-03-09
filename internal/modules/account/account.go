package account

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	account_service "github.com/jourloy/X-Backend/internal/modules/account/service"
	"github.com/jourloy/X-Backend/internal/repositories"
	"github.com/jourloy/X-Backend/internal/tools"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[account]`,
		Level:  log.DebugLevel,
	})
)

type Controller struct {
	service account_service.Service
}

// Init создает сервис аккаунта
func Init() *Controller {
	service := account_service.Init()

	return &Controller{
		service: *service,
	}
}

type CreateResponse200 struct {
	Error   string               `json:"error"`
	Account repositories.Account `json:"account"`
}

type CreateResponse400 struct {
	Error string `json:"error"`
}

// Create создает аккаунт
func (s *Controller) Create(c *gin.Context) {
	var body repositories.AccountCreate
	if err := tools.ParseBody(c, &body); err != nil {
		logger.Error(`parse body error`)
		c.JSON(400, CreateResponse400{Error: `parse body error`})
		return
	}

	resp := s.service.Create(body)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, CreateResponse400{Error: resp.Err.Error()})
		return
	}

	c.JSON(200, CreateResponse200{Account: *resp.Account})
}

type GetResponse200 struct {
	Error   string               `json:"error"`
	Account repositories.Account `json:"account"`
}

type GetResponse400 struct {
	Error string `json:"error"`
}

// GetMe получает аккаунт авторизованного пользователя
func (s *Controller) GetMe(c *gin.Context) {
	accountID := c.GetString(`accountID`)

	resp := s.service.GetOne(accountID)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, GetResponse400{Error: resp.Err.Error()})
		return
	}

	c.JSON(200, GetResponse200{Account: resp.Account})
}

// UpdateOne обновляет аккаунт
func (s *Controller) UpdateOne(c *gin.Context) {
	var b repositories.Account
	if err := tools.ParseBody(c, &b); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	resp := s.service.UpdateOne(b)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}

// DeleteOne удаляет аккаунт
func (s *Controller) DeleteOne(c *gin.Context) {
	var b repositories.Account
	if err := tools.ParseBody(c, &b); err != nil {
		logger.Error(`Parse body error`)
		c.JSON(400, gin.H{`error`: `Parse body error`})
	}

	resp := s.service.DeleteOne(b)
	if resp.Err != nil {
		logger.Error(resp.Err)
		c.JSON(400, gin.H{`error`: resp.Err.Error()})
	}

	c.JSON(200, gin.H{`error`: ``})
}
