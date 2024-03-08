package internal

import (
	"io"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/handlers"
	"github.com/jourloy/X-Backend/internal/repositories/account_rep"
	"github.com/jourloy/X-Backend/internal/repositories/deposit_rep"
	"github.com/jourloy/X-Backend/internal/repositories/item_rep"
	"github.com/jourloy/X-Backend/internal/repositories/item_template_rep"
	"github.com/jourloy/X-Backend/internal/repositories/market_rep"
	"github.com/jourloy/X-Backend/internal/repositories/node_rep"
	"github.com/jourloy/X-Backend/internal/repositories/plan_rep"
	"github.com/jourloy/X-Backend/internal/repositories/resource_rep"
	"github.com/jourloy/X-Backend/internal/repositories/resource_template_rep"
	"github.com/jourloy/X-Backend/internal/repositories/scout_rep"
	"github.com/jourloy/X-Backend/internal/repositories/sector_rep"
	"github.com/jourloy/X-Backend/internal/repositories/storage_rep"
	"github.com/jourloy/X-Backend/internal/repositories/tower_rep"
	"github.com/jourloy/X-Backend/internal/repositories/townhall_rep"
	"github.com/jourloy/X-Backend/internal/repositories/trader_rep"
	"github.com/jourloy/X-Backend/internal/repositories/warrior_rep"
	"github.com/jourloy/X-Backend/internal/repositories/worker_rep"

	_ "github.com/jourloy/X-Backend/docs"
	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/middlewares"
	"github.com/jourloy/X-Backend/internal/storage"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	logger = log.NewWithOptions(os.Stderr, log.Options{
		Prefix: `[gin]`,
		Level:  log.DebugLevel,
	})
)

func StartServer() {
	totalTime := time.Now()
	tempTime := time.Now()

	gin.DefaultWriter = NewDebugWrite()

	// Инициализация хранилища
	storage.InitDB()
	logger.Debug(`Storage initialized`, `latency`, time.Since(tempTime))
	tempTime = time.Now()

	// Инициализация кэша
	cache.InitCache()
	logger.Debug(`Cache initialized`, `latency`, time.Since(tempTime))
	tempTime = time.Now()

	// Инициализация репозиториев
	initReps()
	logger.Debug(`Repositories initialized`, `latency`, time.Since(tempTime))
	tempTime = time.Now()

	r := gin.New()

	// Middlewares
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Группы
	accountGroup := r.Group(`account`)
	depositGroup := r.Group(`deposit`)
	itemGroup := r.Group(`item`)
	marketGroup := r.Group(`market`)
	planGroup := r.Group(`plan`)
	resourceGroup := r.Group(`resource`)
	scoutGroup := r.Group(`scout`)
	sectorGroup := r.Group(`sector`)
	storageGroup := r.Group(`storage`)
	towerGroup := r.Group(`tower`)
	townhallGroup := r.Group(`townhall`)
	traderGroup := r.Group(`trader`)
	wallGroup := r.Group(`wall`)
	warriorGroup := r.Group(`warrior`)
	workerGroup := r.Group(`worker`)

	// Инициализация групп
	handlers.InitAccount(accountGroup)
	handlers.InitDeposit(depositGroup)
	handlers.InitItem(itemGroup)
	handlers.InitMarket(marketGroup)
	handlers.InitPlan(planGroup)
	handlers.InitResource(resourceGroup)
	handlers.InitScout(scoutGroup)
	handlers.InitSector(sectorGroup)
	handlers.InitStorage(storageGroup)
	handlers.InitTower(towerGroup)
	handlers.InitTownhall(townhallGroup)
	handlers.InitTrader(traderGroup)
	handlers.InitWall(wallGroup)
	handlers.InitWarrior(warriorGroup)
	handlers.InitWorker(workerGroup)

	logger.Debug(`Handlers initialized`, `latency`, time.Since(tempTime))

	// Запуск сервера
	logger.Info(`Server started`, `port`, 3001, `latency (total)`, time.Since(totalTime))
	if err := r.Run(`0.0.0.0:3001`); err != nil {
		log.Fatal(err)
	}
}

// Инициализация репозиториев
func initReps() {
	account_rep.Init()

	sector_rep.Init()
	node_rep.Init()

	item_rep.Init()
	townhall_rep.Init()
	resource_rep.Init()
	trader_rep.Init()
	warrior_rep.Init()
	worker_rep.Init()
	plan_rep.Init()
	deposit_rep.Init()
	storage_rep.Init()
	tower_rep.Init()
	market_rep.Init()
	resource_template_rep.Init()
	item_template_rep.Init()
	scout_rep.Init()
}

type WriteFunc func([]byte) (int, error)

func (fn WriteFunc) Write(data []byte) (int, error) {
	return fn(data)
}

func NewDebugWrite() io.Writer {
	return WriteFunc(func(data []byte) (int, error) {
		// logger.Debugf("%s", data)
		return 0, nil
	})
}
