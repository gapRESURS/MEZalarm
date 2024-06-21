package web

import (
	"MEZ/dataset"
	"MEZ/logger"
	"fmt"
	"net/http"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Здесь должен быть код для чтения данных с DHT22
		// Для примера, используем статические данные
		temperature := 25.0 // Пример температура
		humidity := 60.0    // Пример влажность

		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<meta charset="UTF-8">
				<title>Данные с датчика</title>
				Температура: %f°C<br>
				Влажность: %f%%<br>
				<script>
					function updatePage() {
						fetch('/')
							.then(response => response.text())
							.then(data => document.body.innerHTML = data);
					}
					setInterval(updatePage, 1000); // Обновление каждые 5 секунд
				</script>
			</head>
			<body>
				<p id="data"></p>
			</body>
			</html>
		`, temperature, humidity)
	})
	logger.Log("WEB server start:", dataset.WEB_IP+":"+dataset.WEB_PORT)
	if err := http.ListenAndServe(":"+dataset.WEB_PORT, nil); err != nil {
		logger.Log("WEB:", err)
	}
}
