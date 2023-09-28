package adapters

import (
	"net/http"
	"log"
)

type HTTPServerAdapter struct {
	ZipFilePath string
}

func (h *HTTPServerAdapter) Start() {
	http.HandleFunc("/download", h.downloadHandler)

	log.Println("Server started on :8082")
	http.ListenAndServe(":8082", nil)
}

func (h *HTTPServerAdapter) downloadHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, h.ZipFilePath)
}
