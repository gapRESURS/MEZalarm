package erd

import (
	"MEZ/dataset"
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

func DHT22(ip string) (temperature, humidity int8) {
	stT, _ := RequestSNMP(ip, dataset.PSWD, ".1.3.6.1.4.1.40418.2.6.1.8.0", 161)
	temperature, _ = parserT(stT)
	stH, _ := RequestSNMP(ip, dataset.PSWD, ".1.3.6.1.4.1.40418.2.6.1.9.0", 161)
	humidity, _ = parserH(stH)
	return temperature, humidity
}
