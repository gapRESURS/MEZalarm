package main

import (
	"MEZ/dataset"
	"MEZ/logger"
	"fmt"
)

func main() {
	logger.Log("START")
	dataset.Data()
	fmt.Println("PSWD:", dataset.PSWD)
}
