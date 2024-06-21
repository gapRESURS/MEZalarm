package main

import (
	"MEZ/dataset"
	"MEZ/erd"
	"MEZ/logger"
	"MEZ/web"
	"fmt"
)

func main() {
	logger.Log("MEZalarm", "START APP")
	dataset.Data()
	fmt.Println("PSWD:", dataset.PSWD)

	erd.TrapClient()

	web.Start()
}
