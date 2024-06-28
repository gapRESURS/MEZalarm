package erd

import (
	"MEZ/dataset"
	"MEZ/logger"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func parserH(s string) (int8, error) {
	if len(s) == 0 || !strings.Contains(s, ":") {
		return 0, fmt.Errorf("некорректная строка")
	}
	parts := strings.Split(s, ":")
	num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, fmt.Errorf("ошибка преобразования: %w", err)
	}
	if num < 0 || num > 100 {
		return 0, fmt.Errorf("число вне диапазона 0, 100")
	}
	return int8(num), nil
}

func parserT(s string) (int8, error) {
	if len(s) == 0 || !strings.Contains(s, ":") {
		return 0, fmt.Errorf("некорректная строка")
	}
	parts := strings.Split(s, ":")
	num, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return 0, fmt.Errorf("ошибка преобразования: %w", err)
	}
	if num < -40 || num > 125 {
		return 0, fmt.Errorf("число вне диапазона -40, 125")
	}
	return int8(num), nil
}

func DHT22Json(ip string) []byte {
	stT, err := RequestSNMP(ip, dataset.PSWD, ".1.3.6.1.4.1.40418.2.6.1.8.0", 161)
	if err != nil {
		logger.Log("Error RequestSNMP:", err)
	}
	temperature, err := parserT(stT)
	if err != nil {
		logger.Log("Error parser temperature:", err)
	}
	stH, err := RequestSNMP(ip, dataset.PSWD, ".1.3.6.1.4.1.40418.2.6.1.9.0", 161)
	if err != nil {
		logger.Log("Error RequestSNMP:", err)
	}
	humidity, err := parserH(stH)
	if err != nil {
		logger.Log("Error parser humidity:", err)
	}

	result := map[string]int8{"temperature": temperature, "humidity": humidity}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		logger.Log("Error json.Marshal:", err)
		return nil
	}
	return jsonResult
}
