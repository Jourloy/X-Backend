package internal

import (
	"io"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"

	"github.com/jourloy/X-Backend/internal/cache"
	"github.com/jourloy/X-Backend/internal/middlewares"
	"github.com/jourloy/X-Backend/internal/storage"
)

func StartServer() {
	gin.DefaultWriter = NewDebugWrite()

	// Инициализация модулей
	storage.InitDB()
	cache.InitCache()

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
