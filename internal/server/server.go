package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

// Server представляет структуру HTTP-сервера
type Server struct {
	logger *log.Logger
	server *http.Server
}

// NewServer создает и конфигурирует новый HTTP-сервер
func NewServer(logger *log.Logger) *Server {
	// Создаем HTTP роутер
	mux := http.NewServeMux()

	// Регистрируем хендлеры
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler(logger))

	// Создаем экземпляр http.Server с настройками
	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		server: httpServer,
	}
}

// Start запускает HTTP-сервер
func (s *Server) Start() error {
	s.logger.Printf("Starting server on %s", s.server.Addr)
	return s.server.ListenAndServe()
}
