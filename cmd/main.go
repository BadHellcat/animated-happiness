package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "morse-server: ", log.LstdFlags)

	// Создаем сервер
	srv := server.NewServer(logger)

	// Запускаем сервер
	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}
}
