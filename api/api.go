package api

import (
	"encoding/json"
	"net/http"

	"github.com/ShevonKuan/translate-server/module"
)

func Api(w http.ResponseWriter, r *http.Request) {
	reqj := module.InputObj{}
	err := json.NewDecoder(r.Body).Decode(&reqj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// get query parameter
	translateEngine, _ := module.GetEngine(r)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	res, statusCode, err := module.Engine[translateEngine](&reqj)
	if err != nil || statusCode != http.StatusOK {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	responseJSON := map[string]interface{}{
		"code":         http.StatusOK,
		"data":         res.TransText,
		"alternatives": res.Alternatives,
		"engine":       translateEngine,
	}
	responseBytes, err := json.Marshal(responseJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
}
