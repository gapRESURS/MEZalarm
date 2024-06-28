package web

import (
	"MEZ/dataset"
	"MEZ/erd"
	"net/http"
	"strings"
)

func handlerAPI(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")

	if len(parts) != 4 {
		http.Error(w, "Неверный формат URL", http.StatusBadRequest)
		return
	}
	var jsonResult []byte

	switch parts[2] {
	case "NMEZ":
		if dataset.NMEZ_IP == "0.0.0.0" {
			http.Error(w, "Неверный формат URL", http.StatusBadRequest)
			return
		}
		if parts[3] == "DHT22" {
			jsonResult = erd.DHT22Json(dataset.NMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			jsonResult = erd.RelayStatusJson(dataset.NMEZ_IP)
		}
	case "AMEZ":
		if dataset.AMEZ_IP == "0.0.0.0" {
			http.Error(w, "Неверный формат URL", http.StatusBadRequest)
			return
		}
		if parts[3] == "DHT22" {
			jsonResult = erd.DHT22Json(dataset.AMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			jsonResult = erd.RelayStatusJson(dataset.AMEZ_IP)
		}
	case "GMEZ":
		if dataset.GMEZ_IP == "0.0.0.0" {
			http.Error(w, "Неверный формат URL", http.StatusBadRequest)
			return
		}
		if parts[3] == "DHT22" {
			jsonResult = erd.DHT22Json(dataset.GMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			jsonResult = erd.RelayStatusJson(dataset.GMEZ_IP)
		}
	case "SMEZ":
		if dataset.SMEZ_IP == "0.0.0.0" {
			http.Error(w, "Неверный формат URL", http.StatusBadRequest)
			return
		}
		if parts[3] == "DHT22" {
			jsonResult = erd.DHT22Json(dataset.SMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			jsonResult = erd.RelayStatusJson(dataset.SMEZ_IP)
		}
	case "UMEZ":
		if dataset.UMEZ_IP == "0.0.0.0" {
			http.Error(w, "Неверный формат URL", http.StatusBadRequest)
			return
		}
		if parts[3] == "DHT22" {
			jsonResult = erd.DHT22Json(dataset.UMEZ_IP)
		}
		if parts[3] == "RelayStatus" {
			jsonResult = erd.RelayStatusJson(dataset.UMEZ_IP)
		}
	default:
		http.Error(w, "Неверный формат URL", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResult)

}
