package erd

import (
	"MEZ/dataset"
	"MEZ/logger"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func parseBool(s string) (bool, error) {
	if !strings.Contains(s, ":") {
		return false, fmt.Errorf("некорректная строка: отсутствует символ ':'")
	}

	parts := strings.Split(s, ":")
	if len(parts) < 2 {
		return false, fmt.Errorf("некорректная строка: недостаточно частей после разделения")
	}

	parts[1] = strings.TrimRightFunc(parts[1], func(r rune) bool { return r == '\n' })
	parts[1] = strings.TrimSpace(parts[1])

	num, err := strconv.Atoi(parts[1])
	if err != nil {
		return false, fmt.Errorf("ошибка преобразования: %w", err)
	}

	if num != 0 && num != 1 {
		return false, fmt.Errorf("значение должно быть 0 или 1")
	}
	return num == 1, nil
}

func RelayStatusJson(ip string) []byte {
	stR, err := RequestSNMP(ip, dataset.PSWD, ".1.3.6.1.4.1.40418.2.6.2.2.1.3.1.1", 161)
	logger.Log("RelayStatus() RequestSNMP():", err)
	r, err := parseBool(stR)
	logger.Log("RelayStatus() parseBool():", err)

	result := map[string]bool{"result": r}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return jsonResult
}
