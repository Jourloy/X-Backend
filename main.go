package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"

	"github.com/jourloy/X-Backend/internal"
	"github.com/jourloy/X-Backend/internal/config"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear[`linux`] = func() {
		cmd := exec.Command(`clear`)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear[`darwin`] = func() {
		cmd := exec.Command(`clear`)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear[`windows`] = func() {
		cmd := exec.Command(`cmd`, `/c`, `cls`)
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	// Очистка терминала
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic(`Your platform is unsupported`)
	}

	// Загрузка .env
	if err := godotenv.Load(`.env`); err != nil {
		log.Fatal(`Error loading .env file`)
	}

	// Парсинг .env
	config.ParseENV()

	// Старт сервера
	internal.StartServer()
}
