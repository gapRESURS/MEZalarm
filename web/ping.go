package web

import (
	"MEZ/logger"
	"encoding/json"
	"github.com/go-ping/ping"
	"net"
	"net/http"
	"time"
)

func handlerPing(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	ip := r.URL.Query().Get("ip") // Получаем значение IP-адреса из запроса
	test := net.ParseIP(ip)
	if ip == "0.0.0.0" || test == nil {
		http.Error(w, "Wrong IP address", http.StatusNotFound)
		return
	}

	pinger, err := ping.NewPinger(ip)
	if err != nil {
		logger.Log("Error NewPinger():", err)
	}

	pinger.SetPrivileged(true)
	pinger.Count = 1
	pinger.Timeout = time.Duration(1000) * time.Millisecond

	pinger.Run()
	stats := pinger.Statistics()

	if stats.PacketLoss == 100 {
		result := map[string]bool{"result": false}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			logger.Log("Error json.Marshal:", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	} else {
		result := map[string]bool{"result": true}
		jsonResult, err := json.Marshal(result)
		if err != nil {
			logger.Log("Error json.Marshal:", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResult)
	}
}
