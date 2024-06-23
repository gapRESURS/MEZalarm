package web

//
//import (
//	"fmt"
//	"net"
//	"net/http"
//	"strings"
//
//	"golang.org/x/net/icmp"
//	"golang.org/x/net/ipv4"
//)
//
//func handlerPing(w http.ResponseWriter, r *http.Request) {
//	if r.Method != "GET" {
//		http.Error(w, "Method is not supported.", http.StatusNotFound)
//		return
//	}
//
//	ip := r.URL.Query().Get("ip") // Получаем значение IP-адреса из запроса
//
//	var msg icmp.Message
//	msg.Type = ipv4.ICMPTypeEchoRequest
//	msg.Code = ipv4.ICMPCodeEchoRequest
//
//	// Создаем ICMP сообщение
//	buf := golang.org / x / net / icmp.NewICMPMessage(&msg, nil)
//	if buf == nil {
//		http.Error(w, "Failed to create ICMP message", http.StatusInternalServerError)
//		return
//	}
//
//	// Отправляем ICMP Echo Request
//	conn, err := net.Dial("ip4:icmp", ip)
//	if err != nil {
//		http.Error(w, w, "Failed to dial IP address", http.StatusInternalServerError)
//		return
//	}
//	defer conn.Close()
//
//	_, err = conn.Write(buf.Bytes())
//	if err != nil {
//		http.Error(w, "Failed to send ICMP packet", http.StatusInternalServerError)
//		return
//	}
//
//	// Читаем ответ
//	reply := make([]byte, 1024)
//	n, _, err := conn.ReadFrom(reply)
//	if err != nil {
//		http.Error(w, "Failed to read ICMP reply", http.StatusInternalServerError)
//		return
//	}
//
//	// Анализируем ответ
//	if n > 0 && !strings.Contains(string(reply[:n]), "Destination Host Unreachable") {
//		w.WriteHeader(http.StatusOK)
//		fmt.Fprint(w, "Пинг успешен")
//	} else {
//		w.WriteHeader(http.StatusServiceUnavailable)
//		fmt.Fprint(w, "Пинг неудачен")
//	}
//
//}
