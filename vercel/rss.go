package api

import (
	"net/http"

	"github.com/ShevonKuan/translate-server/controller"
)

func Rss(w http.ResponseWriter, r *http.Request) {
	translateEngine := "google"
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "url not found", http.StatusBadRequest)
		return
	}
	resp, err := controller.RSStranslate(url, translateEngine)
	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	output, err := resp.WriteToString()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(output))
}
