package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

func Log(text string) {
	dt := time.Now()
	file := "LOG/" + dt.Format("2006") + "/" + dt.Format("01-02") + ".log"
	dirPath := filepath.Dir(file)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Ошибка открытия файла логов:", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.Println(text)
}
