package main

import (
	"net/http"
	"strings"
)

func Broadcast(message string, w http.ResponseWriter) (err error) {
	err = bot.Broadcast(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
func RouteFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.Write([]byte("only get supported"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	trimmedUrl := strings.TrimPrefix(r.URL.Path, "/")
	params := strings.Split(trimmedUrl, "/")
	if len(params) == 1 && params[0] == "broadcast" {
		message := r.URL.Query().Get("message")
		if message == "" {
			_, _ = w.Write([]byte("empty message"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		Broadcast(message, w)
	}
}
