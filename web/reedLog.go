package web

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func handlerLog(w http.ResponseWriter, r *http.Request) {
	logFile := "LOG/" + time.Now().Format("2006") + "/" + time.Now().Format("01-02") + ".log"

	dir := filepath.Dir(logFile)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		fmt.Fprintf(w, "Директория %s не существует", dir)
		return
	}

	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		fmt.Fprintf(w, "Ошибка при чтении файла логов: %v", err)
		return
	}

	w.Write(content)
}
