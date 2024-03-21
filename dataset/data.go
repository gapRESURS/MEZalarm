package dataset

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var PSWD string

func Data() {

	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Пример чтения переменной окружения
	PSWD = os.Getenv("PSWD")

}
