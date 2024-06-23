package main

import (
	"MEZ/dataset"
	"MEZ/logger"
	"MEZ/web"
	"fmt"
)

func main() {
	logger.Log("MEZalarm", "START APP")
	dataset.Data()
	fmt.Println("PSWD:", dataset.PSWD)

	//erd.RelayStatus("10.176.200.222")
	//
	//erd.TrapClient()

	web.Start()
}
