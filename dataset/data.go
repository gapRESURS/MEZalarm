package dataset

import (
	"MEZ/logger"
	"github.com/joho/godotenv"
	"os"
)

var PSWD string
var OPERATOR string
var WEB_IP string
var WEB_PORT string

func Data() {
	err := godotenv.Load()
	if err != nil {
		logger.Log("Error loading .env file")
		return
	}

	PSWD = os.Getenv("PSWD")
	OPERATOR = os.Getenv("OPERATOR")
	WEB_IP = os.Getenv("WEB_IP")
	WEB_PORT = os.Getenv("WEB_PORT")

}
