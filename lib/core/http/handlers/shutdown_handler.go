package handlers

import (
	"hogo/lib/core/log"
	"net/http"
)

type ShutdownHandler struct {
}

func (s ShutdownHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
	_, err := w.Write([]byte("Service Shutting Down"))
	if err != nil {
		log.Error("ShutdownHandler.ServeHTTP:", err.Error())
	}

	return
}
