package internal

import (
	"io"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/deposit_rep"
	"github.com/jourloy/X-Backend/internal/repositories/item_rep"
	"github.com/jourloy/X-Backend/internal/repositories/market_rep"
	"github.com/jourloy/X-Backend/internal/repositories/plan_rep"
	"github.com/jourloy/X-Backend/internal/repositories/resource_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/storage_rep"
	"github.com/jourloy/X-Backend/internal/repositories/tower_rep"
	"github.com/jourloy/X-Backend/internal/repositories/townhall_rep"
	"github.com/jourloy/X-Backend/internal/repositories/trader_rep"
	"github.com/jourloy/X-Backend/internal/repositories/warrior_rep"
	"github.com/jourloy/X-Backend/internal/repositories/worker_rep"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/middlewares"
	"github.com/jourloy/X-Backend/internal/storage"
)

func StartServer() {
	gin.DefaultWriter = NewDebugWrite()

	// Инициализация модулей
	storage.InitDB()
	cache.InitCache()

	// Инициализация репозиториев
	initReps()

	r := gin.New()

	// Middlewares
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())

	// Запуск сервера
	log.Info(`Server started on port 10000`)
	if err := r.Run(`0.0.0.0:10000`); err != nil {
		log.Fatal(err)
	}
}

// Инициализация репозиториев
func initReps() {
	account_rep.Init()
	item_rep.Init()
	resource_rep.Init()
	sector_rep.Init()
	trader_rep.Init()
	warrior_rep.Init()
	worker_rep.Init()
	plan_rep.Init()
	deposit_rep.Init()
	storage_rep.Init()
	tower_rep.Init()
	townhall_rep.Init()
	market_rep.Init()
}

type WriteFunc func([]byte) (int, error)

func (fn WriteFunc) Write(data []byte) (int, error) {
	return fn(data)
}

func NewDebugWrite() io.Writer {
	return WriteFunc(func(data []byte) (int, error) {
		// log.Debugf("%s", data)
		return 0, nil
	})
}
