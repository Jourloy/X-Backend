package app

import (
	app_service "github.com/jourloy/X-Backend/internal/modules/app/service"
)

type Controller struct {
	service app_service.Service
}

// Init создает сервис приложения
func Init() *Controller {
	service := app_service.Init()

	return &Controller{
		service: *service,
	}
}
