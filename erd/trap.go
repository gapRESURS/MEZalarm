package erd

import (
	"MEZ/dataset"
	"MEZ/logger"
	"net"
	"regexp"
	"strings"
)

func extractTextAfterColon(input string) string {
	re := regexp.MustCompile(`:\s*(.*)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) > 0 {
		return matches[1]
	}
	return ""
}

func TrapServer() {
	serverAddress, err := net.ResolveUDPAddr("udp", ":162")

	logger.Log("Start Trap server:", dataset.WEB_IP+":162")
	logger.Log("Trap client:", dataset.OPERATOR.String())

	if err != nil {
		logger.Log("ResolveUDPAddr:", err)
		return
	}

	connection, err := net.ListenUDP("udp", serverAddress)
	if err != nil {
		logger.Log("ListenUDP:", err)
		return
	}
	defer connection.Close()

	for {
		inputBytes := make([]byte, 4096) // Увеличиваем размер буфера для надежной обработки больших пакетов
		n, addr, err := connection.ReadFromUDP(inputBytes)
		if err != nil {
			logger.Log("ReadFromUDP:", err)
			continue
		}
		//port := addr.Port
		//ip := addr.IP
		//protocol := addr.Network()
		message := string(inputBytes[:n])

		//fmt.Println("port:", port)
		//fmt.Println("ip:", ip)
		//fmt.Println("protocol:", protocol)
		//
		//fmt.Println("пароль:", strings.Contains(message, "OTOvK!T0"))
		command := extractTextAfterColon(message)

		if addr.IP.Equal(dataset.OPERATOR) {
			logger.Log("CHECK TRAP ip")
			if strings.Contains(message, dataset.PSWD) {
				logger.Log("CHECK TRAP password")
				relayControl(command)
			} else {
				logger.Log("ERROR CHECK TRAP password")
			}
		} else {
			logger.Log("ERROR CHECK TRAP ip:", addr.IP, ", TRUE IP:", dataset.OPERATOR)
		}

		//logger.Log("TRAP:", command, ip)
		//logger.Log("CHECK TRAP ip:", command, ip)
		//logger.Log("CHECK TRAP password:", command, ip)
	}
}
