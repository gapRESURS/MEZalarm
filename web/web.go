package web

import (
	"MEZ/dataset"
	"MEZ/logger"
	"net/http"
)

func Start() {
	js := http.FileServer(http.Dir("web/JS/"))
	http.Handle("/JS/", http.StripPrefix("/JS/", js))
	css := http.FileServer(http.Dir("web/CSS/"))
	http.Handle("/CSS/", http.StripPrefix("/CSS/", css))
	img := http.FileServer(http.Dir("web/IMG/"))
	http.Handle("/IMG/", http.StripPrefix("/IMG/", img))

	http.HandleFunc("/API/", handlerAPI)
	http.HandleFunc("/ping/", handlerPing)
	http.HandleFunc("/", handlerIndex)

	logger.Log("WEB server start:", dataset.WEB_IP+":"+dataset.WEB_PORT)
	if err := http.ListenAndServe(":"+dataset.WEB_PORT, nil); err != nil {
		logger.Log("WEB:", err)
	}
}
