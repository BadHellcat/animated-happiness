package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// ConvertData автоматически определяет тип данных (текст или код Морзе) и конвертирует их.
// Если передан обычный текст, функция конвертирует его в код Морзе.
// Если передан код Морзе, функция конвертирует его в обычный текст.
func ConvertData(data string) (string, error) {
	if data == "" {
		return "", errors.New("empty data")
	}

	// Определяем, является ли строка кодом Морзе
	// Код Морзе содержит только точки, тире и пробелы
	isMorse := true
	for _, char := range data {
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\r' && char != '\t' {
			isMorse = false
			break
		}
	}

	// Убираем лишние пробелы и переносы строк
	data = strings.TrimSpace(data)

	if isMorse {
		// Конвертируем код Морзе в текст
		result := morse.ToText(data)
		return result, nil
	}

	// Конвертируем текст в код Морзе
	result := morse.ToMorse(data)
	return result, nil
}
