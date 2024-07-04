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
	if err != nil {
		logger.Log("Error RelayStatus() RequestSNMP():", err)
	}
	r, err := parseBool(stR)
	if err != nil {
		logger.Log("Error RelayStatus() parseBool():", err)
	}

	result := map[string]bool{"result": r}
	jsonResult, err := json.Marshal(result)
	if err != nil {
		logger.Log("Error json.Marshal:", err)
		return nil
	}
	return jsonResult
}

func RelayStatus(ip string) (bool, error) {
	logger.Log("Запрос статуса реле", ip)
	stR, err := RequestSNMP(ip, dataset.PSWD, ".1.3.6.1.4.1.40418.2.6.2.2.1.3.1.1", 161)
	if err != nil {
		return false, err
	}
	r, err := parseBool(stR)
	if err != nil {
		return false, err
	}
	return r, nil
}

func relayStateChange(ip string, r bool) {
	logger.Log("Смена статуса реле", ip)
	value := 1
	if r {
		value = 0
	}
	err := snmpSet(ip, dataset.PSWD, ".1.3.6.1.4.1.40418.2.6.2.2.1.3.1.1", value)
	if err != nil {
		logger.Log("snmpSet:", err)
	}
}

func relayControl(command string) {

	switch command {
	case "ALARM.NMEZ HIGH":
		if dataset.NMEZ_IP != "0.0.0.0" {
			r, err := RelayStatus(dataset.NMEZ_IP)
			if err != nil {
				logger.Log("Error RelayStatus():", err)
			} else {
				logger.Log("Статус реле:", r)
				relayStateChange(dataset.NMEZ_IP, r)
			}
		}
	case "ALARM.AMEZ HIGH":
		if dataset.AMEZ_IP != "0.0.0.0" {
			r, err := RelayStatus(dataset.AMEZ_IP)
			if err != nil {
				logger.Log("Error RelayStatus():", err)
			} else {
				logger.Log("Статус реле:", r)
				relayStateChange(dataset.AMEZ_IP, r)
			}
		}
	case "ALARM.GMEZ HIGH":
		if dataset.GMEZ_IP != "0.0.0.0" {
			r, err := RelayStatus(dataset.GMEZ_IP)
			if err != nil {
				logger.Log("Error RelayStatus():", err)
			} else {
				logger.Log("Статус реле:", r)
				relayStateChange(dataset.GMEZ_IP, r)
			}
		}
	case "ALARM.SMEZ HIGH":
		if dataset.SMEZ_IP != "0.0.0.0" {
			r, err := RelayStatus(dataset.SMEZ_IP)
			if err != nil {
				logger.Log("Error RelayStatus():", err)
			} else {
				logger.Log("Статус реле:", r)
				relayStateChange(dataset.SMEZ_IP, r)
			}
		}
	case "ALARM.UMEZ HIGH":
		if dataset.UMEZ_IP != "0.0.0.0" {
			r, err := RelayStatus(dataset.UMEZ_IP)
			if err != nil {
				logger.Log("Error RelayStatus():", err)
			} else {
				logger.Log("Статус реле:", r)
				relayStateChange(dataset.UMEZ_IP, r)
			}
		}
	default:
		logger.Log("IP и пароль верные, текст трапа:", command)
	}

}
