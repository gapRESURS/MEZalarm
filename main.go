package main

import (
	"MEZ/dataset"
	"MEZ/erd"
	"MEZ/logger"
	"MEZ/web"
)

func main() {
	logger.Log("MEZalarm", "START APP")
	dataset.Data()

	go erd.TrapServer()

	web.Start()
}
