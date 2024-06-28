package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

//func Log(text string) {
//	dt := time.Now()
//	file := "LOG/" + dt.Format("2006") + "/" + dt.Format("01-02") + ".log"
//	dirPath := filepath.Dir(file)
//	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
//		log.Fatal(err)
//	}
//
//	logFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatal("Ошибка открытия файла логов:", err)
//	}
//	defer logFile.Close()
//
//	log.SetOutput(logFile)
//	log.Println(text)
//}

// Log записывает информацию в файл логов с указанными аргументами.
//func Log(args ...interface{}) {
//	dt := time.Now()
//	fileName := "LOG/" + dt.Format("2006") + "/" + dt.Format("01-02") + ".log"
//	dirPath := filepath.Dir(fileName)
//	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
//		log.Fatalf("Ошибка создания директории: %v", err)
//	}
//
//	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
//	if err != nil {
//		log.Fatalf("Ошибка открытия файла логов: %v", err)
//	}
//	defer logFile.Close()
//
//	var message strings.Builder // Используйте StringBuilder для сборки сообщения
//
//	for _, arg := range args {
//		switch v := arg.(type) {
//		case string:
//			message.WriteString(v + ": ")
//		case error:
//			message.WriteString(fmt.Sprintf("%s: %v\n", v, v)) // Добавляем \n для перевода строки
//		default:
//			message.WriteString(fmt.Sprintf("%v\n", v)) // Добавляем \n для перевода строки
//		}
//	}
//
//	// Записываем собранное сообщение в файл
//	err = ioutil.WriteFile(fileName, []byte(message.String()), 0666)
//	if err != nil {
//		log.Fatalf("Ошибка записи в файл логов: %v", err)
//	}
//}

func Log(v ...interface{}) {
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
	log.Println(v...)
	fmt.Println(v...)
}
