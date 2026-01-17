package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler возвращает HTML-файл с формой для загрузки файлов
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файлов и конвертирует данные
func UploadHandler(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Парсим форму
		err := r.ParseMultipartForm(10 << 20) // 10 MB максимум
		if err != nil {
			logger.Printf("Error parsing form: %v", err)
			http.Error(w, "Error parsing form", http.StatusInternalServerError)
			return
		}

		// Получаем файл из формы
		file, header, err := r.FormFile("myFile")
		if err != nil {
			logger.Printf("Error retrieving file: %v", err)
			http.Error(w, "Error retrieving file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Читаем данные из файла
		data, err := io.ReadAll(file)
		if err != nil {
			logger.Printf("Error reading file: %v", err)
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		// Конвертируем данные
		result, err := service.ConvertData(string(data))
		if err != nil {
			logger.Printf("Error converting data: %v", err)
			http.Error(w, "Error converting data", http.StatusInternalServerError)
			return
		}

		// Генерируем имя для нового файла
		ext := filepath.Ext(header.Filename)
		fileName := time.Now().UTC().String() + ext

		// Создаем локальный файл
		outputFile, err := os.Create(fileName)
		if err != nil {
			logger.Printf("Error creating file: %v", err)
			http.Error(w, "Error creating file", http.StatusInternalServerError)
			return
		}
		defer outputFile.Close()

		// Записываем результат в файл
		_, err = outputFile.WriteString(result)
		if err != nil {
			logger.Printf("Error writing to file: %v", err)
			http.Error(w, "Error writing to file", http.StatusInternalServerError)
			return
		}

		logger.Printf("File converted successfully: %s", fileName)

		// Возвращаем результат клиенту
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(result))
		if err != nil {
			logger.Printf("Error writing response: %v", err)
		}
	}
}
