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
var NMEZ_IP string
var AMEZ_IP string
var GMEZ_IP string
var SMEZ_IP string
var UMEZ_IP string

func Data() {
	err := godotenv.Load()
	if err != nil {
		logger.Log("Error loading .env file")
		return
	}

	PSWD = "OTOvK!T0"
	OPERATOR = os.Getenv("OPERATOR")
	WEB_IP = os.Getenv("WEB_IP")
	WEB_PORT = os.Getenv("WEB_PORT")
	NMEZ_IP = os.Getenv("NMEZ_IP")
	AMEZ_IP = os.Getenv("AMEZ_IP")
	GMEZ_IP = os.Getenv("GMEZ_IP")
	SMEZ_IP = os.Getenv("SMEZ_IP")
	UMEZ_IP = os.Getenv("UMEZ_IP")

}
