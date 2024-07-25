package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func Clearing() {
	Log("Поск прошлогодних логов:", "START")
	lastYear := time.Now().Year() - 1

	logDir := filepath.Join("LOG", fmt.Sprintf("%d", lastYear))

	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		Log("Директория с логами прошлого года не найдена:", logDir)
		return
	}

	err := os.RemoveAll(logDir)
	if err != nil {
		Log("Ошибка при удалении директории с логами:", err.Error())
	} else {
		Log("Успешно удалена директория с логами:", logDir)
	}

	Log("Удаление прошлогодних логов:", "FINISH")
}
