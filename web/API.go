package web

import (
	"MEZ/dataset"
	"MEZ/erd"
	"fmt"
	"net/http"
	"strings"
)

func handlerAPI(w http.ResponseWriter, r *http.Request) {
	// Разбиваем URL на части
	parts := strings.Split(r.URL.Path, "/")
	for i, part := range parts {
		fmt.Println(i, part)
	}
	fmt.Println(parts, len(parts))

	// Проверяем, есть ли достаточно частей для извлечения значений
	if len(parts) != 4 {
		http.Error(w, "Неверный формат URL", http.StatusBadRequest)
		return
	}
	var t, h int8
	var rl string
	switch parts[2] {
	case "NMEZ":
		if parts[3] == "DHT22" {
			fmt.Println("DHT22")
			t, h = erd.DHT22(dataset.NMEZ_IP)
			fmt.Println(t, h)
		}
		if parts[3] == "RelayStatus" {
			rl = erd.RelayStatus(dataset.NMEZ_IP)
		}
	case "AMEZ":
		if parts[3] == "DHT22" {
			t, h = erd.DHT22(dataset.AMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			rl = erd.RelayStatus(dataset.AMEZ_IP)
		}
	case "GMEZ":
		if parts[3] == "DHT22" {
			t, h = erd.DHT22(dataset.GMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			rl = erd.RelayStatus(dataset.GMEZ_IP)
		}
	case "SMEZ":
		if parts[3] == "DHT22" {
			t, h = erd.DHT22(dataset.SMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			rl = erd.RelayStatus(dataset.SMEZ_IP)
		}
	case "UMEZ":
		if parts[3] == "DHT22" {
			t, h = erd.DHT22(dataset.UMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			rl = erd.RelayStatus(dataset.UMEZ_IP)
		}
	default:
		http.Error(w, "Неверный формат URL", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Получены параметры: DHT22%s, %s relay %s\n", t, h, rl)

}
